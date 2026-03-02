package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/confidentialhttp"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	"whitewall-cre/contracts/evm/src/generated/validation_registry"
)

func onKYCValidation(
	config *Config,
	runtime cre.Runtime,
	payload *bindings.DecodedLog[validation_registry.ValidationRequestDecoded],
) (*KYCReport, error) {
	logger := runtime.Logger()
	event := payload.Data

	// 1. KYC validator인지 확인
	stripeAddr := common.HexToAddress(config.StripeKYCValidatorAddress)
	if event.ValidatorAddress != stripeAddr {
		logger.Info("Not our validator, skipping")
		return nil, nil
	}

	sessionId := strings.TrimPrefix(event.RequestURI, "stripe:")
	logger.Info("KYC validation request",
		"agentId", event.AgentId,
		"sessionId", sessionId,
		"requestHash", common.BytesToHash(event.RequestHash[:]).Hex())

	// 2. Confidential HTTP로 Stripe API 호출
	verified, err := callStripeAPI(runtime, sessionId)
	if err != nil {
		logger.Error("Stripe API error", "error", err)
		// API 에러여도 리포트는 전송 (verified=false)
		verified = false
	}

	// 3. 리포트 생성 및 전송
	return submitKYCReport(config, runtime, &event, verified, sessionId)
}

func callStripeAPI(runtime cre.Runtime, sessionId string) (bool, error) {
	logger := runtime.Logger()
	client := confidentialhttp.Client{}

	apiUrl := fmt.Sprintf("https://api.stripe.com/v1/identity/verification_sessions/%s", sessionId)

	resp, err := client.SendRequest(runtime,
		&confidentialhttp.ConfidentialHTTPRequest{
			Request: &confidentialhttp.HTTPRequest{
				Url:    apiUrl,
				Method: "GET",
				MultiHeaders: map[string]*confidentialhttp.HeaderValues{
					"Authorization": {
						Values: []string{"Basic {{.STRIPE_SECRET_KEY_B64}}"},
					},
				},
			},
			VaultDonSecrets: []*confidentialhttp.SecretIdentifier{
				{Key: "STRIPE_SECRET_KEY_B64", Namespace: ""},
			},
		}).Await()

	if err != nil {
		return false, fmt.Errorf("HTTP request failed: %w", err)
	}

	// API 응답 파싱
	var result map[string]any
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return false, fmt.Errorf("failed to parse response: %w", err)
	}

	// Stripe API 에러 체크
	if errObj, hasError := result["error"]; hasError {
		errMap, _ := errObj.(map[string]any)
		errMsg, _ := errMap["message"].(string)
		errType, _ := errMap["type"].(string)
		logger.Error("Stripe API returned error",
			"type", errType,
			"message", errMsg)
		return false, fmt.Errorf("Stripe API error: %s - %s", errType, errMsg)
	}

	status, _ := result["status"].(string)
	logger.Info("Stripe verification result",
		"sessionId", sessionId,
		"status", status)

	return status == "verified", nil
}

func submitKYCReport(
	config *Config,
	runtime cre.Runtime,
	event *validation_registry.ValidationRequestDecoded,
	verified bool,
	sessionId string,
) (*KYCReport, error) {
	logger := runtime.Logger()

	// sessionHash 생성 (keccak256 - EVM 표준)
	sessionHash := crypto.Keccak256Hash([]byte(sessionId))

	// ABI 인코딩
	uint256Type, _ := abi.NewType("uint256", "", nil)
	boolType, _ := abi.NewType("bool", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)

	args := abi.Arguments{
		{Name: "agentId", Type: uint256Type},
		{Name: "verified", Type: boolType},
		{Name: "requestHash", Type: bytes32Type},
		{Name: "sessionHash", Type: bytes32Type},
	}

	encodedReport, err := args.Pack(
		event.AgentId,
		verified,
		event.RequestHash,
		sessionHash,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode KYC report: %w", err)
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
	validatorAddr := common.HexToAddress(config.StripeKYCValidatorAddress)
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
		return nil, fmt.Errorf("failed to write KYC report: %w", err)
	}

	logger.Info("KYC report submitted",
		"txHash", common.BytesToHash(writeResult.TxHash).Hex(),
		"agentId", event.AgentId,
		"verified", verified)

	return &KYCReport{
		AgentID:     event.AgentId,
		Verified:    verified,
		RequestHash: event.RequestHash,
		SessionHash: sessionHash,
	}, nil
}
