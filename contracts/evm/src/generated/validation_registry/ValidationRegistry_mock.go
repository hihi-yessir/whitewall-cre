// Code generated — DO NOT EDIT.

//go:build !wasip1

package validation_registry

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

// ValidationRegistryMock is a mock implementation of ValidationRegistry for testing.
type ValidationRegistryMock struct {
	UPGRADEINTERFACEVERSION func() (string, error)
	GetAgentValidations     func(GetAgentValidationsInput) ([][32]byte, error)
	GetIdentityRegistry     func() (common.Address, error)
	GetSummary              func(GetSummaryInput) (GetSummaryOutput, error)
	GetValidationStatus     func(GetValidationStatusInput) (GetValidationStatusOutput, error)
	GetValidatorRequests    func(GetValidatorRequestsInput) ([][32]byte, error)
	Owner                   func() (common.Address, error)
	ProxiableUUID           func() ([32]byte, error)
}

// NewValidationRegistryMock creates a new ValidationRegistryMock for testing.
func NewValidationRegistryMock(address common.Address, clientMock *evmmock.ClientCapability) *ValidationRegistryMock {
	mock := &ValidationRegistryMock{}

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
		string(abi.Methods["getAgentValidations"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetAgentValidations == nil {
				return nil, errors.New("getAgentValidations method not mocked")
			}
			inputs := abi.Methods["getAgentValidations"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetAgentValidationsInput{
				AgentId: values[0].(*big.Int),
			}

			result, err := mock.GetAgentValidations(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getAgentValidations"].Outputs.Pack(result)
		},
		string(abi.Methods["getIdentityRegistry"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetIdentityRegistry == nil {
				return nil, errors.New("getIdentityRegistry method not mocked")
			}
			result, err := mock.GetIdentityRegistry()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getIdentityRegistry"].Outputs.Pack(result)
		},
		string(abi.Methods["getSummary"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetSummary == nil {
				return nil, errors.New("getSummary method not mocked")
			}
			inputs := abi.Methods["getSummary"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 3 {
				return nil, errors.New("expected 3 input values")
			}

			args := GetSummaryInput{
				AgentId:            values[0].(*big.Int),
				ValidatorAddresses: values[1].([]common.Address),
				Tag:                values[2].(string),
			}

			result, err := mock.GetSummary(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getSummary"].Outputs.Pack(
				result.Count,
				result.AvgResponse,
			)
		},
		string(abi.Methods["getValidationStatus"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetValidationStatus == nil {
				return nil, errors.New("getValidationStatus method not mocked")
			}
			inputs := abi.Methods["getValidationStatus"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetValidationStatusInput{
				RequestHash: values[0].([32]byte),
			}

			result, err := mock.GetValidationStatus(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getValidationStatus"].Outputs.Pack(
				result.ValidatorAddress,
				result.AgentId,
				result.Response,
				result.ResponseHash,
				result.Tag,
				result.LastUpdate,
			)
		},
		string(abi.Methods["getValidatorRequests"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetValidatorRequests == nil {
				return nil, errors.New("getValidatorRequests method not mocked")
			}
			inputs := abi.Methods["getValidatorRequests"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetValidatorRequestsInput{
				ValidatorAddress: values[0].(common.Address),
			}

			result, err := mock.GetValidatorRequests(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getValidatorRequests"].Outputs.Pack(result)
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
