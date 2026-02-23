//go:build wasip1

package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	"whitewall-cre/contracts/evm/src/generated/identity_registry"
	"whitewall-cre/contracts/evm/src/generated/plaid_credit_validator"
	"whitewall-cre/contracts/evm/src/generated/stripe_kyc_validator"
)

// ===== Config & Types =====
// config.stagine.json애서 읽어옴
type Config struct {
	ChainName                   string `json:"chainName"`
	IdentityRegistryAddress     string `json:"identityRegistryAddress"`
	ConsumerAddress             string `json:"consumerAddress"`
	WorldIDValidatorAddress     string `json:"worldIDValidatorAddress"`
	StripeKYCValidatorAddress   string `json:"stripeKYCValidatorAddress"`
	PlaidCreditValidatorAddress string `json:"plaidCreditValidatorAddress"`
}

type AccessRequest struct {
	AgentID uint64 `json:"agentId"`
}

type AccessReport struct {
	AgentID          *big.Int       `abi:"agentId"`
	Approved         bool           `abi:"approved"`
	Tier             uint8          `abi:"tier"`
	AccountableHuman common.Address `abi:"accountableHuman"`
	Reason           [32]byte       `abi:"reason"`
}

// TierCheckContext - 티어 체크에 필요한 컨텍스트
type TierCheckContext struct {
	Config      *Config
	Runtime     cre.Runtime
	Logger      *slog.Logger
	AgentID     *big.Int
	Owner       common.Address
	Registry    *identity_registry.IdentityRegistry
	StripeKYC   *stripe_kyc_validator.StripeKYCValidator
	PlaidCredit *plaid_credit_validator.PlaidCreditValidator
	EVMClient   *evm.Client
}

// ===== Workflow Entry =====

func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	httpTrigger := http.Trigger(&http.Config{})
	return cre.Workflow[*Config]{
		cre.Handler(httpTrigger, onAccessRequest),
	}, nil
}

// ===== Main Handler =====

func onAccessRequest(config *Config, runtime cre.Runtime, trigger *http.Payload) (*AccessReport, error) {
	logger := runtime.Logger()
	logger.Info("Access request received")

	// 1. HTTP 파싱
	var req AccessRequest
	if err := json.Unmarshal(trigger.GetInput(), &req); err != nil {
		logger.Error("Failed to decode request body", "error", err)
		return finalizeReport(config, runtime, logger, &AccessReport{
			AgentID:  big.NewInt(0),
			Approved: false,
			Tier:     0,
			Reason:   formatReason("INVALID_REQUEST"),
		})
	}

	agentIdBig := new(big.Int).SetUint64(req.AgentID)
	logger.Info("Processing access request", "agentId", req.AgentID)

	// 2. EVM 설정
	chainSelector, _ := evm.ChainSelectorFromName(config.ChainName)
	evmClient := &evm.Client{ChainSelector: chainSelector}

	registryAddress := common.HexToAddress(config.IdentityRegistryAddress)
	registry, err := identity_registry.NewIdentityRegistry(evmClient, registryAddress, nil)
	if err != nil {
		return nil, err
	}

	stripeAddr := common.HexToAddress(config.StripeKYCValidatorAddress)
	stripeVal, _ := stripe_kyc_validator.NewStripeKYCValidator(evmClient, stripeAddr, nil)
	if err != nil {
		return nil, err
	}

	plaidAddr := common.HexToAddress(config.PlaidCreditValidatorAddress)
	plaidVal, _ := plaid_credit_validator.NewPlaidCreditValidator(evmClient, plaidAddr, nil)
	if err != nil {
		return nil, err
	}

	// 3. 티어 체크 시작
	ctx := &TierCheckContext{
		Config:      config,
		Runtime:     runtime,
		Logger:      logger,
		AgentID:     agentIdBig,
		Registry:    registry,
		StripeKYC:   stripeVal,
		PlaidCredit: plaidVal,
		EVMClient:   evmClient,
	}

	// 티어 체크 파이프라인
	report := runTierChecks(ctx)

	return finalizeReport(config, runtime, logger, report)
}

// ===== Tier Check Pipeline =====

func runTierChecks(ctx *TierCheckContext) *AccessReport {
	var report *AccessReport

	// Tier 0 → 1: 등록 체크
	report = checkTier1(ctx)
	if !canProceedToNextTier(report, 1) {
		return report
	}

	// Tier 1 → 2: Human Verification 체크
	report = checkTier2(ctx)
	if !canProceedToNextTier(report, 2) {
		return report
	}

	// Tier 2 → 3: (확장용) 추가 체크
	report = checkTier3(ctx)
	if !canProceedToNextTier(report, 3) {
		return report
	}

	// Tier 3 → 4: (확장용) 최상위 체크
	report = checkTier4(ctx)

	return report
}

func canProceedToNextTier(report *AccessReport, requiredTier uint8) bool {
	return report.Tier >= requiredTier
}

// ===== Tier 1: Registration Check =====
func checkTier1(ctx *TierCheckContext) *AccessReport {
	owner, err := ctx.Registry.OwnerOf(ctx.Runtime, identity_registry.OwnerOfInput{
		TokenId: ctx.AgentID,
	}, big.NewInt(rpc.LatestBlockNumber.Int64())).Await()

	if err != nil {
		ctx.Logger.Warn("Agent not registered", "agentId", ctx.AgentID, "error", err)
		return &AccessReport{
			AgentID:          ctx.AgentID,
			Approved:         false,
			Tier:             0,
			AccountableHuman: common.Address{},
			Reason:           formatReason("NOT_REGISTERED"),
		}
	}

	ctx.Owner = owner
	ctx.Logger.Info("Tier 1: Agent registered", "agentId", ctx.AgentID, "owner", owner.Hex())

	return &AccessReport{
		AgentID:          ctx.AgentID,
		Approved:         false,
		Tier:             1,
		AccountableHuman: owner,
		Reason:           [32]byte{},
	}
}

// ===== Tier 2: Human Verification Check =====
func checkTier2(ctx *TierCheckContext) *AccessReport {
	metadata, err := ctx.Registry.GetMetadata(ctx.Runtime, identity_registry.GetMetadataInput{
		AgentId:     ctx.AgentID,
		MetadataKey: "humanVerified",
	}, big.NewInt(rpc.LatestBlockNumber.Int64())).Await()

	if err != nil || len(metadata) < 160 {
		ctx.Logger.Warn("Human verification metadata missing", "agentId", ctx.AgentID, "error", err)
		return &AccessReport{
			AgentID:          ctx.AgentID,
			Approved:         false,
			Tier:             1,
			AccountableHuman: ctx.Owner,
			Reason:           formatReason("NO_VERIFICATION_DATA"),
		}
	}

	// metadata 파싱
	isVerified := metadata[31] == 1
	validatorAddr := common.BytesToAddress(metadata[32+12 : 64])
	requesterAddr := common.BytesToAddress(metadata[128+12 : 160])
	trustedValidator := common.HexToAddress(ctx.Config.WorldIDValidatorAddress)

	if isVerified && validatorAddr == trustedValidator && requesterAddr == ctx.Owner {
		ctx.Logger.Info("Tier 2: Human verification passed",
			"validator", validatorAddr.Hex(),
			"requester", requesterAddr.Hex())
		return &AccessReport{
			AgentID:          ctx.AgentID,
			Approved:         true,
			Tier:             2,
			AccountableHuman: ctx.Owner,
			Reason:           [32]byte{},
		}
	}

	ctx.Logger.Warn("Invalid human verification", "agentId", ctx.AgentID)
	return &AccessReport{
		AgentID:          ctx.AgentID,
		Approved:         false,
		Tier:             1,
		AccountableHuman: ctx.Owner,
		Reason:           formatReason("INVALID_HUMAN_DATA"),
	}
}

// ===== Tier 3: KYC (Stripe) =====
func checkTier3(ctx *TierCheckContext) *AccessReport {
	// 1. 바인딩 파일의 IsKYCVerified 함수는 (runtime, IsKYCVerifiedInput, blockNumber)를 인자로 받습니다.
	// 2. IsKYCVerifiedInput은 포인터가 아니라 일반 구조체 값!
	promise := ctx.StripeKYC.IsKYCVerified(
		ctx.Runtime,
		stripe_kyc_validator.IsKYCVerifiedInput{
			AgentId: ctx.AgentID,
		},
		nil, // 최신 블록을 보려면 nil을 넣으면 바인딩 내부에서 처리를 해줍니다.
	)

	// 3. Promise의 결과를 Await()으로 기다립니다.
	isVerified, err := promise.Await()

	if err != nil || !isVerified {
		ctx.Logger.Warn("Tier 3 Check Failed (KYC)", "agentId", ctx.AgentID, "error", err)
		return &AccessReport{
			AgentID:          ctx.AgentID,
			Approved:         true,
			Tier:             2,
			AccountableHuman: ctx.Owner,
			Reason:           formatReason("KYC_REQUIRED"),
		}
	}

	ctx.Logger.Info("Tier 3 Check Passed", "agentId", ctx.AgentID)
	return &AccessReport{
		AgentID:          ctx.AgentID,
		Approved:         true,
		Tier:             3,
		AccountableHuman: ctx.Owner,
		Reason:           [32]byte{},
	}
}

// ===== Tier 4: Credit (Plaid) =====
func checkTier4(ctx *TierCheckContext) *AccessReport {
	// 1. Plaid도 동일하게 GetCreditScoreInput 구조체 값을 인자로 받습니다
	promise := ctx.PlaidCredit.GetCreditScore(
		ctx.Runtime,
		plaid_credit_validator.GetCreditScoreInput{
			AgentId: ctx.AgentID,
		},
		nil,
	)

	score, err := promise.Await()

	// 신용 점수 문턱값 50점
	if err != nil || score < 50 {
		ctx.Logger.Warn("Tier 4 Check Failed (Credit)", "agentId", ctx.AgentID, "score", score, "error", err)
		return &AccessReport{
			AgentID:          ctx.AgentID,
			Approved:         true,
			Tier:             3, // KYC는 통과했으므로 Tier 3 유지
			AccountableHuman: ctx.Owner,
			Reason:           formatReason("LOW_CREDIT_SCORE"),
		}
	}

	ctx.Logger.Info("Tier 4 Check Passed (Final)", "agentId", ctx.AgentID, "score", score)

	//tier 4 통과 ㅋ
	return &AccessReport{
		AgentID:          ctx.AgentID,
		Approved:         true,
		Tier:             4,
		AccountableHuman: ctx.Owner,
		Reason:           [32]byte{},
	}
}

// ===== Report Submission =====

func finalizeReport(config *Config, runtime cre.Runtime, logger *slog.Logger, report *AccessReport) (*AccessReport, error) {
	if err := submitReport(config, runtime, report, logger); err != nil {
		logger.Error("Failed to submit report", "error", err)
	}
	return report, nil
}

func submitReport(config *Config, runtime cre.Runtime, report *AccessReport, logger *slog.Logger) error {
	// 1. 각 필드별 타입 정의
	uint256Type, _ := abi.NewType("uint256", "", nil)
	boolType, _ := abi.NewType("bool", "", nil)
	uint8Type, _ := abi.NewType("uint8", "", nil)
	addressType, _ := abi.NewType("address", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)

	// 2. Arguments 정의 (5개 개별 필드)
	args := abi.Arguments{
		{Name: "agentId", Type: uint256Type},
		{Name: "approved", Type: boolType},
		{Name: "tier", Type: uint8Type},
		{Name: "accountableHuman", Type: addressType},
		{Name: "reason", Type: bytes32Type},
	}

	// 3. Pack (5개 인자)
	encodedReport, err := args.Pack(
		report.AgentID,
		report.Approved,
		report.Tier,
		report.AccountableHuman,
		report.Reason,
	)
	if err != nil {
		return fmt.Errorf("failed to encode report: %w", err)
	}

	// 2. 서명된 Report 생성
	signedReport, err := runtime.GenerateReport(&cre.ReportRequest{
		EncodedPayload: encodedReport,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	}).Await()
	if err != nil {
		return fmt.Errorf("failed to generate report: %w", err)
	}

	// 3. Consumer Contract에 제출
	consumerAddr := common.HexToAddress(config.ConsumerAddress)
	chainSelector, _ := evm.ChainSelectorFromName(config.ChainName)
	evmClient := &evm.Client{ChainSelector: chainSelector}

	writeResult, err := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver: consumerAddr.Bytes(),
		Report:   signedReport,
	}).Await()
	if err != nil {
		return fmt.Errorf("failed to write report: %w", err)
	}

	logger.Info("Report submitted",
		"txHash", writeResult.TxHash,
		"status", writeResult.TxStatus,
		"agentId", report.AgentID,
		"tier", report.Tier,
		"approved", report.Approved)

	return nil
}

// ===== Helpers =====

func formatReason(reason string) [32]byte {
	var b [32]byte
	copy(b[:], []byte(reason))
	return b
}
