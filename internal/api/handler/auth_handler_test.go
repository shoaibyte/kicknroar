package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"kicknroar/internal/ent"
	"kicknroar/internal/pkg/jwt"
	"kicknroar/internal/repository"
	"kicknroar/internal/service"
	"kicknroar/internal/testutil"
	"kicknroar/internal/util"
)

func setupAuthTest(t *testing.T) (*AuthHandler, *ent.Client, func()) {
	client := testutil.SetupTestDB(t)
	userRepo := repository.NewUserRepository(client)
	
	cfg := testutil.TestConfig()
	jwtManager := jwt.NewManager(&cfg.JWT)
	authService := service.NewAuthService(userRepo, jwtManager)
	authHandler := NewAuthHandler(authService)

	cleanup := func() {
		testutil.CleanupTestDB(client)
	}

	return authHandler, client, cleanup
}

func TestAuthHandler_Signup(t *testing.T) {
	handler, _, cleanup := setupAuthTest(t)
	defer cleanup()

	e := echo.New()
	
	tests := []struct {
		name           string
		body           map[string]interface{}
		expectedStatus int
		expectToken    bool
	}{
		{
			name: "Valid signup",
			body: map[string]interface{}{
				"email":     "test@example.com",
				"password":  "SecurePass123!",
				"full_name": "Test User",
				"phone":     "+8801712345678",
			},
			expectedStatus: http.StatusCreated,
			expectToken:    true,
		},
		{
			name: "Invalid email",
			body: map[string]interface{}{
				"email":     "invalid-email",
				"password":  "SecurePass123!",
				"full_name": "Test User",
				"phone":     "+8801712345678",
			},
			expectedStatus: http.StatusBadRequest,
			expectToken:    false,
		},
		{
			name: "Missing fields",
			body: map[string]interface{}{
				"email": "test@example.com",
			},
			expectedStatus: http.StatusBadRequest,
			expectToken:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.Signup(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectToken {
				var response map[string]interface{}
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.NotEmpty(t, response["token"])
				assert.NotEmpty(t, response["user"])
			}
		})
	}
}

func TestAuthHandler_Signup_DuplicateEmail(t *testing.T) {
	handler, client, cleanup := setupAuthTest(t)
	defer cleanup()

	// Create existing user
	ctx := context.Background()
	_, _ = testutil.CreateTestUser(ctx, client, "existing@example.com", "hash", "Existing User", "+8801711111111")

	e := echo.New()
	body, _ := json.Marshal(map[string]interface{}{
		"email":     "existing@example.com",
		"password":  "SecurePass123!",
		"full_name": "New User",
		"phone":     "+8801722222222",
	})

	req := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.Signup(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusConflict, rec.Code)
}

func TestAuthHandler_Login(t *testing.T) {
	handler, client, cleanup := setupAuthTest(t)
	defer cleanup()

	ctx := context.Background()
	passwordHash, _ := util.HashPassword("SecurePass123!")
	_, _ = testutil.CreateTestUser(ctx, client, "test@example.com", passwordHash, "Test User", "+8801712345678")

	e := echo.New()
	
	tests := []struct {
		name           string
		body           map[string]interface{}
		expectedStatus int
		expectToken    bool
	}{
		{
			name: "Valid login",
			body: map[string]interface{}{
				"email":    "test@example.com",
				"password": "SecurePass123!",
			},
			expectedStatus: http.StatusOK,
			expectToken:    true,
		},
		{
			name: "Invalid password",
			body: map[string]interface{}{
				"email":    "test@example.com",
				"password": "WrongPassword",
			},
			expectedStatus: http.StatusUnauthorized,
			expectToken:    false,
		},
		{
			name: "Non-existent user",
			body: map[string]interface{}{
				"email":    "nonexistent@example.com",
				"password": "SecurePass123!",
			},
			expectedStatus: http.StatusUnauthorized,
			expectToken:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.Login(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectToken {
				var response map[string]interface{}
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.NotEmpty(t, response["token"])
			}
		})
	}
}

