// Code generated — DO NOT EDIT.

//go:build !wasip1

package identity_registry

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

// IdentityRegistryMock is a mock implementation of IdentityRegistry for testing.
type IdentityRegistryMock struct {
	UPGRADEINTERFACEVERSION func() (string, error)
	BalanceOf               func(BalanceOfInput) (*big.Int, error)
	Eip712Domain            func() (Eip712DomainOutput, error)
	GetAgentWallet          func(GetAgentWalletInput) (common.Address, error)
	GetApproved             func(GetApprovedInput) (common.Address, error)
	GetMetadata             func(GetMetadataInput) ([]byte, error)
	IsApprovedForAll        func(IsApprovedForAllInput) (bool, error)
	Name                    func() (string, error)
	Owner                   func() (common.Address, error)
	OwnerOf                 func(OwnerOfInput) (common.Address, error)
	ProxiableUUID           func() ([32]byte, error)
	SupportsInterface       func(SupportsInterfaceInput) (bool, error)
	Symbol                  func() (string, error)
	TokenURI                func(TokenURIInput) (string, error)
}

// NewIdentityRegistryMock creates a new IdentityRegistryMock for testing.
func NewIdentityRegistryMock(address common.Address, clientMock *evmmock.ClientCapability) *IdentityRegistryMock {
	mock := &IdentityRegistryMock{}

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
		string(abi.Methods["balanceOf"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.BalanceOf == nil {
				return nil, errors.New("balanceOf method not mocked")
			}
			inputs := abi.Methods["balanceOf"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := BalanceOfInput{
				Owner: values[0].(common.Address),
			}

			result, err := mock.BalanceOf(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["balanceOf"].Outputs.Pack(result)
		},
		string(abi.Methods["eip712Domain"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Eip712Domain == nil {
				return nil, errors.New("eip712Domain method not mocked")
			}
			result, err := mock.Eip712Domain()
			if err != nil {
				return nil, err
			}
			return abi.Methods["eip712Domain"].Outputs.Pack(
				result.Fields,
				result.Name,
				result.Version,
				result.ChainId,
				result.VerifyingContract,
				result.Salt,
				result.Extensions,
			)
		},
		string(abi.Methods["getAgentWallet"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetAgentWallet == nil {
				return nil, errors.New("getAgentWallet method not mocked")
			}
			inputs := abi.Methods["getAgentWallet"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetAgentWalletInput{
				AgentId: values[0].(*big.Int),
			}

			result, err := mock.GetAgentWallet(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getAgentWallet"].Outputs.Pack(result)
		},
		string(abi.Methods["getApproved"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetApproved == nil {
				return nil, errors.New("getApproved method not mocked")
			}
			inputs := abi.Methods["getApproved"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetApprovedInput{
				TokenId: values[0].(*big.Int),
			}

			result, err := mock.GetApproved(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getApproved"].Outputs.Pack(result)
		},
		string(abi.Methods["getMetadata"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetMetadata == nil {
				return nil, errors.New("getMetadata method not mocked")
			}
			inputs := abi.Methods["getMetadata"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 2 {
				return nil, errors.New("expected 2 input values")
			}

			args := GetMetadataInput{
				AgentId:     values[0].(*big.Int),
				MetadataKey: values[1].(string),
			}

			result, err := mock.GetMetadata(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getMetadata"].Outputs.Pack(result)
		},
		string(abi.Methods["isApprovedForAll"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.IsApprovedForAll == nil {
				return nil, errors.New("isApprovedForAll method not mocked")
			}
			inputs := abi.Methods["isApprovedForAll"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 2 {
				return nil, errors.New("expected 2 input values")
			}

			args := IsApprovedForAllInput{
				Owner:    values[0].(common.Address),
				Operator: values[1].(common.Address),
			}

			result, err := mock.IsApprovedForAll(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["isApprovedForAll"].Outputs.Pack(result)
		},
		string(abi.Methods["name"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Name == nil {
				return nil, errors.New("name method not mocked")
			}
			result, err := mock.Name()
			if err != nil {
				return nil, err
			}
			return abi.Methods["name"].Outputs.Pack(result)
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
		string(abi.Methods["ownerOf"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.OwnerOf == nil {
				return nil, errors.New("ownerOf method not mocked")
			}
			inputs := abi.Methods["ownerOf"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := OwnerOfInput{
				TokenId: values[0].(*big.Int),
			}

			result, err := mock.OwnerOf(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["ownerOf"].Outputs.Pack(result)
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
		string(abi.Methods["supportsInterface"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.SupportsInterface == nil {
				return nil, errors.New("supportsInterface method not mocked")
			}
			inputs := abi.Methods["supportsInterface"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := SupportsInterfaceInput{
				InterfaceId: values[0].([4]byte),
			}

			result, err := mock.SupportsInterface(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["supportsInterface"].Outputs.Pack(result)
		},
		string(abi.Methods["symbol"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Symbol == nil {
				return nil, errors.New("symbol method not mocked")
			}
			result, err := mock.Symbol()
			if err != nil {
				return nil, err
			}
			return abi.Methods["symbol"].Outputs.Pack(result)
		},
		string(abi.Methods["tokenURI"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TokenURI == nil {
				return nil, errors.New("tokenURI method not mocked")
			}
			inputs := abi.Methods["tokenURI"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := TokenURIInput{
				TokenId: values[0].(*big.Int),
			}

			result, err := mock.TokenURI(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["tokenURI"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
