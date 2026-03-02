package middleware

import (
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"kicknroar/internal/config"
	"kicknroar/internal/util"
)

// RateLimiter implements a simple in-memory rate limiter
type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.Mutex
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}

	// Cleanup old entries periodically
	go rl.cleanup()

	return rl
}

// cleanup removes old entries periodically
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for key, times := range rl.requests {
			filtered := []time.Time{}
			for _, t := range times {
				if now.Sub(t) < rl.window {
					filtered = append(filtered, t)
				}
			}
			if len(filtered) == 0 {
				delete(rl.requests, key)
			} else {
				rl.requests[key] = filtered
			}
		}
		rl.mu.Unlock()
	}
}

// Allow checks if a request is allowed
func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	times := rl.requests[key]

	// Remove old entries
	filtered := []time.Time{}
	for _, t := range times {
		if now.Sub(t) < rl.window {
			filtered = append(filtered, t)
		}
	}

	if len(filtered) >= rl.limit {
		return false
	}

	filtered = append(filtered, now)
	rl.requests[key] = filtered
	return true
}

// RateLimit returns rate limiting middleware
func RateLimit(cfg *config.RateLimitConfig) echo.MiddlewareFunc {
	limiter := NewRateLimiter(cfg.RequestsPerMinute, 1*time.Minute)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.RealIP()
			if !limiter.Allow(key) {
				return c.JSON(429, util.NewAppError(
					util.ErrorCodeInternalServer,
					"Too many requests",
					429,
				))
			}
			return next(c)
		}
	}
}

// AuthRateLimit returns rate limiting middleware for auth endpoints
func AuthRateLimit(cfg *config.RateLimitConfig) echo.MiddlewareFunc {
	limiter := NewRateLimiter(cfg.AuthRequestsPerMinute, 1*time.Minute)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.RealIP()
			if !limiter.Allow(key) {
				return c.JSON(429, util.NewAppError(
					util.ErrorCodeInternalServer,
					"Too many requests",
					429,
				))
			}
			return next(c)
		}
	}
}

// UploadRateLimit returns rate limiting middleware for upload endpoints
func UploadRateLimit(cfg *config.RateLimitConfig) echo.MiddlewareFunc {
	limiter := NewRateLimiter(cfg.UploadRequestsPerMinute, 1*time.Minute)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.RealIP()
			if !limiter.Allow(key) {
				return c.JSON(429, util.NewAppError(
					util.ErrorCodeInternalServer,
					"Too many requests",
					429,
				))
			}
			return next(c)
		}
	}
}

