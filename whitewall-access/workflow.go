//go:build wasip1

package main

import (
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/triggers/webapitr"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	"whitewall-cre/contracts/evm/src/generated/identity_registry"
)

// Config - config.staging.json에서 읽어옴
type Config struct {
	ChainName               string `json:"chainName"`
	IdentityRegistryAddress string `json:"identityRegistryAddress"`
	ConsumerAddress         string `json:"consumerAddress"`
}

// http 요청 페이로드
type AccessRequest struct {
	AgentID uint64 `json:"agentId"`
}

// report 결과
type AccessReport struct {
	AgentID          *big.Int       `abi:"agentId"`
	Approved         bool           `abi:"approved"`
	Tier             uint8          `abi:"tier"`
	AccountableHuman common.Address `abi:"accountableHuman"`
	Reason           [32]byte       `abi:"reason"`
}

// InitWorkflow
func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	//HTTP 트리거
	httpTrigger := webapitr.Trigger(&webapitr.Config{}) //HTTP 트리거 인스턴스

	return cre.Workflow[*Config]{
		cre.Handler(httpTrigger, onAccessRequest), //트리거-함수 연결
	}, nil
}

// 접근 요청 처리 핸들러
func onAccessRequest(config *Config, runtime cre.Runtime, trigger *webapitr.Payload) (*AccessReport, error) {
	logger := runtime.Logger()
	logger.Info("Access request received")

	// 1. HTTP - agentId 파싱
	var req AccessRequest
	if err := trigger.ParseBody(&req); err != nil {
		logger.Error("Failed to parse request body", "error", err)
		// 리턴할 때도 구조체 타입을 맞춰야 빨간 줄이 안 납니다!
		return &AccessReport{
			AgentID:          big.NewInt(0),
			Approved:         false,
			Tier:             0,
			AccountableHuman: common.Address{},
			Reason:           formatReason("INVALID_REQUEST"),
		}, nil
	}

	agentIdBig := new(big.Int).SetUint64(req.AgentID)
	logger.Info("Processing access request", "agentId", req.AgentID)

	// 2. EVM client 및 Registry 설정 (기존 로직 동일)
	chainSelector, _ := evm.ChainSelectorFromName(config.ChainName)
	evmClient := &evm.Client{ChainSelector: chainSelector}
	registryAddress := common.HexToAddress(config.IdentityRegistryAddress)
	registry, err := identity_registry.NewIdentityRegistry(evmClient, registryAddress, nil)
	if err != nil {
		return nil, err
	}

	// 3. ownerOf 호출
	// 여기서 만약 identity_registry 패키지에 따라 필드명이 다를 수 있으니 주의!
	owner, err := registry.OwnerOf(runtime, identity_registry.OwnerOfInput{TokenId: agentIdBig}, big.NewInt(rpc.LatestBlockNumber.Int64())).Await()

	if err != nil {
		logger.Warn("Agent not registered", "agentId", req.AgentID, "error", err)
		return &AccessReport{
			AgentID:          agentIdBig,
			Approved:         false,
			Tier:             0,
			AccountableHuman: common.Address{},
			Reason:           formatReason("NOT_REGISTERED"),
		}, nil
	}

	// 4. 리포트 최종 생성
	return &AccessReport{
		AgentID:          agentIdBig,
		Approved:         true,
		Tier:             2,
		AccountableHuman: owner,
		Reason:           [32]byte{}, // 0x00...
	}, nil
}

// 헬퍼 함수: string -> bytes32
func formatReason(reason string) [32]byte {
	var b [32]byte
	copy(b[:], []byte(reason))
	return b
}
