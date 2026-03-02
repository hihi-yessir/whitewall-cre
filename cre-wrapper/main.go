package main

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// authorizedKeys (workflow.yaml과 동일하게)
var authorizedSigners = map[string]bool{
	"0x0d10f69243b8a2fe4299fa4cc115c3023f4011cf": true,
}

// Validator 주소들
const (
	ValidationRegistryAddr   = "0x8004Cb1BF31DAf7788923b405b754f57acEB4272" // 나중에 받으면 추가
	StripeKYCValidatorAddr   = "0x12b456dcc0e669eeb1d96806c8ef87b713d39cc8"
	PlaidCreditValidatorAddr = "0x9a0ed706f1714961bf607404521a58decddc2636"
)

// ValidationRequest 이벤트 시그니처
var ValidationRequestTopic = crypto.Keccak256Hash([]byte("ValidationRequest(address,uint256,string,bytes32)"))

// JSON-RPC 요청 구조체
type JSONRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		WorkflowID string                 `json:"workflow_id"`
		Input      map[string]interface{} `json:"input"`
		Signature  string                 `json:"signature"`
		Signer     string                 `json:"signer"`
	} `json:"params"`
	ID int `json:"id"`
}

// JSON-RPC 응답 구조체
type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      int         `json:"id"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// AccessReport 결과
type AccessResult struct {
	AgentID          int    `json:"agentId"`
	Approved         bool   `json:"approved"`
	Tier             int    `json:"tier"`
	AccountableHuman string `json:"accountableHuman"`
	Reason           string `json:"reason"`
}

func main() {
	// Render에서 PORT 환경변수 사용
	_ = godotenv.Load("../.env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 9분마다 셀프 핑 시작
	startSelfPing()
	//
	go startEventListener()

	http.HandleFunc("/trigger", handleAccessRequest)
	http.HandleFunc("/health_check", handleHealth)
	http.HandleFunc("/api/plaid/proxy", handlePlaidProxy)

	addr := "0.0.0.0:" + port
	fmt.Printf("CRE Wrapper Server running on %s\n", addr)
	fmt.Println("   Endpoints:")
	fmt.Println("   - POST /trigger (CRE workflow)")
	fmt.Println("   - GET  /health_check  (health check)")

	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}

func handleAccessRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// 1. POST만 허용
	if r.Method != http.MethodPost {
		sendError(w, -32600, "Method not allowed", 0)
		return
	}

	// 2. JSON-RPC 요청 파싱
	var req JSONRPCRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, -32700, "Parse error", 0)
		return
	}

	// 3. 메서드 확인
	if req.Method != "workflow_execute" {
		sendError(w, -32601, "Method not found", req.ID)
		return
	}

	// 4. Signer 검증 (authorizedKeys 체크)
	signerLower := strings.ToLower(req.Params.Signer)
	if !authorizedSigners[signerLower] {
		sendError(w, -32000, "Unauthorized signer", req.ID)
		return
	}

	// 5. 서명 검증
	if !verifySignature(req.Params.Input, req.Params.Signature, req.Params.Signer) {
		sendError(w, -32001, "Invalid signature", req.ID)
		return
	}

	// 5. agentId 추출
	agentIdFloat, ok := req.Params.Input["agentId"].(float64)
	if !ok {
		sendError(w, -32602, "Invalid params: agentId required", req.ID)
		return
	}
	agentId := int(agentIdFloat)

	// 6. CRE simulate 실행
	result, err := runCRESimulate(agentId)
	if err != nil {
		sendError(w, -32000, err.Error(), req.ID)
		return
	}
	// 7. 성공 응답
	json.NewEncoder(w).Encode(JSONRPCResponse{
		JSONRPC: "2.0",
		Result:  result,
		ID:      req.ID,
	})

}

// CRE 출력 파싱용 구조체
type CREOutput struct {
	AccountableHuman string `json:"AccountableHuman"` // base64
	AgentID          int    `json:"AgentID"`
	Approved         bool   `json:"Approved"`
	Reason           string `json:"Reason"` // base64
	Tier             int    `json:"Tier"`
}

func runCRESimulate(agentId int) (*AccessResult, error) {
	// JSON 페이로드 생성
	payload := fmt.Sprintf(`{"agentId": %d}`, agentId)
	tmpFile := "/tmp/cre_payload.json"
	os.WriteFile(tmpFile, []byte(payload), 0644)

	// 프로젝트 루트 (Docker에서는 /app, 로컬에서는 상위 디렉토리)
	projectRoot := os.Getenv("CRE_PROJECT_ROOT")
	if projectRoot == "" {
		projectRoot, _ = filepath.Abs("..")
	}

	// CRE 타겟 설정 (환경변수 또는 기본값)
	creTarget := os.Getenv("CRE_TARGET")
	if creTarget == "" {
		creTarget = "staging-settings"
	}

	// 디버그: 환경 체크
	fmt.Printf("DEBUG: projectRoot=%s, creTarget=%s\n", projectRoot, creTarget)
	fmt.Printf("DEBUG: payload file=%s\n", tmpFile)

	// 파일 존재 체크
	if _, err := os.Stat(projectRoot + "/project.yaml"); err != nil {
		fmt.Printf("DEBUG: project.yaml NOT FOUND: %v\n", err)
	} else {
		fmt.Println("DEBUG: project.yaml OK")
	}
	if _, err := os.Stat(projectRoot + "/secrets.yaml"); err != nil {
		fmt.Printf("DEBUG: secrets.yaml (symlink) NOT FOUND: %v\n", err)
	} else {
		fmt.Println("DEBUG: secrets.yaml OK")
	}
	if _, err := os.Stat("/root/.cre/cre.yaml"); err != nil {
		fmt.Printf("DEBUG: cre.yaml NOT FOUND: %v\n", err)
	} else {
		fmt.Println("DEBUG: cre.yaml OK")
	}

	// cre simulate 실행 (2분 타임아웃)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "cre", "workflow", "simulate",
		"whitewall-access",
		"--target", creTarget,
		"--http-payload", "@"+tmpFile,
		"--non-interactive",
		"--trigger-index", "0",
		"--broadcast",
	)
	cmd.Dir = projectRoot

	fmt.Println("CRE 시뮬레이션 시작...")
	output, err := cmd.CombinedOutput()
	outputStr := string(output)
	fmt.Printf("CRE 출력:\n%s\n", outputStr)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("CRE execution timed out after 2 minutes")
	}
	if err != nil {
		return nil, fmt.Errorf("CRE execution failed: %v", err)
	}

	// JSON 부분 추출 (Workflow Simulation Result: 이후)
	re := regexp.MustCompile(`Workflow Simulation Result:\s*(\{[\s\S]*?\})`)
	matches := re.FindStringSubmatch(outputStr)
	if len(matches) < 2 {
		return nil, fmt.Errorf("Failed to parse CRE output")
	}

	// JSON 파싱
	var creOutput CREOutput
	if err := json.Unmarshal([]byte(matches[1]), &creOutput); err != nil {
		return nil, fmt.Errorf("Failed to parse JSON: %v", err)
	}

	// base64 → hex 변환 (AccountableHuman)
	addrBytes, _ := base64.StdEncoding.DecodeString(creOutput.AccountableHuman)
	accountableHuman := "0x" + hex.EncodeToString(addrBytes)

	// base64 → hex 변환 (Reason)
	reasonBytes, _ := base64.StdEncoding.DecodeString(creOutput.Reason)
	reason := "0x" + hex.EncodeToString(reasonBytes)

	return &AccessResult{
		AgentID:          creOutput.AgentID,
		Approved:         creOutput.Approved,
		Tier:             creOutput.Tier,
		AccountableHuman: accountableHuman,
		Reason:           reason,
	}, nil
}

func sendError(w http.ResponseWriter, code int, message string, id int) {
	json.NewEncoder(w).Encode(JSONRPCResponse{
		JSONRPC: "2.0",
		Error: &RPCError{
			Code:    code,
			Message: message,
		},
		ID: id,
	})
}

func verifySignature(input map[string]interface{}, signature string, signer string) bool {
	// 1. 입력 데이터를 JSON으로 직렬화
	message, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Failed to marshal input")
		return false
	}

	// 2. Ethereum 서명 메시지 형식으로 해시
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	hash := crypto.Keccak256Hash([]byte(prefix + string(message)))

	// 3. 서명 디코딩
	sigBytes := common.FromHex(signature)
	if len(sigBytes) != 65 {
		fmt.Println("Invalid signature length")
		return false
	}

	// 4. recovery id 조정 (27, 28 -> 0, 1)
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}

	// 5. 서명에서 공개키 복구
	pubKey, err := crypto.SigToPub(hash.Bytes(), sigBytes)
	if err != nil {
		fmt.Println("Failed to recover public key:", err)
		return false
	}

	// 6. 공개키에서 주소 추출
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// 7. 예상 주소와 비교
	expectedAddr := common.HexToAddress(signer)
	isValid := recoveredAddr == expectedAddr

	fmt.Printf("서명 검증: recovered=%s, expected=%s, valid=%v\n",
		recoveredAddr.Hex(), expectedAddr.Hex(), isValid)

	return isValid
}

// =================
func startEventListener() {
	// 환경변수에서 WebSocket RPC URL 읽기
	wsRpcUrl := os.Getenv("RPC_URL")
	if wsRpcUrl == "" {
		fmt.Println("RPC_URL not set, skipping event listener")
		return
	}

	// ValidationRegistry 주소 확인
	if ValidationRegistryAddr == "" || ValidationRegistryAddr == "0x..." {
		fmt.Println("ValidationRegistryAddr not set, skipping event listener")
		return
	}

	fmt.Println("Starting ValidationRegistry event listener...")

	// 1. WebSocket 연결
	client, err := ethclient.Dial(wsRpcUrl)
	if err != nil {
		fmt.Printf("Failed to connect to WebSocket: %v\n", err)
		return
	}
	defer client.Close()

	// 2. ValidationRegistry 주소
	registryAddr := common.HexToAddress(ValidationRegistryAddr)

	// 3. 이벤트 필터 설정
	query := ethereum.FilterQuery{
		Addresses: []common.Address{registryAddr},
		Topics:    [][]common.Hash{{ValidationRequestTopic}},
	}

	// 4. 이벤트 구독
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Printf("Failed to subscribe to logs: %v\n", err)
		return
	}
	defer sub.Unsubscribe()

	fmt.Println("Listening for ValidationRequest events...")

	// 5. 이벤트 처리 루프
	for {
		select {
		case err := <-sub.Err():
			fmt.Printf("Subscription error: %v\n", err)
			return
		case vLog := <-logs:
			handleValidationEvent(vLog)
		}
	}
}
func handleValidationEvent(vLog types.Log) {
	txHash := vLog.TxHash.Hex()

	// indexed 필드에서 validator 주소 추출 (topic[1])
	validatorAddr := common.BytesToAddress(vLog.Topics[1].Bytes())

	fmt.Printf("ValidationRequest detected! txHash=%s, validator=%s\n",
		txHash, validatorAddr.Hex())

	// trigger-index 결정
	var triggerIndex string
	switch strings.ToLower(validatorAddr.Hex()) {
	case strings.ToLower(StripeKYCValidatorAddr):
		triggerIndex = "1" // KYC
	case strings.ToLower(PlaidCreditValidatorAddr):
		triggerIndex = "2" // Credit
	default:
		fmt.Println("Unknown validator, skipping")
		return
	}

	// CRE simulate 실행
	go runLogTriggerSimulate(txHash, triggerIndex)
}

func runLogTriggerSimulate(txHash string, triggerIndex string) {
	projectRoot := os.Getenv("CRE_PROJECT_ROOT")
	if projectRoot == "" {
		projectRoot, _ = filepath.Abs("..")
	}

	creTarget := os.Getenv("CRE_TARGET")
	if creTarget == "" {
		creTarget = "staging-settings"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "cre", "workflow", "simulate",
		"whitewall-access",
		"--target", creTarget,
		"--evm-tx-hash", txHash,
		"--evm-event-index", "0", //트잭에서  첫번째 이벤트 ㅇㅇ
		"--trigger-index", triggerIndex,
		"--non-interactive",
		"--broadcast",
	)
	cmd.Dir = projectRoot

	fmt.Printf("Running CRE simulate for Log trigger (txHash=%s, index=%s)\n", txHash, triggerIndex)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Log trigger simulation failed: %v\n", err)
	}
	fmt.Printf("Log trigger output:\n%s\n", string(output))
}
