package session

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

var tokenStore map[string]TokenInfo

type TokenInfo struct {
	ID       string
	Username string
}

func NewTokenInfo(id, username string) TokenInfo {
	return TokenInfo{
		ID:       id,
		Username: username,
	}
}

func init() {
	tokenStore = make(map[string]TokenInfo, 1)
}

func generateToken(tokenInfo TokenInfo) string {
	tokenJSON, _ := json.Marshal(tokenInfo)

	hash := sha256.Sum256(tokenJSON)
	return hex.EncodeToString(hash[:])
}

func AddToken(tokenInfo TokenInfo) string {
	token := generateToken(tokenInfo)
	tokenStore[token] = tokenInfo
	return token
}

func GetToken(token string) (TokenInfo, bool) {
	tokenInfo, exists := tokenStore[token]
	return tokenInfo, exists
}

func RemoveToken(token string) {
	delete(tokenStore, token)
}

func GetTokenFromID(id string) (TokenInfo, bool) {
	for _, tokenInfo := range tokenStore {
		if tokenInfo.ID == id {
			return tokenInfo, true
		}
	}
	return TokenInfo{}, false
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if _, exists := tokenStore[token]; !exists {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
