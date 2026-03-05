package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/confidentialhttp"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	"whitewall-cre/contracts/evm/src/generated/validation_registry"
)

// Credit Validation Flow (TEE 기반):
// 1. ValidationRegistry event → CreditValidator 주소 확인
// 2. RequestURI에서 "plaid:" 제거 → public_token 추출
// 3. Confidential HTTP로 TEE Credit Oracle 호출 (client_id, secret은 DON Vault에서 template injection)
// 4. TEE가 SGX Enclave 내에서 Plaid API 호출 → score + SGX Quote 반환
// 5. CRE에서 V2 리포트 인코딩 (quote 포함) → DON 서명 → WriteReport

func onCreditValidation(
	config *Config,
	runtime cre.Runtime,
	payload *bindings.DecodedLog[validation_registry.ValidationRequestDecoded],
) (*CreditReport, error) {
	logger := runtime.Logger()
	event := payload.Data

	// 1. Credit validator인지 확인
	plaidAddr := common.HexToAddress(config.PlaidCreditValidatorAddress)
	if event.ValidatorAddress != plaidAddr {
		logger.Info("Not our validator, skipping")
		return nil, nil
	}

	// 2. RequestURI에서 public_token 추출 ("plaid:xxxxx" → "xxxxx")
	publicToken := strings.TrimPrefix(event.RequestURI, "plaid:")

	// publicToken 로깅 (길이 체크)
	tokenPreview := publicToken
	if len(publicToken) > 20 {
		tokenPreview = publicToken[:20] + "..."
	}
	logger.Info("Credit validation request received",
		"agentId", event.AgentId,
		"publicToken", tokenPreview,
		"requestHash", common.BytesToHash(event.RequestHash[:]).Hex())

	// 3. TEE Credit Oracle 호출 → score + quote 반환
	score, sgxQuote, err := callTEECreditOracle(config, runtime, event.AgentId, event.RequestHash, publicToken)
	if err != nil {
		logger.Error("Failed to get score from TEE", "error", err)
		return nil, err
	}

	logger.Info("Credit score received from TEE",
		"agentId", event.AgentId,
		"score", score,
		"quoteLength", len(sgxQuote))

	// 4. V2 리포트 인코딩 (quote 포함) → DON 서명 → WriteReport
	return submitCreditReportV2(config, runtime, &event, score, sgxQuote)
}

// callTEECreditOracle: Confidential HTTP로 TEE Credit Oracle 호출
// client_id, secret은 DON Vault에서 template injection
// TEE(Intel SGX)가 Plaid API 호출 후 score + SGX Quote 반환
func callTEECreditOracle(
	config *Config,
	runtime cre.Runtime,
	agentId *big.Int,
	requestHash [32]byte,
	publicToken string,
) (uint8, []byte, error) {
	logger := runtime.Logger()
	client := confidentialhttp.Client{}

	// Template injection으로 DON Vault 값 주입
	body := fmt.Sprintf(`{
		"clientId": "{{.PLAID_CLIENT_ID}}",
		"secret": "{{.PLAID_SECRET}}",
		"publicToken": "%s",
		"agentId": "%s",
		"requestHash": "%s"
	}`, publicToken, agentId.String(), common.BytesToHash(requestHash[:]).Hex())

	logger.Info("Calling TEE Credit Oracle...",
		"url", config.CreditTEEUrl,
		"agentId", agentId.String())

	resp, err := client.SendRequest(runtime,
		&confidentialhttp.ConfidentialHTTPRequest{
			Request: &confidentialhttp.HTTPRequest{
				Url:    config.CreditTEEUrl,
				Method: "POST",
				MultiHeaders: map[string]*confidentialhttp.HeaderValues{
					"Content-Type": {Values: []string{"application/json"}},
				},
				Body: &confidentialhttp.HTTPRequest_BodyString{BodyString: body},
			},
			VaultDonSecrets: []*confidentialhttp.SecretIdentifier{
				{Key: "PLAID_CLIENT_ID", Namespace: ""},
				{Key: "PLAID_SECRET", Namespace: ""},
			},
		}).Await()

	if err != nil {
		return 0, nil, fmt.Errorf("TEE call failed: %w", err)
	}

	logger.Info("TEE response received", "status", resp.StatusCode)

	// 응답 파싱 (score + quote)
	var teeResp struct {
		Score   uint8  `json:"score"`
		Success bool   `json:"success"`
		Quote   string `json:"quote"` // SGX DCAP Quote (hex string)
		Error   string `json:"error,omitempty"`
	}
	if err := json.Unmarshal(resp.Body, &teeResp); err != nil {
		return 0, nil, fmt.Errorf("failed to parse TEE response: %w", err)
	}

	if !teeResp.Success {
		return 0, nil, fmt.Errorf("TEE error: %s", teeResp.Error)
	}

	// Quote hex string → bytes
	sgxQuote := common.FromHex(teeResp.Quote)
	logger.Info("SGX Quote received", "quoteLength", len(sgxQuote))

	return teeResp.Score, sgxQuote, nil
}

// submitCreditReportV2: V2 리포트 인코딩 (SGX Quote 포함) → DON 서명 → WriteReport
// V2 인코딩: abi.encode(agentId, score, requestHash, dataHash, sgxQuote)
func submitCreditReportV2(
	config *Config,
	runtime cre.Runtime,
	event *validation_registry.ValidationRequestDecoded,
	score uint8,
	sgxQuote []byte,
) (*CreditReport, error) {
	logger := runtime.Logger()

	// dataHash 계산: sha256("agent:{agentId}|hash:{requestHash}|score:{score}")
	// TEE가 생성한 reportData와 동일한 형식
	dataString := fmt.Sprintf("agent:%s|hash:%s|score:%d",
		event.AgentId.String(),
		common.BytesToHash(event.RequestHash[:]).Hex(),
		score)
	dataHash := sha256.Sum256([]byte(dataString))

	logger.Info("V2 Report encoding",
		"dataString", dataString,
		"dataHash", common.BytesToHash(dataHash[:]).Hex())

	// V2 ABI 인코딩: (agentId, score, requestHash, dataHash, sgxQuote)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	uint8Type, _ := abi.NewType("uint8", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)
	bytesType, _ := abi.NewType("bytes", "", nil)

	args := abi.Arguments{
		{Name: "agentId", Type: uint256Type},
		{Name: "score", Type: uint8Type},
		{Name: "requestHash", Type: bytes32Type},
		{Name: "dataHash", Type: bytes32Type},
		{Name: "sgxQuote", Type: bytesType},
	}

	encodedReport, err := args.Pack(
		event.AgentId,
		score,
		event.RequestHash,
		dataHash,
		sgxQuote,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode V2 credit report: %w", err)
	}

	logger.Info("V2 Report encoded", "reportLength", len(encodedReport))

	// DON 서명
	signedReport, err := runtime.GenerateReport(&cre.ReportRequest{
		EncodedPayload: encodedReport,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	}).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to generate report: %w", err)
	}

	// Forwarder로 전송
	validatorAddr := common.HexToAddress(config.PlaidCreditValidatorAddress)
	chainSelector, _ := evm.ChainSelectorFromName(config.ChainName)
	evmClient := &evm.Client{ChainSelector: chainSelector}

	logger.Info("Credit V2 WriteReport 준비중",
		"configGasLimit", config.GasLimit,
		"receiver", validatorAddr.Hex(),
		"chain", config.ChainName,
		"quoteLength", len(sgxQuote))

	writeResult, err := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver: validatorAddr.Bytes(),
		Report:   signedReport,
		GasConfig: &evm.GasConfig{
			GasLimit: config.GasLimit,
		},
	}).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to write V2 credit report: %w", err)
	}

	logger.Info("Credit V2 report submitted",
		"txHash", common.BytesToHash(writeResult.TxHash).Hex(),
		"agentId", event.AgentId,
		"score", score)

	return &CreditReport{
		AgentID:     event.AgentId,
		Score:       score,
		RequestHash: event.RequestHash,
		TokenHash:   dataHash, // V2에서는 dataHash를 TokenHash 필드에 저장
	}, nil
}
