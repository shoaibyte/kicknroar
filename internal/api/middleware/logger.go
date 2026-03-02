package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger returns logger middleware
func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${status} ${method} ${uri} ${latency_human} ${error}` + "\n",
	})
}

// Recover returns recover middleware
func Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

