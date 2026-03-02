package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"kicknroar/internal/config"
)

func TestCORS(t *testing.T) {
	cfg := &config.ServerConfig{
		AllowedOrigins: []string{"http://localhost:3000", "https://example.com"},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	req.Header.Set("Access-Control-Request-Method", "POST")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := CORS(cfg)
	err := handler(func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})(c)

	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:3000", rec.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", rec.Header().Get("Access-Control-Allow-Credentials"))
}

func TestCORS_InvalidOrigin(t *testing.T) {
	cfg := &config.ServerConfig{
		AllowedOrigins: []string{"http://localhost:3000"},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Origin", "http://malicious.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := CORS(cfg)
	err := handler(func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})(c)

	assert.NoError(t, err)
	// CORS middleware should not set Allow-Origin for invalid origins
	assert.Empty(t, rec.Header().Get("Access-Control-Allow-Origin"))
}

