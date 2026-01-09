package session

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"sync"
)

type contextKey string

const (
	userContextKey contextKey = "user"
)

var ts TokenStore

type TokenStore struct {
	rwMutex sync.RWMutex
	store   map[string]TokenInfo
}

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
	ts = TokenStore{
		store: make(map[string]TokenInfo, 1),
	}

}

func generateToken(tokenInfo TokenInfo) string {
	tokenJSON, _ := json.Marshal(tokenInfo)

	hash := sha256.Sum256(tokenJSON)
	return hex.EncodeToString(hash[:])
}

func AddToken(tokenInfo TokenInfo) string {
	ts.rwMutex.Lock()
	defer ts.rwMutex.Unlock()

	token := generateToken(tokenInfo)
	ts.store[token] = tokenInfo

	return token
}

func GetToken(token string) (TokenInfo, bool) {
	ts.rwMutex.RLock()
	defer ts.rwMutex.RUnlock()

	tokenInfo, exists := ts.store[token]

	return tokenInfo, exists
}

func RemoveToken(token string) {
	ts.rwMutex.Lock()
	defer ts.rwMutex.Unlock()

	delete(ts.store, token)
}

func GetTokenFromSessionID(sessionID string) (TokenInfo, bool) {
	ts.rwMutex.RLock()
	defer ts.rwMutex.RUnlock()

	for _, tokenInfo := range ts.store {
		if tokenInfo.ID == sessionID {
			return tokenInfo, true
		}
	}

	return TokenInfo{}, false
}

func AuthMiddleware(bmcUser, bmcPassword string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()

			if ok {
				if user != bmcUser || pass != bmcPassword {
					w.Header().Set("WWW-Authenticate", `Basic realm="Redfish"`)
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				ctx := context.WithValue(r.Context(), userContextKey, user)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("X-Auth-Token")
			if token == "" {
				w.Header().Set("WWW-Authenticate", `Basic realm="Redfish"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ts.rwMutex.RLock()
			_, exists := ts.store[token]
			ts.rwMutex.RUnlock()

			if !exists {
				w.Header().Set("WWW-Authenticate", `Basic realm="Redfish"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
