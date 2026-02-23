// Code generated — DO NOT EDIT.

//go:build !wasip1

package plaid_credit_validator

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
)

var (
	_ = errors.New
	_ = fmt.Errorf
	_ = big.NewInt
	_ = common.Big1
)

// PlaidCreditValidatorMock is a mock implementation of PlaidCreditValidator for testing.
type PlaidCreditValidatorMock struct {
	UPGRADEINTERFACEVERSION func() (string, error)
	GetConfig               func() (GetConfigOutput, error)
	GetCreditData           func(GetCreditDataInput) (GetCreditDataOutput, error)
	GetCreditScore          func(GetCreditScoreInput) (uint8, error)
	HasCreditScore          func(HasCreditScoreInput) (bool, error)
	Owner                   func() (common.Address, error)
	ProxiableUUID           func() ([32]byte, error)
}

// NewPlaidCreditValidatorMock creates a new PlaidCreditValidatorMock for testing.
func NewPlaidCreditValidatorMock(address common.Address, clientMock *evmmock.ClientCapability) *PlaidCreditValidatorMock {
	mock := &PlaidCreditValidatorMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["UPGRADE_INTERFACE_VERSION"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.UPGRADEINTERFACEVERSION == nil {
				return nil, errors.New("UPGRADE_INTERFACE_VERSION method not mocked")
			}
			result, err := mock.UPGRADEINTERFACEVERSION()
			if err != nil {
				return nil, err
			}
			return abi.Methods["UPGRADE_INTERFACE_VERSION"].Outputs.Pack(result)
		},
		string(abi.Methods["getConfig"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetConfig == nil {
				return nil, errors.New("getConfig method not mocked")
			}
			result, err := mock.GetConfig()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getConfig"].Outputs.Pack(
				result.ForwarderAddress,
				result.IdentityRegistry,
				result.ValidationRegistry,
			)
		},
		string(abi.Methods["getCreditData"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetCreditData == nil {
				return nil, errors.New("getCreditData method not mocked")
			}
			inputs := abi.Methods["getCreditData"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetCreditDataInput{
				AgentId: values[0].(*big.Int),
			}

			result, err := mock.GetCreditData(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getCreditData"].Outputs.Pack(
				result.Score,
				result.DataHash,
				result.VerifiedAt,
				result.HasScore,
			)
		},
		string(abi.Methods["getCreditScore"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetCreditScore == nil {
				return nil, errors.New("getCreditScore method not mocked")
			}
			inputs := abi.Methods["getCreditScore"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetCreditScoreInput{
				AgentId: values[0].(*big.Int),
			}

			result, err := mock.GetCreditScore(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getCreditScore"].Outputs.Pack(result)
		},
		string(abi.Methods["hasCreditScore"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.HasCreditScore == nil {
				return nil, errors.New("hasCreditScore method not mocked")
			}
			inputs := abi.Methods["hasCreditScore"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := HasCreditScoreInput{
				AgentId: values[0].(*big.Int),
			}

			result, err := mock.HasCreditScore(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["hasCreditScore"].Outputs.Pack(result)
		},
		string(abi.Methods["owner"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Owner == nil {
				return nil, errors.New("owner method not mocked")
			}
			result, err := mock.Owner()
			if err != nil {
				return nil, err
			}
			return abi.Methods["owner"].Outputs.Pack(result)
		},
		string(abi.Methods["proxiableUUID"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.ProxiableUUID == nil {
				return nil, errors.New("proxiableUUID method not mocked")
			}
			result, err := mock.ProxiableUUID()
			if err != nil {
				return nil, err
			}
			return abi.Methods["proxiableUUID"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
