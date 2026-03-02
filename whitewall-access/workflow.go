//go:build wasip1

package main

import (
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	"whitewall-cre/contracts/evm/src/generated/validation_registry"
)

// ===== Workflow Entry =====
func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	httpTrigger := http.Trigger(&http.Config{})

	// EVM 설정
	chainSelector, err := evm.ChainSelectorFromName(config.ChainName)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain selector: %w", err)
	}
	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}
	registryAddr := common.HexToAddress(config.ValidationRegistryAddress)
	registry, _ := validation_registry.NewValidationRegistry(evmClient, registryAddr, nil)

	//Log trigger -kyc용
	kycLogTrigger, _ := registry.LogTriggerValidationRequestLog(
		chainSelector,
		evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
		[]validation_registry.ValidationRequestTopics{},
	)
	//Log Trigger - Credit용 (같은 이벤트, 핸들러에서 필터링)
	creditLogTrigger, _ := registry.LogTriggerValidationRequestLog(
		chainSelector,
		evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
		[]validation_registry.ValidationRequestTopics{},
	)

	return cre.Workflow[*Config]{
		cre.Handler(httpTrigger, onAccessRequest),         // index 0 - Tier 체크
		cre.Handler(kycLogTrigger, onKYCValidation),       // index 1 - KYC
		cre.Handler(creditLogTrigger, onCreditValidation), // index 2 - Credit
	}, nil
}
