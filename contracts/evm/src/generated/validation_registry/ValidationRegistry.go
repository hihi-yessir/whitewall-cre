// Code generated — DO NOT EDIT.

package validation_registry

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
	_ = rpc.API{}
	_ = json.Unmarshal
	_ = reflect.Bool
)

var ValidationRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestURI\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"ValidationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"response\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"responseURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"responseHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"ValidationResponse\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"getAgentValidations\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIdentityRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"validatorAddresses\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"getSummary\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"avgResponse\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"getValidationStatus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"response\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"responseHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"getValidatorRequests\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identityRegistry_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"requestURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"validationRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"response\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"responseURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"responseHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"validationResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Structs

// Contract Method Inputs
type GetAgentValidationsInput struct {
	AgentId *big.Int
}

type GetSummaryInput struct {
	AgentId            *big.Int
	ValidatorAddresses []common.Address
	Tag                string
}

type GetValidationStatusInput struct {
	RequestHash [32]byte
}

type GetValidatorRequestsInput struct {
	ValidatorAddress common.Address
}

type InitializeInput struct {
	IdentityRegistry common.Address
}

type TransferOwnershipInput struct {
	NewOwner common.Address
}

type UpgradeToAndCallInput struct {
	NewImplementation common.Address
	Data              []byte
}

type ValidationRequestInput struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestURI       string
	RequestHash      [32]byte
}

type ValidationResponseInput struct {
	RequestHash  [32]byte
	Response     uint8
	ResponseURI  string
	ResponseHash [32]byte
	Tag          string
}

// Contract Method Outputs
type GetSummaryOutput struct {
	Count       uint64
	AvgResponse uint8
}

type GetValidationStatusOutput struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	Response         uint8
	ResponseHash     [32]byte
	Tag              string
	LastUpdate       *big.Int
}

// Errors
type AddressEmptyCode struct {
	Target common.Address
}

type ERC1967InvalidImplementation struct {
	Implementation common.Address
}

type ERC1967NonPayable struct {
}

type FailedCall struct {
}

type InvalidInitialization struct {
}

type NotInitializing struct {
}

type OwnableInvalidOwner struct {
	Owner common.Address
}

type OwnableUnauthorizedAccount struct {
	Account common.Address
}

type UUPSUnauthorizedCallContext struct {
}

type UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// Events
// The <Event>Topics struct should be used as a filter (for log triggers).
// Note: It is only possible to filter on indexed fields.
// Indexed (string and bytes) fields will be of type common.Hash.
// They need to he (crypto.Keccak256) hashed and passed in.
// Indexed (tuple/slice/array) fields can be passed in as is, the Encode<Event>Topics function will handle the hashing.
//
// The <Event>Decoded struct will be the result of calling decode (Adapt) on the log trigger result.
// Indexed dynamic type fields will be of type common.Hash.

type InitializedTopics struct {
}

type InitializedDecoded struct {
	Version uint64
}

type OwnershipTransferredTopics struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type OwnershipTransferredDecoded struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type UpgradedTopics struct {
	Implementation common.Address
}

type UpgradedDecoded struct {
	Implementation common.Address
}

type ValidationRequestTopics struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestHash      [32]byte
}

type ValidationRequestDecoded struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestURI       string
	RequestHash      [32]byte
}

type ValidationResponseTopics struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestHash      [32]byte
}

type ValidationResponseDecoded struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestHash      [32]byte
	Response         uint8
	ResponseURI      string
	ResponseHash     [32]byte
	Tag              string
}

// Main Binding Type for ValidationRegistry
type ValidationRegistry struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   ValidationRegistryCodec
}

type ValidationRegistryCodec interface {
	EncodeUPGRADEINTERFACEVERSIONMethodCall() ([]byte, error)
	DecodeUPGRADEINTERFACEVERSIONMethodOutput(data []byte) (string, error)
	EncodeGetAgentValidationsMethodCall(in GetAgentValidationsInput) ([]byte, error)
	DecodeGetAgentValidationsMethodOutput(data []byte) ([][32]byte, error)
	EncodeGetIdentityRegistryMethodCall() ([]byte, error)
	DecodeGetIdentityRegistryMethodOutput(data []byte) (common.Address, error)
	EncodeGetSummaryMethodCall(in GetSummaryInput) ([]byte, error)
	DecodeGetSummaryMethodOutput(data []byte) (GetSummaryOutput, error)
	EncodeGetValidationStatusMethodCall(in GetValidationStatusInput) ([]byte, error)
	DecodeGetValidationStatusMethodOutput(data []byte) (GetValidationStatusOutput, error)
	EncodeGetValidatorRequestsMethodCall(in GetValidatorRequestsInput) ([]byte, error)
	DecodeGetValidatorRequestsMethodOutput(data []byte) ([][32]byte, error)
	EncodeGetVersionMethodCall() ([]byte, error)
	DecodeGetVersionMethodOutput(data []byte) (string, error)
	EncodeInitializeMethodCall(in InitializeInput) ([]byte, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodeProxiableUUIDMethodCall() ([]byte, error)
	DecodeProxiableUUIDMethodOutput(data []byte) ([32]byte, error)
	EncodeRenounceOwnershipMethodCall() ([]byte, error)
	EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error)
	EncodeUpgradeToAndCallMethodCall(in UpgradeToAndCallInput) ([]byte, error)
	EncodeValidationRequestMethodCall(in ValidationRequestInput) ([]byte, error)
	EncodeValidationResponseMethodCall(in ValidationResponseInput) ([]byte, error)
	InitializedLogHash() []byte
	EncodeInitializedTopics(evt abi.Event, values []InitializedTopics) ([]*evm.TopicValues, error)
	DecodeInitialized(log *evm.Log) (*InitializedDecoded, error)
	OwnershipTransferredLogHash() []byte
	EncodeOwnershipTransferredTopics(evt abi.Event, values []OwnershipTransferredTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error)
	UpgradedLogHash() []byte
	EncodeUpgradedTopics(evt abi.Event, values []UpgradedTopics) ([]*evm.TopicValues, error)
	DecodeUpgraded(log *evm.Log) (*UpgradedDecoded, error)
	ValidationRequestLogHash() []byte
	EncodeValidationRequestTopics(evt abi.Event, values []ValidationRequestTopics) ([]*evm.TopicValues, error)
	DecodeValidationRequest(log *evm.Log) (*ValidationRequestDecoded, error)
	ValidationResponseLogHash() []byte
	EncodeValidationResponseTopics(evt abi.Event, values []ValidationResponseTopics) ([]*evm.TopicValues, error)
	DecodeValidationResponse(log *evm.Log) (*ValidationResponseDecoded, error)
}

func NewValidationRegistry(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*ValidationRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidationRegistryMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &ValidationRegistry{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type Codec struct {
	abi *abi.ABI
}

func NewCodec() (ValidationRegistryCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidationRegistryMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeUPGRADEINTERFACEVERSIONMethodCall() ([]byte, error) {
	return c.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

func (c *Codec) DecodeUPGRADEINTERFACEVERSIONMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["UPGRADE_INTERFACE_VERSION"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(string), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result string
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(string), fmt.Errorf("failed to unmarshal to string: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetAgentValidationsMethodCall(in GetAgentValidationsInput) ([]byte, error) {
	return c.abi.Pack("getAgentValidations", in.AgentId)
}

func (c *Codec) DecodeGetAgentValidationsMethodOutput(data []byte) ([][32]byte, error) {
	vals, err := c.abi.Methods["getAgentValidations"].Outputs.Unpack(data)
	if err != nil {
		return *new([][32]byte), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([][32]byte), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result [][32]byte
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([][32]byte), fmt.Errorf("failed to unmarshal to [][32]byte: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetIdentityRegistryMethodCall() ([]byte, error) {
	return c.abi.Pack("getIdentityRegistry")
}

func (c *Codec) DecodeGetIdentityRegistryMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["getIdentityRegistry"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetSummaryMethodCall(in GetSummaryInput) ([]byte, error) {
	return c.abi.Pack("getSummary", in.AgentId, in.ValidatorAddresses, in.Tag)
}

func (c *Codec) DecodeGetSummaryMethodOutput(data []byte) (GetSummaryOutput, error) {
	vals, err := c.abi.Methods["getSummary"].Outputs.Unpack(data)
	if err != nil {
		return GetSummaryOutput{}, err
	}
	if len(vals) != 2 {
		return GetSummaryOutput{}, fmt.Errorf("expected 2 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return GetSummaryOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 uint64
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return GetSummaryOutput{}, fmt.Errorf("failed to unmarshal to uint64: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return GetSummaryOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 uint8
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return GetSummaryOutput{}, fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}

	return GetSummaryOutput{
		Count:       result0,
		AvgResponse: result1,
	}, nil
}

func (c *Codec) EncodeGetValidationStatusMethodCall(in GetValidationStatusInput) ([]byte, error) {
	return c.abi.Pack("getValidationStatus", in.RequestHash)
}

func (c *Codec) DecodeGetValidationStatusMethodOutput(data []byte) (GetValidationStatusOutput, error) {
	vals, err := c.abi.Methods["getValidationStatus"].Outputs.Unpack(data)
	if err != nil {
		return GetValidationStatusOutput{}, err
	}
	if len(vals) != 6 {
		return GetValidationStatusOutput{}, fmt.Errorf("expected 6 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 common.Address
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 *big.Int
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData2, err := json.Marshal(vals[2])
	if err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to marshal ABI result 2: %w", err)
	}

	var result2 uint8
	if err := json.Unmarshal(jsonData2, &result2); err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}
	jsonData3, err := json.Marshal(vals[3])
	if err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to marshal ABI result 3: %w", err)
	}

	var result3 [32]byte
	if err := json.Unmarshal(jsonData3, &result3); err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to unmarshal to [32]byte: %w", err)
	}
	jsonData4, err := json.Marshal(vals[4])
	if err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to marshal ABI result 4: %w", err)
	}

	var result4 string
	if err := json.Unmarshal(jsonData4, &result4); err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to unmarshal to string: %w", err)
	}
	jsonData5, err := json.Marshal(vals[5])
	if err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to marshal ABI result 5: %w", err)
	}

	var result5 *big.Int
	if err := json.Unmarshal(jsonData5, &result5); err != nil {
		return GetValidationStatusOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return GetValidationStatusOutput{
		ValidatorAddress: result0,
		AgentId:          result1,
		Response:         result2,
		ResponseHash:     result3,
		Tag:              result4,
		LastUpdate:       result5,
	}, nil
}

func (c *Codec) EncodeGetValidatorRequestsMethodCall(in GetValidatorRequestsInput) ([]byte, error) {
	return c.abi.Pack("getValidatorRequests", in.ValidatorAddress)
}

func (c *Codec) DecodeGetValidatorRequestsMethodOutput(data []byte) ([][32]byte, error) {
	vals, err := c.abi.Methods["getValidatorRequests"].Outputs.Unpack(data)
	if err != nil {
		return *new([][32]byte), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([][32]byte), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result [][32]byte
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([][32]byte), fmt.Errorf("failed to unmarshal to [][32]byte: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetVersionMethodCall() ([]byte, error) {
	return c.abi.Pack("getVersion")
}

func (c *Codec) DecodeGetVersionMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["getVersion"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(string), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result string
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(string), fmt.Errorf("failed to unmarshal to string: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeInitializeMethodCall(in InitializeInput) ([]byte, error) {
	return c.abi.Pack("initialize", in.IdentityRegistry)
}

func (c *Codec) EncodeOwnerMethodCall() ([]byte, error) {
	return c.abi.Pack("owner")
}

func (c *Codec) DecodeOwnerMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["owner"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeProxiableUUIDMethodCall() ([]byte, error) {
	return c.abi.Pack("proxiableUUID")
}

func (c *Codec) DecodeProxiableUUIDMethodOutput(data []byte) ([32]byte, error) {
	vals, err := c.abi.Methods["proxiableUUID"].Outputs.Unpack(data)
	if err != nil {
		return *new([32]byte), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([32]byte), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result [32]byte
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([32]byte), fmt.Errorf("failed to unmarshal to [32]byte: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeRenounceOwnershipMethodCall() ([]byte, error) {
	return c.abi.Pack("renounceOwnership")
}

func (c *Codec) EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error) {
	return c.abi.Pack("transferOwnership", in.NewOwner)
}

func (c *Codec) EncodeUpgradeToAndCallMethodCall(in UpgradeToAndCallInput) ([]byte, error) {
	return c.abi.Pack("upgradeToAndCall", in.NewImplementation, in.Data)
}

func (c *Codec) EncodeValidationRequestMethodCall(in ValidationRequestInput) ([]byte, error) {
	return c.abi.Pack("validationRequest", in.ValidatorAddress, in.AgentId, in.RequestURI, in.RequestHash)
}

func (c *Codec) EncodeValidationResponseMethodCall(in ValidationResponseInput) ([]byte, error) {
	return c.abi.Pack("validationResponse", in.RequestHash, in.Response, in.ResponseURI, in.ResponseHash, in.Tag)
}

func (c *Codec) InitializedLogHash() []byte {
	return c.abi.Events["Initialized"].ID.Bytes()
}

func (c *Codec) EncodeInitializedTopics(
	evt abi.Event,
	values []InitializedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeInitialized decodes a log into a Initialized struct.
func (c *Codec) DecodeInitialized(log *evm.Log) (*InitializedDecoded, error) {
	event := new(InitializedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Initialized", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Initialized"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) OwnershipTransferredLogHash() []byte {
	return c.abi.Events["OwnershipTransferred"].ID.Bytes()
}

func (c *Codec) EncodeOwnershipTransferredTopics(
	evt abi.Event,
	values []OwnershipTransferredTopics,
) ([]*evm.TopicValues, error) {
	var previousOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.PreviousOwner).IsZero() {
			previousOwnerRule = append(previousOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.PreviousOwner)
		if err != nil {
			return nil, err
		}
		previousOwnerRule = append(previousOwnerRule, fieldVal)
	}
	var newOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.NewOwner).IsZero() {
			newOwnerRule = append(newOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.NewOwner)
		if err != nil {
			return nil, err
		}
		newOwnerRule = append(newOwnerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		previousOwnerRule,
		newOwnerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeOwnershipTransferred decodes a log into a OwnershipTransferred struct.
func (c *Codec) DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error) {
	event := new(OwnershipTransferredDecoded)
	if err := c.abi.UnpackIntoInterface(event, "OwnershipTransferred", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["OwnershipTransferred"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) UpgradedLogHash() []byte {
	return c.abi.Events["Upgraded"].ID.Bytes()
}

func (c *Codec) EncodeUpgradedTopics(
	evt abi.Event,
	values []UpgradedTopics,
) ([]*evm.TopicValues, error) {
	var implementationRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Implementation).IsZero() {
			implementationRule = append(implementationRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Implementation)
		if err != nil {
			return nil, err
		}
		implementationRule = append(implementationRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		implementationRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeUpgraded decodes a log into a Upgraded struct.
func (c *Codec) DecodeUpgraded(log *evm.Log) (*UpgradedDecoded, error) {
	event := new(UpgradedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Upgraded", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Upgraded"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) ValidationRequestLogHash() []byte {
	return c.abi.Events["ValidationRequest"].ID.Bytes()
}

func (c *Codec) EncodeValidationRequestTopics(
	evt abi.Event,
	values []ValidationRequestTopics,
) ([]*evm.TopicValues, error) {
	var validatorAddressRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.ValidatorAddress).IsZero() {
			validatorAddressRule = append(validatorAddressRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.ValidatorAddress)
		if err != nil {
			return nil, err
		}
		validatorAddressRule = append(validatorAddressRule, fieldVal)
	}
	var agentIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.AgentId).IsZero() {
			agentIdRule = append(agentIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.AgentId)
		if err != nil {
			return nil, err
		}
		agentIdRule = append(agentIdRule, fieldVal)
	}
	var requestHashRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.RequestHash).IsZero() {
			requestHashRule = append(requestHashRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[3], v.RequestHash)
		if err != nil {
			return nil, err
		}
		requestHashRule = append(requestHashRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		validatorAddressRule,
		agentIdRule,
		requestHashRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeValidationRequest decodes a log into a ValidationRequest struct.
func (c *Codec) DecodeValidationRequest(log *evm.Log) (*ValidationRequestDecoded, error) {
	event := new(ValidationRequestDecoded)
	if err := c.abi.UnpackIntoInterface(event, "ValidationRequest", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["ValidationRequest"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) ValidationResponseLogHash() []byte {
	return c.abi.Events["ValidationResponse"].ID.Bytes()
}

func (c *Codec) EncodeValidationResponseTopics(
	evt abi.Event,
	values []ValidationResponseTopics,
) ([]*evm.TopicValues, error) {
	var validatorAddressRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.ValidatorAddress).IsZero() {
			validatorAddressRule = append(validatorAddressRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.ValidatorAddress)
		if err != nil {
			return nil, err
		}
		validatorAddressRule = append(validatorAddressRule, fieldVal)
	}
	var agentIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.AgentId).IsZero() {
			agentIdRule = append(agentIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.AgentId)
		if err != nil {
			return nil, err
		}
		agentIdRule = append(agentIdRule, fieldVal)
	}
	var requestHashRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.RequestHash).IsZero() {
			requestHashRule = append(requestHashRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[2], v.RequestHash)
		if err != nil {
			return nil, err
		}
		requestHashRule = append(requestHashRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		validatorAddressRule,
		agentIdRule,
		requestHashRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeValidationResponse decodes a log into a ValidationResponse struct.
func (c *Codec) DecodeValidationResponse(log *evm.Log) (*ValidationResponseDecoded, error) {
	event := new(ValidationResponseDecoded)
	if err := c.abi.UnpackIntoInterface(event, "ValidationResponse", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["ValidationResponse"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c ValidationRegistry) UPGRADEINTERFACEVERSION(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[string] {
	calldata, err := c.Codec.EncodeUPGRADEINTERFACEVERSIONMethodCall()
	if err != nil {
		return cre.PromiseFromResult[string](*new(string), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (string, error) {
		return c.Codec.DecodeUPGRADEINTERFACEVERSIONMethodOutput(response.Data)
	})

}

func (c ValidationRegistry) GetAgentValidations(
	runtime cre.Runtime,
	args GetAgentValidationsInput,
	blockNumber *big.Int,
) cre.Promise[[][32]byte] {
	calldata, err := c.Codec.EncodeGetAgentValidationsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[[][32]byte](*new([][32]byte), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([][32]byte, error) {
		return c.Codec.DecodeGetAgentValidationsMethodOutput(response.Data)
	})

}

func (c ValidationRegistry) GetIdentityRegistry(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeGetIdentityRegistryMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeGetIdentityRegistryMethodOutput(response.Data)
	})

}

func (c ValidationRegistry) GetSummary(
	runtime cre.Runtime,
	args GetSummaryInput,
	blockNumber *big.Int,
) cre.Promise[GetSummaryOutput] {
	calldata, err := c.Codec.EncodeGetSummaryMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[GetSummaryOutput](GetSummaryOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (GetSummaryOutput, error) {
		return c.Codec.DecodeGetSummaryMethodOutput(response.Data)
	})

}

func (c ValidationRegistry) GetValidationStatus(
	runtime cre.Runtime,
	args GetValidationStatusInput,
	blockNumber *big.Int,
) cre.Promise[GetValidationStatusOutput] {
	calldata, err := c.Codec.EncodeGetValidationStatusMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[GetValidationStatusOutput](GetValidationStatusOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (GetValidationStatusOutput, error) {
		return c.Codec.DecodeGetValidationStatusMethodOutput(response.Data)
	})

}

func (c ValidationRegistry) GetValidatorRequests(
	runtime cre.Runtime,
	args GetValidatorRequestsInput,
	blockNumber *big.Int,
) cre.Promise[[][32]byte] {
	calldata, err := c.Codec.EncodeGetValidatorRequestsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[[][32]byte](*new([][32]byte), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([][32]byte, error) {
		return c.Codec.DecodeGetValidatorRequestsMethodOutput(response.Data)
	})

}

func (c ValidationRegistry) Owner(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOwnerMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeOwnerMethodOutput(response.Data)
	})

}

func (c ValidationRegistry) ProxiableUUID(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[[32]byte] {
	calldata, err := c.Codec.EncodeProxiableUUIDMethodCall()
	if err != nil {
		return cre.PromiseFromResult[[32]byte](*new([32]byte), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([32]byte, error) {
		return c.Codec.DecodeProxiableUUIDMethodOutput(response.Data)
	})

}

func (c ValidationRegistry) WriteReport(
	runtime cre.Runtime,
	report *cre.Report,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  c.Address.Bytes(),
		Report:    report,
		GasConfig: gasConfig,
	})
}

// DecodeAddressEmptyCodeError decodes a AddressEmptyCode error from revert data.
func (c *ValidationRegistry) DecodeAddressEmptyCodeError(data []byte) (*AddressEmptyCode, error) {
	args := c.ABI.Errors["AddressEmptyCode"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	target, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for target in AddressEmptyCode error")
	}

	return &AddressEmptyCode{
		Target: target,
	}, nil
}

// Error implements the error interface for AddressEmptyCode.
func (e *AddressEmptyCode) Error() string {
	return fmt.Sprintf("AddressEmptyCode error: target=%v;", e.Target)
}

// DecodeERC1967InvalidImplementationError decodes a ERC1967InvalidImplementation error from revert data.
func (c *ValidationRegistry) DecodeERC1967InvalidImplementationError(data []byte) (*ERC1967InvalidImplementation, error) {
	args := c.ABI.Errors["ERC1967InvalidImplementation"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	implementation, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for implementation in ERC1967InvalidImplementation error")
	}

	return &ERC1967InvalidImplementation{
		Implementation: implementation,
	}, nil
}

// Error implements the error interface for ERC1967InvalidImplementation.
func (e *ERC1967InvalidImplementation) Error() string {
	return fmt.Sprintf("ERC1967InvalidImplementation error: implementation=%v;", e.Implementation)
}

// DecodeERC1967NonPayableError decodes a ERC1967NonPayable error from revert data.
func (c *ValidationRegistry) DecodeERC1967NonPayableError(data []byte) (*ERC1967NonPayable, error) {
	args := c.ABI.Errors["ERC1967NonPayable"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &ERC1967NonPayable{}, nil
}

// Error implements the error interface for ERC1967NonPayable.
func (e *ERC1967NonPayable) Error() string {
	return fmt.Sprintf("ERC1967NonPayable error:")
}

// DecodeFailedCallError decodes a FailedCall error from revert data.
func (c *ValidationRegistry) DecodeFailedCallError(data []byte) (*FailedCall, error) {
	args := c.ABI.Errors["FailedCall"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &FailedCall{}, nil
}

// Error implements the error interface for FailedCall.
func (e *FailedCall) Error() string {
	return fmt.Sprintf("FailedCall error:")
}

// DecodeInvalidInitializationError decodes a InvalidInitialization error from revert data.
func (c *ValidationRegistry) DecodeInvalidInitializationError(data []byte) (*InvalidInitialization, error) {
	args := c.ABI.Errors["InvalidInitialization"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &InvalidInitialization{}, nil
}

// Error implements the error interface for InvalidInitialization.
func (e *InvalidInitialization) Error() string {
	return fmt.Sprintf("InvalidInitialization error:")
}

// DecodeNotInitializingError decodes a NotInitializing error from revert data.
func (c *ValidationRegistry) DecodeNotInitializingError(data []byte) (*NotInitializing, error) {
	args := c.ABI.Errors["NotInitializing"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &NotInitializing{}, nil
}

// Error implements the error interface for NotInitializing.
func (e *NotInitializing) Error() string {
	return fmt.Sprintf("NotInitializing error:")
}

// DecodeOwnableInvalidOwnerError decodes a OwnableInvalidOwner error from revert data.
func (c *ValidationRegistry) DecodeOwnableInvalidOwnerError(data []byte) (*OwnableInvalidOwner, error) {
	args := c.ABI.Errors["OwnableInvalidOwner"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	owner, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for owner in OwnableInvalidOwner error")
	}

	return &OwnableInvalidOwner{
		Owner: owner,
	}, nil
}

// Error implements the error interface for OwnableInvalidOwner.
func (e *OwnableInvalidOwner) Error() string {
	return fmt.Sprintf("OwnableInvalidOwner error: owner=%v;", e.Owner)
}

// DecodeOwnableUnauthorizedAccountError decodes a OwnableUnauthorizedAccount error from revert data.
func (c *ValidationRegistry) DecodeOwnableUnauthorizedAccountError(data []byte) (*OwnableUnauthorizedAccount, error) {
	args := c.ABI.Errors["OwnableUnauthorizedAccount"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	account, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for account in OwnableUnauthorizedAccount error")
	}

	return &OwnableUnauthorizedAccount{
		Account: account,
	}, nil
}

// Error implements the error interface for OwnableUnauthorizedAccount.
func (e *OwnableUnauthorizedAccount) Error() string {
	return fmt.Sprintf("OwnableUnauthorizedAccount error: account=%v;", e.Account)
}

// DecodeUUPSUnauthorizedCallContextError decodes a UUPSUnauthorizedCallContext error from revert data.
func (c *ValidationRegistry) DecodeUUPSUnauthorizedCallContextError(data []byte) (*UUPSUnauthorizedCallContext, error) {
	args := c.ABI.Errors["UUPSUnauthorizedCallContext"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &UUPSUnauthorizedCallContext{}, nil
}

// Error implements the error interface for UUPSUnauthorizedCallContext.
func (e *UUPSUnauthorizedCallContext) Error() string {
	return fmt.Sprintf("UUPSUnauthorizedCallContext error:")
}

// DecodeUUPSUnsupportedProxiableUUIDError decodes a UUPSUnsupportedProxiableUUID error from revert data.
func (c *ValidationRegistry) DecodeUUPSUnsupportedProxiableUUIDError(data []byte) (*UUPSUnsupportedProxiableUUID, error) {
	args := c.ABI.Errors["UUPSUnsupportedProxiableUUID"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	slot, ok0 := values[0].([32]byte)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for slot in UUPSUnsupportedProxiableUUID error")
	}

	return &UUPSUnsupportedProxiableUUID{
		Slot: slot,
	}, nil
}

// Error implements the error interface for UUPSUnsupportedProxiableUUID.
func (e *UUPSUnsupportedProxiableUUID) Error() string {
	return fmt.Sprintf("UUPSUnsupportedProxiableUUID error: slot=%v;", e.Slot)
}

func (c *ValidationRegistry) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["AddressEmptyCode"].ID.Bytes()[:4]):
		return c.DecodeAddressEmptyCodeError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]):
		return c.DecodeERC1967InvalidImplementationError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC1967NonPayable"].ID.Bytes()[:4]):
		return c.DecodeERC1967NonPayableError(data)
	case common.Bytes2Hex(c.ABI.Errors["FailedCall"].ID.Bytes()[:4]):
		return c.DecodeFailedCallError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidInitialization"].ID.Bytes()[:4]):
		return c.DecodeInvalidInitializationError(data)
	case common.Bytes2Hex(c.ABI.Errors["NotInitializing"].ID.Bytes()[:4]):
		return c.DecodeNotInitializingError(data)
	case common.Bytes2Hex(c.ABI.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]):
		return c.DecodeOwnableInvalidOwnerError(data)
	case common.Bytes2Hex(c.ABI.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]):
		return c.DecodeOwnableUnauthorizedAccountError(data)
	case common.Bytes2Hex(c.ABI.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]):
		return c.DecodeUUPSUnauthorizedCallContextError(data)
	case common.Bytes2Hex(c.ABI.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]):
		return c.DecodeUUPSUnsupportedProxiableUUIDError(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

// InitializedTrigger wraps the raw log trigger and provides decoded InitializedDecoded data
type InitializedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ValidationRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into Initialized data
func (t *InitializedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[InitializedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeInitialized(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Initialized log: %w", err)
	}

	return &bindings.DecodedLog[InitializedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ValidationRegistry) LogTriggerInitializedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []InitializedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[InitializedDecoded]], error) {
	event := c.ABI.Events["Initialized"]
	topics, err := c.Codec.EncodeInitializedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Initialized: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &InitializedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ValidationRegistry) FilterLogsInitialized(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.InitializedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OwnershipTransferredTrigger wraps the raw log trigger and provides decoded OwnershipTransferredDecoded data
type OwnershipTransferredTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ValidationRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into OwnershipTransferred data
func (t *OwnershipTransferredTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[OwnershipTransferredDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeOwnershipTransferred(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OwnershipTransferred log: %w", err)
	}

	return &bindings.DecodedLog[OwnershipTransferredDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ValidationRegistry) LogTriggerOwnershipTransferredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipTransferredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipTransferredDecoded]], error) {
	event := c.ABI.Events["OwnershipTransferred"]
	topics, err := c.Codec.EncodeOwnershipTransferredTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for OwnershipTransferred: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &OwnershipTransferredTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ValidationRegistry) FilterLogsOwnershipTransferred(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.OwnershipTransferredLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// UpgradedTrigger wraps the raw log trigger and provides decoded UpgradedDecoded data
type UpgradedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ValidationRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into Upgraded data
func (t *UpgradedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[UpgradedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeUpgraded(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Upgraded log: %w", err)
	}

	return &bindings.DecodedLog[UpgradedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ValidationRegistry) LogTriggerUpgradedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []UpgradedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[UpgradedDecoded]], error) {
	event := c.ABI.Events["Upgraded"]
	topics, err := c.Codec.EncodeUpgradedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Upgraded: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &UpgradedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ValidationRegistry) FilterLogsUpgraded(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.UpgradedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// ValidationRequestTrigger wraps the raw log trigger and provides decoded ValidationRequestDecoded data
type ValidationRequestTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ValidationRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into ValidationRequest data
func (t *ValidationRequestTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[ValidationRequestDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeValidationRequest(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ValidationRequest log: %w", err)
	}

	return &bindings.DecodedLog[ValidationRequestDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ValidationRegistry) LogTriggerValidationRequestLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []ValidationRequestTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[ValidationRequestDecoded]], error) {
	event := c.ABI.Events["ValidationRequest"]
	topics, err := c.Codec.EncodeValidationRequestTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for ValidationRequest: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &ValidationRequestTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ValidationRegistry) FilterLogsValidationRequest(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.ValidationRequestLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// ValidationResponseTrigger wraps the raw log trigger and provides decoded ValidationResponseDecoded data
type ValidationResponseTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ValidationRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into ValidationResponse data
func (t *ValidationResponseTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[ValidationResponseDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeValidationResponse(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ValidationResponse log: %w", err)
	}

	return &bindings.DecodedLog[ValidationResponseDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ValidationRegistry) LogTriggerValidationResponseLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []ValidationResponseTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[ValidationResponseDecoded]], error) {
	event := c.ABI.Events["ValidationResponse"]
	topics, err := c.Codec.EncodeValidationResponseTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for ValidationResponse: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &ValidationResponseTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ValidationRegistry) FilterLogsValidationResponse(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.ValidationResponseLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}
