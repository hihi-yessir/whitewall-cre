//go:build wasip1

package main

import (
	"encoding/json"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	"whitewall-cre/contracts/evm/src/generated/identity_registry"
)

// Config - config.staging.json에서 읽어옴
type Config struct {
	ChainName               string `json:"chainName"`
	IdentityRegistryAddress string `json:"identityRegistryAddress"`
	ConsumerAddress         string `json:"consumerAddress"`
	WorldIDValidatorAddress string `json:"worldIDValidatorAddress"`
}

// http 요청 페이로드
type AccessRequest struct {
	AgentID uint64 `json:"agentId"`
}

// report 결과
type AccessReport struct {
	AgentID          *big.Int       `abi:"agentId"`
	Approved         bool           `abi:"approved"`
	Tier             uint8          `abi:"tier"`
	AccountableHuman common.Address `abi:"accountableHuman"`
	Reason           [32]byte       `abi:"reason"`
}

// InitWorkflow
func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	//HTTP 트리거
	httpTrigger := http.Trigger(&http.Config{}) //HTTP 트리거 인스턴스

	return cre.Workflow[*Config]{
		cre.Handler(httpTrigger, onAccessRequest), //트리거-함수 연결
	}, nil
}

// 접근 요청 처리 핸들러
func onAccessRequest(config *Config, runtime cre.Runtime, trigger *http.Payload) (*AccessReport, error) {
	logger := runtime.Logger()
	//초기값
	currentTier := uint8(0)
	logger.Info("Access request received")

	// 1. HTTP - agentId 파싱
	var req AccessRequest

	if err := json.Unmarshal(trigger.GetInput(), &req); err != nil {
		logger.Error("Failed to decode request body", "error", err)
		// 리턴할 때도 구조체 타입을 맞춰야 빨간 줄이 안 납니다!
		return &AccessReport{
			AgentID:          big.NewInt(0),
			Approved:         false,
			Tier:             currentTier,
			AccountableHuman: common.Address{},
			Reason:           formatReason("INVALID_REQUEST"),
		}, nil
	}

	agentIdBig := new(big.Int).SetUint64(req.AgentID)
	logger.Info("Processing access request", "agentId", req.AgentID)

	// 2. EVM client 및 Registry 설정
	chainSelector, _ := evm.ChainSelectorFromName(config.ChainName)
	evmClient := &evm.Client{ChainSelector: chainSelector}
	registryAddress := common.HexToAddress(config.IdentityRegistryAddress)
	registry, err := identity_registry.NewIdentityRegistry(evmClient, registryAddress, nil)
	if err != nil {
		return nil, err
	}

	// 3. CRE 조건 체크 시작
	// 조건 1 8004(IdentityRegistry) ownerOf 호출 (등록여부)
	owner, err := registry.OwnerOf(runtime, identity_registry.OwnerOfInput{TokenId: agentIdBig}, big.NewInt(rpc.LatestBlockNumber.Int64())).Await()

	if err != nil {
		logger.Warn("Agent not registered", "agentId", req.AgentID, "error", err)
		return &AccessReport{
			AgentID:          agentIdBig,
			Approved:         false,
			Tier:             currentTier,
			AccountableHuman: common.Address{}, //zero address
			Reason:           formatReason("NOT_REGISTERED"),
		}, nil
	}
	//등록 확인 -> Tier 1로 업뎃
	currentTier = 1
	logger.Info("Tier 1: Agent registered", "agentId", req.AgentID, "owner", owner.Hex())

	//3. 조건 2 getMetadata 호출 (humanVerified 태그 확인)
	metadata, err := registry.GetMetadata(runtime, identity_registry.GetMetadataInput{
		AgentId:     agentIdBig,
		MetadataKey: "humanVerified",
	}, big.NewInt(rpc.LatestBlockNumber.Int64())).Await()

	if err != nil || len(metadata) < 160 {
		logger.Warn("Human verification metadata missing or invalid length", "agentId", req.AgentID, "error", err)
		return &AccessReport{
			AgentID:          agentIdBig,
			Approved:         false,
			Tier:             1, // 등록은 확인됐으니 Tier 1은 유지
			AccountableHuman: owner,
			Reason:           formatReason("NO_VERIFICATION_DATA"),
		}, nil
	}

	// meatada[humanVerified] 데이터 파싱
	isVerified := metadata[31] == 1                                         //true 여서 막비트 1이어야함.(isVerified: 첫 번째 32바이트의 마지막 바이트)
	validatorAddr := common.BytesToAddress(metadata[32+12 : 64])            //(validatorAddr: 두 번째 32바이트 (앞 12바이트 패딩 제외))
	requesterAddr := common.BytesToAddress(metadata[128+12 : 160])          //(requesterAddr: 다섯 번째 32바이트 (앞 12바이트 패딩 제외))
	trustedValidator := common.HexToAddress(config.WorldIDValidatorAddress) //우리가 설정한 Validator 주소와 비교

	// 인증 값 && WorldIDValidator인지 && 그 검증 요청자가 owner인지! (태그는 있고 주인이 바뀌었을 수도 있으니!!!! 중요)
	if isVerified && validatorAddr == trustedValidator && requesterAddr == owner {
		logger.Info("Tier 2 Check: Human verification verified",
			"validator", validatorAddr.Hex(),
			"requester", requesterAddr.Hex())
		currentTier = 2
	} else {
		// 데이터는 있지만 값이 틀린 경우 (위조된 데이터거나 다른 사람이 인증한 경우)
		logger.Warn("Invalid human verification details",
			"agentId", req.AgentID,
			"isVerified", isVerified,
			"isTrustedValidator", validatorAddr == trustedValidator,
			"isOwnerRequester", requesterAddr == owner)

		return &AccessReport{
			AgentID:          agentIdBig,
			Approved:         false,
			Tier:             currentTier,
			AccountableHuman: owner,
			Reason:           formatReason("INVALID_HUMAN_DATA"),
		}, nil
	}

	// 4. 리포트 최종 생성
	return &AccessReport{
		AgentID:          agentIdBig,
		Approved:         true,
		Tier:             currentTier, //계산 로직 있어야함!!
		AccountableHuman: owner,
		Reason:           [32]byte{}, // 0x00...
	}, nil
}

// 헬퍼 함수: string -> bytes32
func formatReason(reason string) [32]byte {
	var b [32]byte
	copy(b[:], []byte(reason))
	return b
}
