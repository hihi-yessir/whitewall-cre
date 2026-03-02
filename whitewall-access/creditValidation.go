package main

import (
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

// Credit Validation Flow:
// 1. ValidationRegistry event → CreditValidator 주소 확인
// 2. RequestURI에서 "plaid:" 제거 → public_token 추출
// 3. Confidential HTTP로 Proxy 호출 (client_id, secret은 DON Vault에서 template injection)
// 4. Proxy가 Plaid API 호출 → score 반환
// 5. CRE에서 DON 서명 + Report 제출

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

	// 3. Proxy 호출 → score 반환
	score, err := callPlaidProxy(config, runtime, event.AgentId, event.RequestHash, publicToken)
	if err != nil {
		logger.Error("Failed to get score from proxy", "error", err)
		return nil, err
	}

	logger.Info("Credit score received from proxy",
		"agentId", event.AgentId,
		"score", score)

	// 4. DON 서명 + Report 제출
	return submitCreditReport(config, runtime, &event, score)
}

// callPlaidProxy: Confidential HTTP로 Proxy 호출
// client_id, secret은 DON Vault에서 template injection
func callPlaidProxy(
	config *Config,
	runtime cre.Runtime,
	agentId *big.Int,
	requestHash [32]byte,
	publicToken string,
) (uint8, error) {
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

	logger.Info("Calling Plaid proxy...",
		"url", config.CreditProxyURL,
		"agentId", agentId.String())

	resp, err := client.SendRequest(runtime,
		&confidentialhttp.ConfidentialHTTPRequest{
			Request: &confidentialhttp.HTTPRequest{
				Url:    config.CreditProxyURL,
				Method: "POST",
				MultiHeaders: map[string]*confidentialhttp.HeaderValues{
					"Content-Type": {Values: []string{"application/json"}},
				},
				Body: &confidentialhttp.HTTPRequest_BodyString{BodyString: body},
				// EncryptOutput: false - score는 민감하지 않음
			},
			VaultDonSecrets: []*confidentialhttp.SecretIdentifier{
				{Key: "PLAID_CLIENT_ID", Namespace: ""},
				{Key: "PLAID_SECRET", Namespace: ""},
			},
		}).Await()

	if err != nil {
		return 0, fmt.Errorf("proxy call failed: %w", err)
	}

	logger.Info("Proxy response", "status", resp.StatusCode, "body", string(resp.Body))

	// 응답 파싱
	var proxyResp struct {
		Score   uint8  `json:"score"`
		Success bool   `json:"success"`
		Error   string `json:"error,omitempty"`
	}
	if err := json.Unmarshal(resp.Body, &proxyResp); err != nil {
		return 0, fmt.Errorf("failed to parse proxy response: %w", err)
	}

	if !proxyResp.Success {
		return 0, fmt.Errorf("proxy error: %s", proxyResp.Error)
	}

	return proxyResp.Score, nil
}

// submitCreditReport: DON 서명 후 온체인 제출
func submitCreditReport(
	config *Config,
	runtime cre.Runtime,
	event *validation_registry.ValidationRequestDecoded,
	score uint8,
) (*CreditReport, error) {
	logger := runtime.Logger()

	// tokenHash는 requestHash 재사용
	tokenHash := event.RequestHash

	// ABI 인코딩
	uint256Type, _ := abi.NewType("uint256", "", nil)
	uint8Type, _ := abi.NewType("uint8", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)

	args := abi.Arguments{
		{Name: "agentId", Type: uint256Type},
		{Name: "score", Type: uint8Type},
		{Name: "requestHash", Type: bytes32Type},
		{Name: "tokenHash", Type: bytes32Type},
	}

	encodedReport, err := args.Pack(
		event.AgentId,
		score,
		event.RequestHash,
		tokenHash,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode credit report: %w", err)
	}

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

	writeResult, err := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver: validatorAddr.Bytes(),
		Report:   signedReport,
		GasConfig: &evm.GasConfig{
			GasLimit: config.GasLimit,
		},
	}).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to write credit report: %w", err)
	}

	logger.Info("Credit report submitted",
		"txHash", common.BytesToHash(writeResult.TxHash).Hex(),
		"agentId", event.AgentId,
		"score", score)

	return &CreditReport{
		AgentID:     event.AgentId,
		Score:       score,
		RequestHash: event.RequestHash,
		TokenHash:   tokenHash,
	}, nil
}
