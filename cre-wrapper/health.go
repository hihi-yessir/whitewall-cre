package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// HealthResponse - health check 응답 구조체
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// handleHealth - Render health check 엔드포인트
// Render가 주기적으로 이 엔드포인트를 호출해서 서버가 살아있는지 확인함
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

// startSelfPing - 9분마다 자기 자신에게 핑 보내기
// Render Free Plan은 15분 동안 요청이 없으면 sleep 모드로 전환됨
// 9분마다 핑을 보내서 sleep 방지!
func startSelfPing() {
	// 서버 URL 가져오기 (Render가 자동으로 설정해줌)
	renderURL := os.Getenv("RENDER_EXTERNAL_URL")
	if renderURL == "" {
		fmt.Println("RENDER_EXTERNAL_URL not set, self-ping disabled")
		return
	}

	healthURL := renderURL + "/health_check"
	fmt.Printf("Self-ping enabled: %s (every 9 minutes)\n", healthURL)

	// 고루틴으로 백그라운드에서 실행
	go func() {
		ticker := time.NewTicker(9 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			resp, err := http.Get(healthURL)
			if err != nil {
				fmt.Printf("Self-ping failed: %v\n", err)
				continue
			}
			resp.Body.Close()
			fmt.Printf("Self-ping success at %s\n", time.Now().UTC().Format("15:04:05"))
		}
	}()
}
