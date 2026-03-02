package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"kicknroar/internal/config"
)

func TestJWTManager_GenerateAccessToken(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:        "test-secret-key",
		Expiry:        15 * time.Minute,
		RefreshExpiry: 7 * 24 * time.Hour,
	}
	manager := NewManager(cfg)

	token, err := manager.GenerateAccessToken("user-123", "test@example.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestJWTManager_GenerateRefreshToken(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:        "test-secret-key",
		Expiry:        15 * time.Minute,
		RefreshExpiry: 7 * 24 * time.Hour,
	}
	manager := NewManager(cfg)

	token, err := manager.GenerateRefreshToken("user-123")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestJWTManager_ValidateToken(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:        "test-secret-key",
		Expiry:        15 * time.Minute,
		RefreshExpiry: 7 * 24 * time.Hour,
	}
	manager := NewManager(cfg)

	// Generate and validate token
	token, err := manager.GenerateAccessToken("user-123", "test@example.com")
	assert.NoError(t, err)

	claims, err := manager.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "user-123", claims.UserID)
	assert.Equal(t, "test@example.com", claims.Email)
}

func TestJWTManager_ValidateToken_Invalid(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:        "test-secret-key",
		Expiry:        15 * time.Minute,
		RefreshExpiry: 7 * 24 * time.Hour,
	}
	manager := NewManager(cfg)

	// Invalid token
	_, err := manager.ValidateToken("invalid-token")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidToken, err)
}

func TestJWTManager_ValidateToken_WrongSecret(t *testing.T) {
	cfg1 := &config.JWTConfig{
		Secret: "secret-1",
		Expiry: 15 * time.Minute,
	}
	cfg2 := &config.JWTConfig{
		Secret: "secret-2",
		Expiry: 15 * time.Minute,
	}
	manager1 := NewManager(cfg1)
	manager2 := NewManager(cfg2)

	// Generate token with manager1
	token, err := manager1.GenerateAccessToken("user-123", "test@example.com")
	assert.NoError(t, err)

	// Try to validate with manager2 (different secret)
	_, err = manager2.ValidateToken(token)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidToken, err)
}

