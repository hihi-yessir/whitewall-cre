package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	http.HandleFunc("/cre/access", handleAccessRequest)

	fmt.Println("CRE Wrapper Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleAccessRequest(w http.ResponseWriter, r *http.Request) {
	// 1. POST만 허용
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. 요청 Body 읽기
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 3. JSON 파일로 저장 (cre simulate가 읽을 수 있게)
	payloadBytes, _ := json.Marshal(payload)
	tmpFile := "/tmp/cre_payload.json"
	os.WriteFile(tmpFile, payloadBytes, 0644)

	// 4. 프로젝트 루트 경로 (cre-wrapper 상위 폴더)
	projectRoot, _ := filepath.Abs("..")

	// 5. cre simulate 실행
	cmd := exec.Command("cre", "workflow", "simulate",
		"whitewall-access",
		"--target", "staging-settings",
		"--http-payload", "@"+tmpFile,
		"--non-interactive",
		"--trigger-index", "0",
	)
	cmd.Dir = projectRoot

	output, err := cmd.CombinedOutput()

	//결과 반환
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Printf("CRE Error: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"output":  string(output),
		})
		return
	}

	fmt.Printf("CRE Success\n")
	fmt.Printf("Output: %s\n", string(output))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"output":  string(output),
	})
}
