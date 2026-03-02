package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// =================== Plaid Proxy ===================
// CRE workflow calls this proxy with confidential http
// CRE sends client_id, secret via template injection from DON Vault
//
// Flow:
// 1. CRE confidential http → body에 {{.PLAID_CLIENT_ID}}, {{.PLAID_SECRET}} (TEE가 치환)
// 2. Proxy receives actual values (already substituted by TEE)
// 3. Exchange public_token → access_token
// 4. Get balance info from Plaid
// 5. Calculate credit score
// 6. Return score (no encryption) → CRE handles DON signing

type PlaidProxyRequest struct {
	ClientID    string `json:"clientId"`    // from DON Vault via template injection
	Secret      string `json:"secret"`      // from DON Vault via template injection
	PublicToken string `json:"publicToken"` // from ValidationRegistry event
	AgentID     string `json:"agentId"`
	RequestHash string `json:"requestHash"`
}

type PlaidProxyResponse struct {
	Score   uint8  `json:"score"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// Plaid API responses
type PlaidTokenExchangeResponse struct {
	AccessToken string      `json:"access_token"`
	ItemID      string      `json:"item_id"`
	RequestID   string      `json:"request_id"`
	Error       *PlaidError `json:"error,omitempty"`
}

type PlaidBalanceResponse struct {
	Accounts  []PlaidAccount `json:"accounts"`
	Item      interface{}    `json:"item"`
	RequestID string         `json:"request_id"`
	Error     *PlaidError    `json:"error,omitempty"`
}

type PlaidAccount struct {
	AccountID string        `json:"account_id"`
	Balances  PlaidBalances `json:"balances"`
	Name      string        `json:"name"`
	Type      string        `json:"type"`
	Subtype   string        `json:"subtype"`
}

type PlaidBalances struct {
	Available              *float64 `json:"available"`
	Current                *float64 `json:"current"`
	Limit                  *float64 `json:"limit"`
	IsoCurrencyCode        string   `json:"iso_currency_code"`
	UnofficialCurrencyCode string   `json:"unofficial_currency_code"`
}

type PlaidError struct {
	ErrorType      string `json:"error_type"`
	ErrorCode      string `json:"error_code"`
	ErrorMessage   string `json:"error_message"`
	DisplayMessage string `json:"display_message"`
}

const PlaidSandboxURL = "https://sandbox.plaid.com"

func handlePlaidProxy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(PlaidProxyResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	var req PlaidProxyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(PlaidProxyResponse{
			Success: false,
			Error:   "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Printf("Plaid proxy request: agentId=%s, requestHash=%s\n", req.AgentID, req.RequestHash)

	// Validate required fields
	if req.PublicToken == "" || req.ClientID == "" || req.Secret == "" {
		json.NewEncoder(w).Encode(PlaidProxyResponse{
			Success: false,
			Error:   "Missing required fields: publicToken, clientId, secret",
		})
		return
	}

	// 1. Exchange public_token → access_token
	accessToken, err := exchangePublicToken(req.ClientID, req.Secret, req.PublicToken)
	if err != nil {
		fmt.Printf("Token exchange error: %v\n", err)
		json.NewEncoder(w).Encode(PlaidProxyResponse{
			Success: false,
			Error:   "Token exchange failed: " + err.Error(),
		})
		return
	}
	fmt.Printf("Token exchange successful, access_token obtained\n")

	// 2. Get balance info
	balanceResp, err := getAccountBalance(req.ClientID, req.Secret, accessToken)
	if err != nil {
		fmt.Printf("Balance fetch error: %v\n", err)
		json.NewEncoder(w).Encode(PlaidProxyResponse{
			Success: false,
			Error:   "Balance fetch failed: " + err.Error(),
		})
		return
	}
	fmt.Printf("Balance fetch successful, accounts=%d\n", len(balanceResp.Accounts))

	// 3. Calculate credit score
	score := computeCreditScore(balanceResp.Accounts)
	fmt.Printf("Credit score calculated: %d for agentId=%s\n", score, req.AgentID)

	// 4. Return score (no encryption - CRE will handle DON signing)
	json.NewEncoder(w).Encode(PlaidProxyResponse{
		Score:   score,
		Success: true,
	})
}

// exchangePublicToken: public_token → access_token
func exchangePublicToken(clientID, secret, publicToken string) (string, error) {
	payload := map[string]string{
		"client_id":    clientID,
		"secret":       secret,
		"public_token": publicToken,
	}
	payloadBytes, _ := json.Marshal(payload)

	resp, err := http.Post(
		PlaidSandboxURL+"/item/public_token/exchange",
		"application/json",
		bytes.NewReader(payloadBytes),
	)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var tokenResp PlaidTokenExchangeResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("parse error: %w", err)
	}

	if tokenResp.Error != nil {
		return "", fmt.Errorf("plaid error: %s - %s",
			tokenResp.Error.ErrorCode, tokenResp.Error.ErrorMessage)
	}

	return tokenResp.AccessToken, nil
}

// getAccountBalance: access_token으로 계좌 잔액 조회
func getAccountBalance(clientID, secret, accessToken string) (*PlaidBalanceResponse, error) {
	payload := map[string]string{
		"client_id":    clientID,
		"secret":       secret,
		"access_token": accessToken,
	}
	payloadBytes, _ := json.Marshal(payload)

	resp, err := http.Post(
		PlaidSandboxURL+"/accounts/balance/get",
		"application/json",
		bytes.NewReader(payloadBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var balanceResp PlaidBalanceResponse
	if err := json.Unmarshal(body, &balanceResp); err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	if balanceResp.Error != nil {
		return nil, fmt.Errorf("plaid error: %s - %s",
			balanceResp.Error.ErrorCode, balanceResp.Error.ErrorMessage)
	}

	return &balanceResp, nil
}

// computeCreditScore: Plaid 계좌 정보로 신용 점수 계산 (0-100)
// - loan, credit 타입은 부채로 처리
// - depository, investment 타입은 자산으로 처리
// - 순자산(Net Worth) 기반 점수 계산
func computeCreditScore(accounts []PlaidAccount) uint8 {
	if len(accounts) == 0 {
		return 0
	}

	var score uint8 = 0
	var totalAssets float64 = 0
	var totalLiabilities float64 = 0
	hasNegative := false

	// 1. Account count score (max 10)
	if len(accounts) >= 3 {
		score += 10
	} else {
		score += uint8(len(accounts) * 3)
	}

	for _, acc := range accounts {
		if acc.Balances.Current == nil {
			continue
		}
		current := *acc.Balances.Current

		// loan, credit → 부채 (빚)
		if acc.Type == "loan" || acc.Type == "credit" {
			totalLiabilities += current
		} else {
			// depository, investment → 자산
			totalAssets += current
			if current < 0 {
				hasNegative = true
			}
		}
	}

	// 2. Net worth score (max 40)
	netWorth := totalAssets - totalLiabilities
	if netWorth >= 50000 {
		score += 40
	} else if netWorth >= 20000 {
		score += 30
	} else if netWorth >= 5000 {
		score += 20
	} else if netWorth > 0 {
		score += 10
	}
	// netWorth <= 0 이면 0점

	// 3. No negative balance score (30)
	if !hasNegative {
		score += 30
	}

	// 4. Diversity score (max 20)
	typeSet := make(map[string]bool)
	for _, acc := range accounts {
		if acc.Type != "" {
			typeSet[acc.Type] = true
		}
	}
	diversityScore := uint8(len(typeSet) * 5)
	if diversityScore > 20 {
		diversityScore = 20
	}
	score += diversityScore

	return score
}
