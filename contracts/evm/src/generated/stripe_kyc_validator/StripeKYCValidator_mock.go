// Code generated — DO NOT EDIT.

//go:build !wasip1

package stripe_kyc_validator

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

// StripeKYCValidatorMock is a mock implementation of StripeKYCValidator for testing.
type StripeKYCValidatorMock struct {
	UPGRADEINTERFACEVERSION func() (string, error)
	GetConfig               func() (GetConfigOutput, error)
	GetKYCData              func(GetKYCDataInput) (GetKYCDataOutput, error)
	IsKYCVerified           func(IsKYCVerifiedInput) (bool, error)
	Owner                   func() (common.Address, error)
	ProxiableUUID           func() ([32]byte, error)
}

// NewStripeKYCValidatorMock creates a new StripeKYCValidatorMock for testing.
func NewStripeKYCValidatorMock(address common.Address, clientMock *evmmock.ClientCapability) *StripeKYCValidatorMock {
	mock := &StripeKYCValidatorMock{}

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
		string(abi.Methods["getKYCData"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetKYCData == nil {
				return nil, errors.New("getKYCData method not mocked")
			}
			inputs := abi.Methods["getKYCData"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetKYCDataInput{
				AgentId: values[0].(*big.Int),
			}

			result, err := mock.GetKYCData(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getKYCData"].Outputs.Pack(
				result.Verified,
				result.SessionHash,
				result.VerifiedAt,
			)
		},
		string(abi.Methods["isKYCVerified"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.IsKYCVerified == nil {
				return nil, errors.New("isKYCVerified method not mocked")
			}
			inputs := abi.Methods["isKYCVerified"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := IsKYCVerifiedInput{
				AgentId: values[0].(*big.Int),
			}

			result, err := mock.IsKYCVerified(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["isKYCVerified"].Outputs.Pack(result)
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
