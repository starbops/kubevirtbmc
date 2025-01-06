package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddAndGetToken(t *testing.T) {
	testCases := []struct {
		id       string
		username string
	}{
		{"123", "user1"},
		{"456", "user2"},
		{"789", "user3"},
	}

	for _, tc := range testCases {
		tokenInfo := NewTokenInfo(tc.id, tc.username)
		token := AddToken(tokenInfo)

		fetchedTokenInfo, exists := GetToken(token)
		if !exists {
			t.Errorf("Could not retrieve token for ID: %s", tc.id)
		}
		if fetchedTokenInfo != tokenInfo {
			t.Errorf("Expected %v, got %v", tokenInfo, fetchedTokenInfo)
		}
	}
}

func TestRemoveToken(t *testing.T) {
	testCases := []struct {
		id       string
		username string
	}{
		{"321", "user4"},
		{"654", "user5"},
	}

	for _, tc := range testCases {
		tokenInfo := NewTokenInfo(tc.id, tc.username)
		token := AddToken(tokenInfo)
		RemoveToken(token)

		if _, exists := ts.store[token]; exists {
			t.Errorf("Token was not removed for ID: %s", tc.id)
		}
	}
}

func TestGetTokenFromSessionID(t *testing.T) {
	testCases := []struct {
		sessionID string
		username  string
	}{
		{"101", "user6"},
		{"202", "user7"},
	}

	for _, tc := range testCases {
		tokenInfo := NewTokenInfo(tc.sessionID, tc.username)
		AddToken(tokenInfo)

		fetchedTokenInfo, exists := GetTokenFromSessionID(tc.sessionID)
		if !exists {
			t.Errorf("Could not retrieve token by session ID: %s", tc.sessionID)
		}
		if fetchedTokenInfo != tokenInfo {
			t.Errorf("Expected %v, got %v", tokenInfo, fetchedTokenInfo)
		}
	}
}

func TestAuthMiddleware(t *testing.T) {
	tokenInfo := NewTokenInfo("303", "user8")
	token := AddToken(tokenInfo)

	testCases := []struct {
		name       string
		token      string
		wantStatus int
	}{
		{"ValidToken", token, http.StatusOK},
		{"MissingToken", "", http.StatusUnauthorized},
		{"InvalidToken", "invalidToken", http.StatusUnauthorized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			authMiddleware := AuthMiddleware(nextHandler)

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("X-Auth-Token", tc.token)

			rr := httptest.NewRecorder()
			authMiddleware.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.wantStatus {
				t.Errorf("AuthMiddleware returned wrong status code: got %v want %v", status, tc.wantStatus)
			}
		})
	}
}
