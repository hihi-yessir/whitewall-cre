package main

import (
	"log/slog"
	"math/big"
	"whitewall-cre/contracts/evm/src/generated/identity_registry"
	"whitewall-cre/contracts/evm/src/generated/plaid_credit_validator"
	"whitewall-cre/contracts/evm/src/generated/stripe_kyc_validator"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

// ===== Config ===== config.stagine.json애서 읽어옴
type Config struct {
	ChainName                   string `json:"chainName"`
	IdentityRegistryAddress     string `json:"identityRegistryAddress"`
	ConsumerAddress             string `json:"consumerAddress"`
	WorldIDValidatorAddress     string `json:"worldIDValidatorAddress"`
	StripeKYCValidatorAddress   string `json:"stripeKYCValidatorAddress"`
	PlaidCreditValidatorAddress string `json:"plaidCreditValidatorAddress"`
	ForwarderAddress            string `json:"forwarderAddress"`
	ValidationRegistryAddress   string `json:"validationRegistryAddress"`
	GasLimit                    uint64 `json:"gasLimit"`
	CreditProxyURL              string `json:"creditProxyUrl"` // Plaid proxy URL for credit score calculation
}

// ===== Request/Response Types =====
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

// ===== Tier Check Context ===== 티어 체크에 필요한 컨텍스트
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

// ===== KYC Types =====
type KYCReport struct {
	AgentID     *big.Int
	Verified    bool
	RequestHash [32]byte
	SessionHash [32]byte
}

// ===== Credit Types =====
type CreditReport struct {
	AgentID     *big.Int
	Score       uint8
	RequestHash [32]byte
	TokenHash   [32]byte
}
