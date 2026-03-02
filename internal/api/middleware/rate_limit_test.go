package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"kicknroar/internal/config"
)

func TestRateLimit(t *testing.T) {
	cfg := &config.RateLimitConfig{
		RequestsPerMinute: 2, // Low limit for testing
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := RateLimit(cfg)

	// First request should succeed
	err := handler(func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Second request should succeed
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req, rec2)
	err = handler(func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})(c2)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec2.Code)

	// Third request should be rate limited
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req, rec3)
	err = handler(func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})(c3)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusTooManyRequests, rec3.Code)
}

func TestAuthRateLimit(t *testing.T) {
	cfg := &config.RateLimitConfig{
		AuthRequestsPerMinute: 2,
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := AuthRateLimit(cfg)

	// First request should succeed
	err := handler(func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Second request should succeed
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req, rec2)
	err = handler(func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})(c2)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec2.Code)

	// Third request should be rate limited
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req, rec3)
	err = handler(func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})(c3)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusTooManyRequests, rec3.Code)
}

func TestRateLimiterCleanup(t *testing.T) {
	limiter := NewRateLimiter(2, 100*time.Millisecond)

	key := "test-key"
	assert.True(t, limiter.Allow(key))
	assert.True(t, limiter.Allow(key))
	assert.False(t, limiter.Allow(key)) // Should be rate limited

	// Wait for cleanup
	time.Sleep(150 * time.Millisecond)

	// Should be able to make requests again after cleanup
	assert.True(t, limiter.Allow(key))
}

