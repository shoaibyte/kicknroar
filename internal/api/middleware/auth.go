package middleware

import (
	"net/http"
	"strings"

	"kicknroar/internal/pkg/jwt"
	"kicknroar/internal/util"

	"github.com/labstack/echo/v4"
)

// Auth returns authentication middleware
func Auth(jwtManager *jwt.Manager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, util.ErrUnauthorized())
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, util.ErrUnauthorized())
			}

			token := parts[1]
			claims, err := jwtManager.ValidateToken(token)
			if err != nil {
				if err == jwt.ErrExpiredToken {
					return c.JSON(http.StatusUnauthorized, util.ErrTokenExpired())
				}
				return c.JSON(http.StatusUnauthorized, util.ErrUnauthorized())
			}

			// Store user info in context
			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)

			return next(c)
		}
	}
}

// GetUserID extracts user ID from context
func GetUserID(c echo.Context) string {
	userID, ok := c.Get("user_id").(string)
	if !ok {
		return ""
	}
	return userID
}

// GetEmail extracts email from context
func GetEmail(c echo.Context) string {
	email, ok := c.Get("email").(string)
	if !ok {
		return ""
	}
	return email
}
