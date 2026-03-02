package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"kicknroar/internal/config"
	"kicknroar/internal/pkg/jwt"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:        "test-secret-key-for-testing-only-change-in-production",
		Expiry:        15 * time.Minute,
		RefreshExpiry: 7 * 24 * time.Hour,
	}
	jwtManager := jwt.NewManager(cfg)

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
		expectUserID   bool
	}{
		{
			name:           "Valid token",
			authHeader:     "Bearer " + generateTestToken(t, jwtManager, "user-123", "test@example.com"),
			expectedStatus: http.StatusOK,
			expectUserID:   true,
		},
		{
			name:           "Missing Authorization header",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			expectUserID:   false,
		},
		{
			name:           "Invalid token format",
			authHeader:     "InvalidFormat token",
			expectedStatus: http.StatusUnauthorized,
			expectUserID:   false,
		},
		{
			name:           "Expired token",
			authHeader:     "Bearer expired-token",
			expectedStatus: http.StatusUnauthorized,
			expectUserID:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			handler := Auth(jwtManager)
			handler(func(c echo.Context) error {
				if tt.expectUserID {
					userID := GetUserID(c)
					assert.NotEmpty(t, userID)
				}
				return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
			})(c)

			// Check response code (Echo middleware returns errors as JSON responses)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func generateTestToken(t *testing.T, jwtManager *jwt.Manager, userID, email string) string {
	token, err := jwtManager.GenerateAccessToken(userID, email)
	if err != nil {
		t.Fatalf("Failed to generate test token: %v", err)
	}
	return token
}
