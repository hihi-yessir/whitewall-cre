// Code generated — DO NOT EDIT.

package identity_registry

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

var IdentityRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedMetadataKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadataValue\",\"type\":\"bytes\"}],\"name\":\"MetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"agentURI\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updatedBy\",\"type\":\"address\"}],\"name\":\"URIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"getAgentWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataKey\",\"type\":\"string\"}],\"name\":\"getMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"agentURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"metadataKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"metadataValue\",\"type\":\"bytes\"}],\"internalType\":\"structIdentityRegistryUpgradeable.MetadataEntry[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"agentURI\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"setAgentURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"setAgentWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"metadataValue\",\"type\":\"bytes\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"unsetAgentWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Structs
type UpgradeableMetadataEntry struct {
	MetadataKey   string
	MetadataValue []byte
}

// Contract Method Inputs
type ApproveInput struct {
	To      common.Address
	TokenId *big.Int
}

type BalanceOfInput struct {
	Owner common.Address
}

type GetAgentWalletInput struct {
	AgentId *big.Int
}

type GetApprovedInput struct {
	TokenId *big.Int
}

type GetMetadataInput struct {
	AgentId     *big.Int
	MetadataKey string
}

type IsApprovedForAllInput struct {
	Owner    common.Address
	Operator common.Address
}

type OwnerOfInput struct {
	TokenId *big.Int
}

type Register0Input struct {
	AgentURI string
	Metadata []UpgradeableMetadataEntry
}

type Register1Input struct {
	AgentURI string
}

type SafeTransferFromInput struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

type SafeTransferFrom0Input struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Data    []byte
}

type SetAgentURIInput struct {
	AgentId *big.Int
	NewURI  string
}

type SetAgentWalletInput struct {
	AgentId   *big.Int
	NewWallet common.Address
	Deadline  *big.Int
	Signature []byte
}

type SetApprovalForAllInput struct {
	Operator common.Address
	Approved bool
}

type SetMetadataInput struct {
	AgentId       *big.Int
	MetadataKey   string
	MetadataValue []byte
}

type SupportsInterfaceInput struct {
	InterfaceId [4]byte
}

type TokenURIInput struct {
	TokenId *big.Int
}

type TransferFromInput struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

type TransferOwnershipInput struct {
	NewOwner common.Address
}

type UnsetAgentWalletInput struct {
	AgentId *big.Int
}

type UpgradeToAndCallInput struct {
	NewImplementation common.Address
	Data              []byte
}

// Contract Method Outputs
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

// Errors
type AddressEmptyCode struct {
	Target common.Address
}

type ECDSAInvalidSignature struct {
}

type ECDSAInvalidSignatureLength struct {
	Length *big.Int
}

type ECDSAInvalidSignatureS struct {
	S [32]byte
}

type ERC1967InvalidImplementation struct {
	Implementation common.Address
}

type ERC1967NonPayable struct {
}

type ERC721IncorrectOwner struct {
	Sender  common.Address
	TokenId *big.Int
	Owner   common.Address
}

type ERC721InsufficientApproval struct {
	Operator common.Address
	TokenId  *big.Int
}

type ERC721InvalidApprover struct {
	Approver common.Address
}

type ERC721InvalidOperator struct {
	Operator common.Address
}

type ERC721InvalidOwner struct {
	Owner common.Address
}

type ERC721InvalidReceiver struct {
	Receiver common.Address
}

type ERC721InvalidSender struct {
	Sender common.Address
}

type ERC721NonexistentToken struct {
	TokenId *big.Int
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

type ApprovalTopics struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
}

type ApprovalDecoded struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
}

type ApprovalForAllTopics struct {
	Owner    common.Address
	Operator common.Address
}

type ApprovalForAllDecoded struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
}

type BatchMetadataUpdateTopics struct {
}

type BatchMetadataUpdateDecoded struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
}

type EIP712DomainChangedTopics struct {
}

type EIP712DomainChangedDecoded struct {
}

type InitializedTopics struct {
}

type InitializedDecoded struct {
	Version uint64
}

type MetadataSetTopics struct {
	AgentId            *big.Int
	IndexedMetadataKey common.Hash
}

type MetadataSetDecoded struct {
	AgentId            *big.Int
	IndexedMetadataKey common.Hash
	MetadataKey        string
	MetadataValue      []byte
}

type MetadataUpdateTopics struct {
}

type MetadataUpdateDecoded struct {
	TokenId *big.Int
}

type OwnershipTransferredTopics struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type OwnershipTransferredDecoded struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type RegisteredTopics struct {
	AgentId *big.Int
	Owner   common.Address
}

type RegisteredDecoded struct {
	AgentId  *big.Int
	AgentURI string
	Owner    common.Address
}

type TransferTopics struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

type TransferDecoded struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

type URIUpdatedTopics struct {
	AgentId   *big.Int
	UpdatedBy common.Address
}

type URIUpdatedDecoded struct {
	AgentId   *big.Int
	NewURI    string
	UpdatedBy common.Address
}

type UpgradedTopics struct {
	Implementation common.Address
}

type UpgradedDecoded struct {
	Implementation common.Address
}

// Main Binding Type for IdentityRegistry
type IdentityRegistry struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   IdentityRegistryCodec
}

type IdentityRegistryCodec interface {
	EncodeUPGRADEINTERFACEVERSIONMethodCall() ([]byte, error)
	DecodeUPGRADEINTERFACEVERSIONMethodOutput(data []byte) (string, error)
	EncodeApproveMethodCall(in ApproveInput) ([]byte, error)
	EncodeBalanceOfMethodCall(in BalanceOfInput) ([]byte, error)
	DecodeBalanceOfMethodOutput(data []byte) (*big.Int, error)
	EncodeEip712DomainMethodCall() ([]byte, error)
	DecodeEip712DomainMethodOutput(data []byte) (Eip712DomainOutput, error)
	EncodeGetAgentWalletMethodCall(in GetAgentWalletInput) ([]byte, error)
	DecodeGetAgentWalletMethodOutput(data []byte) (common.Address, error)
	EncodeGetApprovedMethodCall(in GetApprovedInput) ([]byte, error)
	DecodeGetApprovedMethodOutput(data []byte) (common.Address, error)
	EncodeGetMetadataMethodCall(in GetMetadataInput) ([]byte, error)
	DecodeGetMetadataMethodOutput(data []byte) ([]byte, error)
	EncodeGetVersionMethodCall() ([]byte, error)
	DecodeGetVersionMethodOutput(data []byte) (string, error)
	EncodeInitializeMethodCall() ([]byte, error)
	EncodeIsApprovedForAllMethodCall(in IsApprovedForAllInput) ([]byte, error)
	DecodeIsApprovedForAllMethodOutput(data []byte) (bool, error)
	EncodeNameMethodCall() ([]byte, error)
	DecodeNameMethodOutput(data []byte) (string, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodeOwnerOfMethodCall(in OwnerOfInput) ([]byte, error)
	DecodeOwnerOfMethodOutput(data []byte) (common.Address, error)
	EncodeProxiableUUIDMethodCall() ([]byte, error)
	DecodeProxiableUUIDMethodOutput(data []byte) ([32]byte, error)
	EncodeRegisterMethodCall() ([]byte, error)
	DecodeRegisterMethodOutput(data []byte) (*big.Int, error)
	EncodeRegister0MethodCall(in Register0Input) ([]byte, error)
	DecodeRegister0MethodOutput(data []byte) (*big.Int, error)
	EncodeRegister1MethodCall(in Register1Input) ([]byte, error)
	DecodeRegister1MethodOutput(data []byte) (*big.Int, error)
	EncodeRenounceOwnershipMethodCall() ([]byte, error)
	EncodeSafeTransferFromMethodCall(in SafeTransferFromInput) ([]byte, error)
	EncodeSafeTransferFrom0MethodCall(in SafeTransferFrom0Input) ([]byte, error)
	EncodeSetAgentURIMethodCall(in SetAgentURIInput) ([]byte, error)
	EncodeSetAgentWalletMethodCall(in SetAgentWalletInput) ([]byte, error)
	EncodeSetApprovalForAllMethodCall(in SetApprovalForAllInput) ([]byte, error)
	EncodeSetMetadataMethodCall(in SetMetadataInput) ([]byte, error)
	EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error)
	DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error)
	EncodeSymbolMethodCall() ([]byte, error)
	DecodeSymbolMethodOutput(data []byte) (string, error)
	EncodeTokenURIMethodCall(in TokenURIInput) ([]byte, error)
	DecodeTokenURIMethodOutput(data []byte) (string, error)
	EncodeTransferFromMethodCall(in TransferFromInput) ([]byte, error)
	EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error)
	EncodeUnsetAgentWalletMethodCall(in UnsetAgentWalletInput) ([]byte, error)
	EncodeUpgradeToAndCallMethodCall(in UpgradeToAndCallInput) ([]byte, error)
	EncodeUpgradeableMetadataEntryStruct(in UpgradeableMetadataEntry) ([]byte, error)
	ApprovalLogHash() []byte
	EncodeApprovalTopics(evt abi.Event, values []ApprovalTopics) ([]*evm.TopicValues, error)
	DecodeApproval(log *evm.Log) (*ApprovalDecoded, error)
	ApprovalForAllLogHash() []byte
	EncodeApprovalForAllTopics(evt abi.Event, values []ApprovalForAllTopics) ([]*evm.TopicValues, error)
	DecodeApprovalForAll(log *evm.Log) (*ApprovalForAllDecoded, error)
	BatchMetadataUpdateLogHash() []byte
	EncodeBatchMetadataUpdateTopics(evt abi.Event, values []BatchMetadataUpdateTopics) ([]*evm.TopicValues, error)
	DecodeBatchMetadataUpdate(log *evm.Log) (*BatchMetadataUpdateDecoded, error)
	EIP712DomainChangedLogHash() []byte
	EncodeEIP712DomainChangedTopics(evt abi.Event, values []EIP712DomainChangedTopics) ([]*evm.TopicValues, error)
	DecodeEIP712DomainChanged(log *evm.Log) (*EIP712DomainChangedDecoded, error)
	InitializedLogHash() []byte
	EncodeInitializedTopics(evt abi.Event, values []InitializedTopics) ([]*evm.TopicValues, error)
	DecodeInitialized(log *evm.Log) (*InitializedDecoded, error)
	MetadataSetLogHash() []byte
	EncodeMetadataSetTopics(evt abi.Event, values []MetadataSetTopics) ([]*evm.TopicValues, error)
	DecodeMetadataSet(log *evm.Log) (*MetadataSetDecoded, error)
	MetadataUpdateLogHash() []byte
	EncodeMetadataUpdateTopics(evt abi.Event, values []MetadataUpdateTopics) ([]*evm.TopicValues, error)
	DecodeMetadataUpdate(log *evm.Log) (*MetadataUpdateDecoded, error)
	OwnershipTransferredLogHash() []byte
	EncodeOwnershipTransferredTopics(evt abi.Event, values []OwnershipTransferredTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error)
	RegisteredLogHash() []byte
	EncodeRegisteredTopics(evt abi.Event, values []RegisteredTopics) ([]*evm.TopicValues, error)
	DecodeRegistered(log *evm.Log) (*RegisteredDecoded, error)
	TransferLogHash() []byte
	EncodeTransferTopics(evt abi.Event, values []TransferTopics) ([]*evm.TopicValues, error)
	DecodeTransfer(log *evm.Log) (*TransferDecoded, error)
	URIUpdatedLogHash() []byte
	EncodeURIUpdatedTopics(evt abi.Event, values []URIUpdatedTopics) ([]*evm.TopicValues, error)
	DecodeURIUpdated(log *evm.Log) (*URIUpdatedDecoded, error)
	UpgradedLogHash() []byte
	EncodeUpgradedTopics(evt abi.Event, values []UpgradedTopics) ([]*evm.TopicValues, error)
	DecodeUpgraded(log *evm.Log) (*UpgradedDecoded, error)
}

func NewIdentityRegistry(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*IdentityRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &IdentityRegistry{
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

func NewCodec() (IdentityRegistryCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryMetaData.ABI))
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

func (c *Codec) EncodeApproveMethodCall(in ApproveInput) ([]byte, error) {
	return c.abi.Pack("approve", in.To, in.TokenId)
}

func (c *Codec) EncodeBalanceOfMethodCall(in BalanceOfInput) ([]byte, error) {
	return c.abi.Pack("balanceOf", in.Owner)
}

func (c *Codec) DecodeBalanceOfMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["balanceOf"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeEip712DomainMethodCall() ([]byte, error) {
	return c.abi.Pack("eip712Domain")
}

func (c *Codec) DecodeEip712DomainMethodOutput(data []byte) (Eip712DomainOutput, error) {
	vals, err := c.abi.Methods["eip712Domain"].Outputs.Unpack(data)
	if err != nil {
		return Eip712DomainOutput{}, err
	}
	if len(vals) != 7 {
		return Eip712DomainOutput{}, fmt.Errorf("expected 7 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 [1]byte
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to unmarshal to [1]byte: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 string
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to unmarshal to string: %w", err)
	}
	jsonData2, err := json.Marshal(vals[2])
	if err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to marshal ABI result 2: %w", err)
	}

	var result2 string
	if err := json.Unmarshal(jsonData2, &result2); err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to unmarshal to string: %w", err)
	}
	jsonData3, err := json.Marshal(vals[3])
	if err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to marshal ABI result 3: %w", err)
	}

	var result3 *big.Int
	if err := json.Unmarshal(jsonData3, &result3); err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData4, err := json.Marshal(vals[4])
	if err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to marshal ABI result 4: %w", err)
	}

	var result4 common.Address
	if err := json.Unmarshal(jsonData4, &result4); err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}
	jsonData5, err := json.Marshal(vals[5])
	if err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to marshal ABI result 5: %w", err)
	}

	var result5 [32]byte
	if err := json.Unmarshal(jsonData5, &result5); err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to unmarshal to [32]byte: %w", err)
	}
	jsonData6, err := json.Marshal(vals[6])
	if err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to marshal ABI result 6: %w", err)
	}

	var result6 []*big.Int
	if err := json.Unmarshal(jsonData6, &result6); err != nil {
		return Eip712DomainOutput{}, fmt.Errorf("failed to unmarshal to []*big.Int: %w", err)
	}

	return Eip712DomainOutput{
		Fields:            result0,
		Name:              result1,
		Version:           result2,
		ChainId:           result3,
		VerifyingContract: result4,
		Salt:              result5,
		Extensions:        result6,
	}, nil
}

func (c *Codec) EncodeGetAgentWalletMethodCall(in GetAgentWalletInput) ([]byte, error) {
	return c.abi.Pack("getAgentWallet", in.AgentId)
}

func (c *Codec) DecodeGetAgentWalletMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["getAgentWallet"].Outputs.Unpack(data)
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

func (c *Codec) EncodeGetApprovedMethodCall(in GetApprovedInput) ([]byte, error) {
	return c.abi.Pack("getApproved", in.TokenId)
}

func (c *Codec) DecodeGetApprovedMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["getApproved"].Outputs.Unpack(data)
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

func (c *Codec) EncodeGetMetadataMethodCall(in GetMetadataInput) ([]byte, error) {
	return c.abi.Pack("getMetadata", in.AgentId, in.MetadataKey)
}

func (c *Codec) DecodeGetMetadataMethodOutput(data []byte) ([]byte, error) {
	vals, err := c.abi.Methods["getMetadata"].Outputs.Unpack(data)
	if err != nil {
		return *new([]byte), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([]byte), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result []byte
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([]byte), fmt.Errorf("failed to unmarshal to []byte: %w", err)
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

func (c *Codec) EncodeInitializeMethodCall() ([]byte, error) {
	return c.abi.Pack("initialize")
}

func (c *Codec) EncodeIsApprovedForAllMethodCall(in IsApprovedForAllInput) ([]byte, error) {
	return c.abi.Pack("isApprovedForAll", in.Owner, in.Operator)
}

func (c *Codec) DecodeIsApprovedForAllMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["isApprovedForAll"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeNameMethodCall() ([]byte, error) {
	return c.abi.Pack("name")
}

func (c *Codec) DecodeNameMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["name"].Outputs.Unpack(data)
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

func (c *Codec) EncodeOwnerOfMethodCall(in OwnerOfInput) ([]byte, error) {
	return c.abi.Pack("ownerOf", in.TokenId)
}

func (c *Codec) DecodeOwnerOfMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["ownerOf"].Outputs.Unpack(data)
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

func (c *Codec) EncodeRegisterMethodCall() ([]byte, error) {
	return c.abi.Pack("register")
}

func (c *Codec) DecodeRegisterMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["register"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeRegister0MethodCall(in Register0Input) ([]byte, error) {
	return c.abi.Pack("register0", in.AgentURI, in.Metadata)
}

func (c *Codec) DecodeRegister0MethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["register0"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeRegister1MethodCall(in Register1Input) ([]byte, error) {
	return c.abi.Pack("register1", in.AgentURI)
}

func (c *Codec) DecodeRegister1MethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["register1"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeRenounceOwnershipMethodCall() ([]byte, error) {
	return c.abi.Pack("renounceOwnership")
}

func (c *Codec) EncodeSafeTransferFromMethodCall(in SafeTransferFromInput) ([]byte, error) {
	return c.abi.Pack("safeTransferFrom", in.From, in.To, in.TokenId)
}

func (c *Codec) EncodeSafeTransferFrom0MethodCall(in SafeTransferFrom0Input) ([]byte, error) {
	return c.abi.Pack("safeTransferFrom0", in.From, in.To, in.TokenId, in.Data)
}

func (c *Codec) EncodeSetAgentURIMethodCall(in SetAgentURIInput) ([]byte, error) {
	return c.abi.Pack("setAgentURI", in.AgentId, in.NewURI)
}

func (c *Codec) EncodeSetAgentWalletMethodCall(in SetAgentWalletInput) ([]byte, error) {
	return c.abi.Pack("setAgentWallet", in.AgentId, in.NewWallet, in.Deadline, in.Signature)
}

func (c *Codec) EncodeSetApprovalForAllMethodCall(in SetApprovalForAllInput) ([]byte, error) {
	return c.abi.Pack("setApprovalForAll", in.Operator, in.Approved)
}

func (c *Codec) EncodeSetMetadataMethodCall(in SetMetadataInput) ([]byte, error) {
	return c.abi.Pack("setMetadata", in.AgentId, in.MetadataKey, in.MetadataValue)
}

func (c *Codec) EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error) {
	return c.abi.Pack("supportsInterface", in.InterfaceId)
}

func (c *Codec) DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["supportsInterface"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeSymbolMethodCall() ([]byte, error) {
	return c.abi.Pack("symbol")
}

func (c *Codec) DecodeSymbolMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["symbol"].Outputs.Unpack(data)
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

func (c *Codec) EncodeTokenURIMethodCall(in TokenURIInput) ([]byte, error) {
	return c.abi.Pack("tokenURI", in.TokenId)
}

func (c *Codec) DecodeTokenURIMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["tokenURI"].Outputs.Unpack(data)
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

func (c *Codec) EncodeTransferFromMethodCall(in TransferFromInput) ([]byte, error) {
	return c.abi.Pack("transferFrom", in.From, in.To, in.TokenId)
}

func (c *Codec) EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error) {
	return c.abi.Pack("transferOwnership", in.NewOwner)
}

func (c *Codec) EncodeUnsetAgentWalletMethodCall(in UnsetAgentWalletInput) ([]byte, error) {
	return c.abi.Pack("unsetAgentWallet", in.AgentId)
}

func (c *Codec) EncodeUpgradeToAndCallMethodCall(in UpgradeToAndCallInput) ([]byte, error) {
	return c.abi.Pack("upgradeToAndCall", in.NewImplementation, in.Data)
}

func (c *Codec) EncodeUpgradeableMetadataEntryStruct(in UpgradeableMetadataEntry) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "metadataKey", Type: "string"},
			{Name: "metadataValue", Type: "bytes"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for UpgradeableMetadataEntry: %w", err)
	}
	args := abi.Arguments{
		{Name: "upgradeableMetadataEntry", Type: tupleType},
	}

	return args.Pack(in)
}

func (c *Codec) ApprovalLogHash() []byte {
	return c.abi.Events["Approval"].ID.Bytes()
}

func (c *Codec) EncodeApprovalTopics(
	evt abi.Event,
	values []ApprovalTopics,
) ([]*evm.TopicValues, error) {
	var ownerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Owner).IsZero() {
			ownerRule = append(ownerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Owner)
		if err != nil {
			return nil, err
		}
		ownerRule = append(ownerRule, fieldVal)
	}
	var approvedRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Approved).IsZero() {
			approvedRule = append(approvedRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.Approved)
		if err != nil {
			return nil, err
		}
		approvedRule = append(approvedRule, fieldVal)
	}
	var tokenIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.TokenId).IsZero() {
			tokenIdRule = append(tokenIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[2], v.TokenId)
		if err != nil {
			return nil, err
		}
		tokenIdRule = append(tokenIdRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		ownerRule,
		approvedRule,
		tokenIdRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeApproval decodes a log into a Approval struct.
func (c *Codec) DecodeApproval(log *evm.Log) (*ApprovalDecoded, error) {
	event := new(ApprovalDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Approval", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Approval"].Inputs {
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

func (c *Codec) ApprovalForAllLogHash() []byte {
	return c.abi.Events["ApprovalForAll"].ID.Bytes()
}

func (c *Codec) EncodeApprovalForAllTopics(
	evt abi.Event,
	values []ApprovalForAllTopics,
) ([]*evm.TopicValues, error) {
	var ownerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Owner).IsZero() {
			ownerRule = append(ownerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Owner)
		if err != nil {
			return nil, err
		}
		ownerRule = append(ownerRule, fieldVal)
	}
	var operatorRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Operator).IsZero() {
			operatorRule = append(operatorRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.Operator)
		if err != nil {
			return nil, err
		}
		operatorRule = append(operatorRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		ownerRule,
		operatorRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeApprovalForAll decodes a log into a ApprovalForAll struct.
func (c *Codec) DecodeApprovalForAll(log *evm.Log) (*ApprovalForAllDecoded, error) {
	event := new(ApprovalForAllDecoded)
	if err := c.abi.UnpackIntoInterface(event, "ApprovalForAll", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["ApprovalForAll"].Inputs {
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

func (c *Codec) BatchMetadataUpdateLogHash() []byte {
	return c.abi.Events["BatchMetadataUpdate"].ID.Bytes()
}

func (c *Codec) EncodeBatchMetadataUpdateTopics(
	evt abi.Event,
	values []BatchMetadataUpdateTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeBatchMetadataUpdate decodes a log into a BatchMetadataUpdate struct.
func (c *Codec) DecodeBatchMetadataUpdate(log *evm.Log) (*BatchMetadataUpdateDecoded, error) {
	event := new(BatchMetadataUpdateDecoded)
	if err := c.abi.UnpackIntoInterface(event, "BatchMetadataUpdate", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["BatchMetadataUpdate"].Inputs {
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

func (c *Codec) EIP712DomainChangedLogHash() []byte {
	return c.abi.Events["EIP712DomainChanged"].ID.Bytes()
}

func (c *Codec) EncodeEIP712DomainChangedTopics(
	evt abi.Event,
	values []EIP712DomainChangedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeEIP712DomainChanged decodes a log into a EIP712DomainChanged struct.
func (c *Codec) DecodeEIP712DomainChanged(log *evm.Log) (*EIP712DomainChangedDecoded, error) {
	event := new(EIP712DomainChangedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "EIP712DomainChanged", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["EIP712DomainChanged"].Inputs {
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

func (c *Codec) MetadataSetLogHash() []byte {
	return c.abi.Events["MetadataSet"].ID.Bytes()
}

func (c *Codec) EncodeMetadataSetTopics(
	evt abi.Event,
	values []MetadataSetTopics,
) ([]*evm.TopicValues, error) {
	var agentIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.AgentId).IsZero() {
			agentIdRule = append(agentIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.AgentId)
		if err != nil {
			return nil, err
		}
		agentIdRule = append(agentIdRule, fieldVal)
	}
	var indexedMetadataKeyRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.IndexedMetadataKey).IsZero() {
			indexedMetadataKeyRule = append(indexedMetadataKeyRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.IndexedMetadataKey)
		if err != nil {
			return nil, err
		}
		indexedMetadataKeyRule = append(indexedMetadataKeyRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		agentIdRule,
		indexedMetadataKeyRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeMetadataSet decodes a log into a MetadataSet struct.
func (c *Codec) DecodeMetadataSet(log *evm.Log) (*MetadataSetDecoded, error) {
	event := new(MetadataSetDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MetadataSet", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["MetadataSet"].Inputs {
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

func (c *Codec) MetadataUpdateLogHash() []byte {
	return c.abi.Events["MetadataUpdate"].ID.Bytes()
}

func (c *Codec) EncodeMetadataUpdateTopics(
	evt abi.Event,
	values []MetadataUpdateTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeMetadataUpdate decodes a log into a MetadataUpdate struct.
func (c *Codec) DecodeMetadataUpdate(log *evm.Log) (*MetadataUpdateDecoded, error) {
	event := new(MetadataUpdateDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MetadataUpdate", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["MetadataUpdate"].Inputs {
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

func (c *Codec) RegisteredLogHash() []byte {
	return c.abi.Events["Registered"].ID.Bytes()
}

func (c *Codec) EncodeRegisteredTopics(
	evt abi.Event,
	values []RegisteredTopics,
) ([]*evm.TopicValues, error) {
	var agentIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.AgentId).IsZero() {
			agentIdRule = append(agentIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.AgentId)
		if err != nil {
			return nil, err
		}
		agentIdRule = append(agentIdRule, fieldVal)
	}
	var ownerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Owner).IsZero() {
			ownerRule = append(ownerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[2], v.Owner)
		if err != nil {
			return nil, err
		}
		ownerRule = append(ownerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		agentIdRule,
		ownerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeRegistered decodes a log into a Registered struct.
func (c *Codec) DecodeRegistered(log *evm.Log) (*RegisteredDecoded, error) {
	event := new(RegisteredDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Registered", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Registered"].Inputs {
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

func (c *Codec) TransferLogHash() []byte {
	return c.abi.Events["Transfer"].ID.Bytes()
}

func (c *Codec) EncodeTransferTopics(
	evt abi.Event,
	values []TransferTopics,
) ([]*evm.TopicValues, error) {
	var fromRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.From).IsZero() {
			fromRule = append(fromRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.From)
		if err != nil {
			return nil, err
		}
		fromRule = append(fromRule, fieldVal)
	}
	var toRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.To).IsZero() {
			toRule = append(toRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.To)
		if err != nil {
			return nil, err
		}
		toRule = append(toRule, fieldVal)
	}
	var tokenIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.TokenId).IsZero() {
			tokenIdRule = append(tokenIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[2], v.TokenId)
		if err != nil {
			return nil, err
		}
		tokenIdRule = append(tokenIdRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		fromRule,
		toRule,
		tokenIdRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeTransfer decodes a log into a Transfer struct.
func (c *Codec) DecodeTransfer(log *evm.Log) (*TransferDecoded, error) {
	event := new(TransferDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Transfer", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Transfer"].Inputs {
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

func (c *Codec) URIUpdatedLogHash() []byte {
	return c.abi.Events["URIUpdated"].ID.Bytes()
}

func (c *Codec) EncodeURIUpdatedTopics(
	evt abi.Event,
	values []URIUpdatedTopics,
) ([]*evm.TopicValues, error) {
	var agentIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.AgentId).IsZero() {
			agentIdRule = append(agentIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.AgentId)
		if err != nil {
			return nil, err
		}
		agentIdRule = append(agentIdRule, fieldVal)
	}
	var updatedByRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.UpdatedBy).IsZero() {
			updatedByRule = append(updatedByRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[2], v.UpdatedBy)
		if err != nil {
			return nil, err
		}
		updatedByRule = append(updatedByRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		agentIdRule,
		updatedByRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeURIUpdated decodes a log into a URIUpdated struct.
func (c *Codec) DecodeURIUpdated(log *evm.Log) (*URIUpdatedDecoded, error) {
	event := new(URIUpdatedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "URIUpdated", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["URIUpdated"].Inputs {
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

func (c IdentityRegistry) UPGRADEINTERFACEVERSION(
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

func (c IdentityRegistry) BalanceOf(
	runtime cre.Runtime,
	args BalanceOfInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeBalanceOfMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeBalanceOfMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) Eip712Domain(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[Eip712DomainOutput] {
	calldata, err := c.Codec.EncodeEip712DomainMethodCall()
	if err != nil {
		return cre.PromiseFromResult[Eip712DomainOutput](Eip712DomainOutput{}, err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (Eip712DomainOutput, error) {
		return c.Codec.DecodeEip712DomainMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) GetAgentWallet(
	runtime cre.Runtime,
	args GetAgentWalletInput,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeGetAgentWalletMethodCall(args)
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
		return c.Codec.DecodeGetAgentWalletMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) GetApproved(
	runtime cre.Runtime,
	args GetApprovedInput,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeGetApprovedMethodCall(args)
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
		return c.Codec.DecodeGetApprovedMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) GetMetadata(
	runtime cre.Runtime,
	args GetMetadataInput,
	blockNumber *big.Int,
) cre.Promise[[]byte] {
	calldata, err := c.Codec.EncodeGetMetadataMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[[]byte](*new([]byte), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) ([]byte, error) {
		return c.Codec.DecodeGetMetadataMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) IsApprovedForAll(
	runtime cre.Runtime,
	args IsApprovedForAllInput,
	blockNumber *big.Int,
) cre.Promise[bool] {
	calldata, err := c.Codec.EncodeIsApprovedForAllMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[bool](*new(bool), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (bool, error) {
		return c.Codec.DecodeIsApprovedForAllMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) Name(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[string] {
	calldata, err := c.Codec.EncodeNameMethodCall()
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
		return c.Codec.DecodeNameMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) Owner(
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

func (c IdentityRegistry) OwnerOf(
	runtime cre.Runtime,
	args OwnerOfInput,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOwnerOfMethodCall(args)
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
		return c.Codec.DecodeOwnerOfMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) ProxiableUUID(
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

func (c IdentityRegistry) SupportsInterface(
	runtime cre.Runtime,
	args SupportsInterfaceInput,
	blockNumber *big.Int,
) cre.Promise[bool] {
	calldata, err := c.Codec.EncodeSupportsInterfaceMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[bool](*new(bool), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (bool, error) {
		return c.Codec.DecodeSupportsInterfaceMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) Symbol(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[string] {
	calldata, err := c.Codec.EncodeSymbolMethodCall()
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
		return c.Codec.DecodeSymbolMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) TokenURI(
	runtime cre.Runtime,
	args TokenURIInput,
	blockNumber *big.Int,
) cre.Promise[string] {
	calldata, err := c.Codec.EncodeTokenURIMethodCall(args)
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
		return c.Codec.DecodeTokenURIMethodOutput(response.Data)
	})

}

func (c IdentityRegistry) WriteReportFromUpgradeableMetadataEntry(
	runtime cre.Runtime,
	input UpgradeableMetadataEntry,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeUpgradeableMetadataEntryStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c IdentityRegistry) WriteReport(
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
func (c *IdentityRegistry) DecodeAddressEmptyCodeError(data []byte) (*AddressEmptyCode, error) {
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

// DecodeECDSAInvalidSignatureError decodes a ECDSAInvalidSignature error from revert data.
func (c *IdentityRegistry) DecodeECDSAInvalidSignatureError(data []byte) (*ECDSAInvalidSignature, error) {
	args := c.ABI.Errors["ECDSAInvalidSignature"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &ECDSAInvalidSignature{}, nil
}

// Error implements the error interface for ECDSAInvalidSignature.
func (e *ECDSAInvalidSignature) Error() string {
	return fmt.Sprintf("ECDSAInvalidSignature error:")
}

// DecodeECDSAInvalidSignatureLengthError decodes a ECDSAInvalidSignatureLength error from revert data.
func (c *IdentityRegistry) DecodeECDSAInvalidSignatureLengthError(data []byte) (*ECDSAInvalidSignatureLength, error) {
	args := c.ABI.Errors["ECDSAInvalidSignatureLength"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	length, ok0 := values[0].(*big.Int)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for length in ECDSAInvalidSignatureLength error")
	}

	return &ECDSAInvalidSignatureLength{
		Length: length,
	}, nil
}

// Error implements the error interface for ECDSAInvalidSignatureLength.
func (e *ECDSAInvalidSignatureLength) Error() string {
	return fmt.Sprintf("ECDSAInvalidSignatureLength error: length=%v;", e.Length)
}

// DecodeECDSAInvalidSignatureSError decodes a ECDSAInvalidSignatureS error from revert data.
func (c *IdentityRegistry) DecodeECDSAInvalidSignatureSError(data []byte) (*ECDSAInvalidSignatureS, error) {
	args := c.ABI.Errors["ECDSAInvalidSignatureS"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	s, ok0 := values[0].([32]byte)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for s in ECDSAInvalidSignatureS error")
	}

	return &ECDSAInvalidSignatureS{
		S: s,
	}, nil
}

// Error implements the error interface for ECDSAInvalidSignatureS.
func (e *ECDSAInvalidSignatureS) Error() string {
	return fmt.Sprintf("ECDSAInvalidSignatureS error: s=%v;", e.S)
}

// DecodeERC1967InvalidImplementationError decodes a ERC1967InvalidImplementation error from revert data.
func (c *IdentityRegistry) DecodeERC1967InvalidImplementationError(data []byte) (*ERC1967InvalidImplementation, error) {
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
func (c *IdentityRegistry) DecodeERC1967NonPayableError(data []byte) (*ERC1967NonPayable, error) {
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

// DecodeERC721IncorrectOwnerError decodes a ERC721IncorrectOwner error from revert data.
func (c *IdentityRegistry) DecodeERC721IncorrectOwnerError(data []byte) (*ERC721IncorrectOwner, error) {
	args := c.ABI.Errors["ERC721IncorrectOwner"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 3 {
		return nil, fmt.Errorf("expected 3 values, got %d", len(values))
	}

	sender, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for sender in ERC721IncorrectOwner error")
	}

	tokenId, ok1 := values[1].(*big.Int)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for tokenId in ERC721IncorrectOwner error")
	}

	owner, ok2 := values[2].(common.Address)
	if !ok2 {
		return nil, fmt.Errorf("unexpected type for owner in ERC721IncorrectOwner error")
	}

	return &ERC721IncorrectOwner{
		Sender:  sender,
		TokenId: tokenId,
		Owner:   owner,
	}, nil
}

// Error implements the error interface for ERC721IncorrectOwner.
func (e *ERC721IncorrectOwner) Error() string {
	return fmt.Sprintf("ERC721IncorrectOwner error: sender=%v; tokenId=%v; owner=%v;", e.Sender, e.TokenId, e.Owner)
}

// DecodeERC721InsufficientApprovalError decodes a ERC721InsufficientApproval error from revert data.
func (c *IdentityRegistry) DecodeERC721InsufficientApprovalError(data []byte) (*ERC721InsufficientApproval, error) {
	args := c.ABI.Errors["ERC721InsufficientApproval"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values, got %d", len(values))
	}

	operator, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for operator in ERC721InsufficientApproval error")
	}

	tokenId, ok1 := values[1].(*big.Int)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for tokenId in ERC721InsufficientApproval error")
	}

	return &ERC721InsufficientApproval{
		Operator: operator,
		TokenId:  tokenId,
	}, nil
}

// Error implements the error interface for ERC721InsufficientApproval.
func (e *ERC721InsufficientApproval) Error() string {
	return fmt.Sprintf("ERC721InsufficientApproval error: operator=%v; tokenId=%v;", e.Operator, e.TokenId)
}

// DecodeERC721InvalidApproverError decodes a ERC721InvalidApprover error from revert data.
func (c *IdentityRegistry) DecodeERC721InvalidApproverError(data []byte) (*ERC721InvalidApprover, error) {
	args := c.ABI.Errors["ERC721InvalidApprover"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	approver, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for approver in ERC721InvalidApprover error")
	}

	return &ERC721InvalidApprover{
		Approver: approver,
	}, nil
}

// Error implements the error interface for ERC721InvalidApprover.
func (e *ERC721InvalidApprover) Error() string {
	return fmt.Sprintf("ERC721InvalidApprover error: approver=%v;", e.Approver)
}

// DecodeERC721InvalidOperatorError decodes a ERC721InvalidOperator error from revert data.
func (c *IdentityRegistry) DecodeERC721InvalidOperatorError(data []byte) (*ERC721InvalidOperator, error) {
	args := c.ABI.Errors["ERC721InvalidOperator"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	operator, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for operator in ERC721InvalidOperator error")
	}

	return &ERC721InvalidOperator{
		Operator: operator,
	}, nil
}

// Error implements the error interface for ERC721InvalidOperator.
func (e *ERC721InvalidOperator) Error() string {
	return fmt.Sprintf("ERC721InvalidOperator error: operator=%v;", e.Operator)
}

// DecodeERC721InvalidOwnerError decodes a ERC721InvalidOwner error from revert data.
func (c *IdentityRegistry) DecodeERC721InvalidOwnerError(data []byte) (*ERC721InvalidOwner, error) {
	args := c.ABI.Errors["ERC721InvalidOwner"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	owner, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for owner in ERC721InvalidOwner error")
	}

	return &ERC721InvalidOwner{
		Owner: owner,
	}, nil
}

// Error implements the error interface for ERC721InvalidOwner.
func (e *ERC721InvalidOwner) Error() string {
	return fmt.Sprintf("ERC721InvalidOwner error: owner=%v;", e.Owner)
}

// DecodeERC721InvalidReceiverError decodes a ERC721InvalidReceiver error from revert data.
func (c *IdentityRegistry) DecodeERC721InvalidReceiverError(data []byte) (*ERC721InvalidReceiver, error) {
	args := c.ABI.Errors["ERC721InvalidReceiver"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	receiver, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for receiver in ERC721InvalidReceiver error")
	}

	return &ERC721InvalidReceiver{
		Receiver: receiver,
	}, nil
}

// Error implements the error interface for ERC721InvalidReceiver.
func (e *ERC721InvalidReceiver) Error() string {
	return fmt.Sprintf("ERC721InvalidReceiver error: receiver=%v;", e.Receiver)
}

// DecodeERC721InvalidSenderError decodes a ERC721InvalidSender error from revert data.
func (c *IdentityRegistry) DecodeERC721InvalidSenderError(data []byte) (*ERC721InvalidSender, error) {
	args := c.ABI.Errors["ERC721InvalidSender"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	sender, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for sender in ERC721InvalidSender error")
	}

	return &ERC721InvalidSender{
		Sender: sender,
	}, nil
}

// Error implements the error interface for ERC721InvalidSender.
func (e *ERC721InvalidSender) Error() string {
	return fmt.Sprintf("ERC721InvalidSender error: sender=%v;", e.Sender)
}

// DecodeERC721NonexistentTokenError decodes a ERC721NonexistentToken error from revert data.
func (c *IdentityRegistry) DecodeERC721NonexistentTokenError(data []byte) (*ERC721NonexistentToken, error) {
	args := c.ABI.Errors["ERC721NonexistentToken"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	tokenId, ok0 := values[0].(*big.Int)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for tokenId in ERC721NonexistentToken error")
	}

	return &ERC721NonexistentToken{
		TokenId: tokenId,
	}, nil
}

// Error implements the error interface for ERC721NonexistentToken.
func (e *ERC721NonexistentToken) Error() string {
	return fmt.Sprintf("ERC721NonexistentToken error: tokenId=%v;", e.TokenId)
}

// DecodeFailedCallError decodes a FailedCall error from revert data.
func (c *IdentityRegistry) DecodeFailedCallError(data []byte) (*FailedCall, error) {
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
func (c *IdentityRegistry) DecodeInvalidInitializationError(data []byte) (*InvalidInitialization, error) {
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
func (c *IdentityRegistry) DecodeNotInitializingError(data []byte) (*NotInitializing, error) {
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
func (c *IdentityRegistry) DecodeOwnableInvalidOwnerError(data []byte) (*OwnableInvalidOwner, error) {
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
func (c *IdentityRegistry) DecodeOwnableUnauthorizedAccountError(data []byte) (*OwnableUnauthorizedAccount, error) {
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
func (c *IdentityRegistry) DecodeUUPSUnauthorizedCallContextError(data []byte) (*UUPSUnauthorizedCallContext, error) {
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
func (c *IdentityRegistry) DecodeUUPSUnsupportedProxiableUUIDError(data []byte) (*UUPSUnsupportedProxiableUUID, error) {
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

func (c *IdentityRegistry) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["AddressEmptyCode"].ID.Bytes()[:4]):
		return c.DecodeAddressEmptyCodeError(data)
	case common.Bytes2Hex(c.ABI.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]):
		return c.DecodeECDSAInvalidSignatureError(data)
	case common.Bytes2Hex(c.ABI.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]):
		return c.DecodeECDSAInvalidSignatureLengthError(data)
	case common.Bytes2Hex(c.ABI.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]):
		return c.DecodeECDSAInvalidSignatureSError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]):
		return c.DecodeERC1967InvalidImplementationError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC1967NonPayable"].ID.Bytes()[:4]):
		return c.DecodeERC1967NonPayableError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC721IncorrectOwner"].ID.Bytes()[:4]):
		return c.DecodeERC721IncorrectOwnerError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC721InsufficientApproval"].ID.Bytes()[:4]):
		return c.DecodeERC721InsufficientApprovalError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC721InvalidApprover"].ID.Bytes()[:4]):
		return c.DecodeERC721InvalidApproverError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC721InvalidOperator"].ID.Bytes()[:4]):
		return c.DecodeERC721InvalidOperatorError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC721InvalidOwner"].ID.Bytes()[:4]):
		return c.DecodeERC721InvalidOwnerError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC721InvalidReceiver"].ID.Bytes()[:4]):
		return c.DecodeERC721InvalidReceiverError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC721InvalidSender"].ID.Bytes()[:4]):
		return c.DecodeERC721InvalidSenderError(data)
	case common.Bytes2Hex(c.ABI.Errors["ERC721NonexistentToken"].ID.Bytes()[:4]):
		return c.DecodeERC721NonexistentTokenError(data)
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

// ApprovalTrigger wraps the raw log trigger and provides decoded ApprovalDecoded data
type ApprovalTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into Approval data
func (t *ApprovalTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[ApprovalDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeApproval(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Approval log: %w", err)
	}

	return &bindings.DecodedLog[ApprovalDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *IdentityRegistry) LogTriggerApprovalLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []ApprovalTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[ApprovalDecoded]], error) {
	event := c.ABI.Events["Approval"]
	topics, err := c.Codec.EncodeApprovalTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Approval: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &ApprovalTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *IdentityRegistry) FilterLogsApproval(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.ApprovalLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// ApprovalForAllTrigger wraps the raw log trigger and provides decoded ApprovalForAllDecoded data
type ApprovalForAllTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into ApprovalForAll data
func (t *ApprovalForAllTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[ApprovalForAllDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeApprovalForAll(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ApprovalForAll log: %w", err)
	}

	return &bindings.DecodedLog[ApprovalForAllDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *IdentityRegistry) LogTriggerApprovalForAllLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []ApprovalForAllTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[ApprovalForAllDecoded]], error) {
	event := c.ABI.Events["ApprovalForAll"]
	topics, err := c.Codec.EncodeApprovalForAllTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for ApprovalForAll: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &ApprovalForAllTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *IdentityRegistry) FilterLogsApprovalForAll(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.ApprovalForAllLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// BatchMetadataUpdateTrigger wraps the raw log trigger and provides decoded BatchMetadataUpdateDecoded data
type BatchMetadataUpdateTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into BatchMetadataUpdate data
func (t *BatchMetadataUpdateTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[BatchMetadataUpdateDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeBatchMetadataUpdate(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode BatchMetadataUpdate log: %w", err)
	}

	return &bindings.DecodedLog[BatchMetadataUpdateDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *IdentityRegistry) LogTriggerBatchMetadataUpdateLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []BatchMetadataUpdateTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[BatchMetadataUpdateDecoded]], error) {
	event := c.ABI.Events["BatchMetadataUpdate"]
	topics, err := c.Codec.EncodeBatchMetadataUpdateTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for BatchMetadataUpdate: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &BatchMetadataUpdateTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *IdentityRegistry) FilterLogsBatchMetadataUpdate(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.BatchMetadataUpdateLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// EIP712DomainChangedTrigger wraps the raw log trigger and provides decoded EIP712DomainChangedDecoded data
type EIP712DomainChangedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into EIP712DomainChanged data
func (t *EIP712DomainChangedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[EIP712DomainChangedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeEIP712DomainChanged(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode EIP712DomainChanged log: %w", err)
	}

	return &bindings.DecodedLog[EIP712DomainChangedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *IdentityRegistry) LogTriggerEIP712DomainChangedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []EIP712DomainChangedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[EIP712DomainChangedDecoded]], error) {
	event := c.ABI.Events["EIP712DomainChanged"]
	topics, err := c.Codec.EncodeEIP712DomainChangedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for EIP712DomainChanged: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &EIP712DomainChangedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *IdentityRegistry) FilterLogsEIP712DomainChanged(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.EIP712DomainChangedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// InitializedTrigger wraps the raw log trigger and provides decoded InitializedDecoded data
type InitializedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
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

func (c *IdentityRegistry) LogTriggerInitializedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []InitializedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[InitializedDecoded]], error) {
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

func (c *IdentityRegistry) FilterLogsInitialized(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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

// MetadataSetTrigger wraps the raw log trigger and provides decoded MetadataSetDecoded data
type MetadataSetTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into MetadataSet data
func (t *MetadataSetTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[MetadataSetDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeMetadataSet(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MetadataSet log: %w", err)
	}

	return &bindings.DecodedLog[MetadataSetDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *IdentityRegistry) LogTriggerMetadataSetLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []MetadataSetTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[MetadataSetDecoded]], error) {
	event := c.ABI.Events["MetadataSet"]
	topics, err := c.Codec.EncodeMetadataSetTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for MetadataSet: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &MetadataSetTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *IdentityRegistry) FilterLogsMetadataSet(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.MetadataSetLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// MetadataUpdateTrigger wraps the raw log trigger and provides decoded MetadataUpdateDecoded data
type MetadataUpdateTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into MetadataUpdate data
func (t *MetadataUpdateTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[MetadataUpdateDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeMetadataUpdate(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MetadataUpdate log: %w", err)
	}

	return &bindings.DecodedLog[MetadataUpdateDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *IdentityRegistry) LogTriggerMetadataUpdateLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []MetadataUpdateTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[MetadataUpdateDecoded]], error) {
	event := c.ABI.Events["MetadataUpdate"]
	topics, err := c.Codec.EncodeMetadataUpdateTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for MetadataUpdate: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &MetadataUpdateTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *IdentityRegistry) FilterLogsMetadataUpdate(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.MetadataUpdateLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OwnershipTransferredTrigger wraps the raw log trigger and provides decoded OwnershipTransferredDecoded data
type OwnershipTransferredTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
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

func (c *IdentityRegistry) LogTriggerOwnershipTransferredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipTransferredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipTransferredDecoded]], error) {
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

func (c *IdentityRegistry) FilterLogsOwnershipTransferred(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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

// RegisteredTrigger wraps the raw log trigger and provides decoded RegisteredDecoded data
type RegisteredTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into Registered data
func (t *RegisteredTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[RegisteredDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeRegistered(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Registered log: %w", err)
	}

	return &bindings.DecodedLog[RegisteredDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *IdentityRegistry) LogTriggerRegisteredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []RegisteredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[RegisteredDecoded]], error) {
	event := c.ABI.Events["Registered"]
	topics, err := c.Codec.EncodeRegisteredTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Registered: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &RegisteredTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *IdentityRegistry) FilterLogsRegistered(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.RegisteredLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// TransferTrigger wraps the raw log trigger and provides decoded TransferDecoded data
type TransferTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into Transfer data
func (t *TransferTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[TransferDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeTransfer(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Transfer log: %w", err)
	}

	return &bindings.DecodedLog[TransferDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *IdentityRegistry) LogTriggerTransferLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []TransferTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[TransferDecoded]], error) {
	event := c.ABI.Events["Transfer"]
	topics, err := c.Codec.EncodeTransferTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Transfer: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &TransferTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *IdentityRegistry) FilterLogsTransfer(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.TransferLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// URIUpdatedTrigger wraps the raw log trigger and provides decoded URIUpdatedDecoded data
type URIUpdatedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
}

// Adapt method that decodes the log into URIUpdated data
func (t *URIUpdatedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[URIUpdatedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeURIUpdated(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode URIUpdated log: %w", err)
	}

	return &bindings.DecodedLog[URIUpdatedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *IdentityRegistry) LogTriggerURIUpdatedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []URIUpdatedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[URIUpdatedDecoded]], error) {
	event := c.ABI.Events["URIUpdated"]
	topics, err := c.Codec.EncodeURIUpdatedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for URIUpdated: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &URIUpdatedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *IdentityRegistry) FilterLogsURIUpdated(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.URIUpdatedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// UpgradedTrigger wraps the raw log trigger and provides decoded UpgradedDecoded data
type UpgradedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                   // Embed the raw trigger
	contract                        *IdentityRegistry // Keep reference for decoding
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

func (c *IdentityRegistry) LogTriggerUpgradedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []UpgradedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[UpgradedDecoded]], error) {
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

func (c *IdentityRegistry) FilterLogsUpgraded(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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
