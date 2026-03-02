package handler

import (
	"net/http"

	"kicknroar/internal/dto/request"
	"kicknroar/internal/dto/response"
	"kicknroar/internal/pkg/validator"
	"kicknroar/internal/service"
	"kicknroar/internal/util"

	"github.com/labstack/echo/v4"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Signup handles user registration
// @Summary      Register a new user
// @Description  Create a new user account with email, password, full name, and phone
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      request.SignupRequest  true  "Signup request"
// @Success      201      {object}  response.AuthResponse
// @Failure      400      {object}  response.ErrorResponse
// @Failure      409      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/auth/signup [post]
func (h *AuthHandler) Signup(c echo.Context) error {
	var req request.SignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "Invalid request data",
			},
		})
	}

	if err := validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: err.Error(),
			},
		})
	}

	authResp, err := h.authService.Signup(c.Request().Context(), &req)
	if err != nil {
		if appErr, ok := err.(*util.AppError); ok {
			return c.JSON(appErr.Status, response.ErrorResponse{
				Error: response.ErrorDetail{
					Code:    string(appErr.Code),
					Message: appErr.Message,
					Details: appErr.Details,
				},
			})
		}
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeInternalServer),
				Message: "Internal server error",
			},
		})
	}

	return c.JSON(http.StatusCreated, authResp)
}

// Login handles user login
// @Summary      Login user
// @Description  Authenticate user with email and password, returns access token and refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      request.LoginRequest  true  "Login request"
// @Success      200      {object}  response.AuthResponse
// @Failure      400      {object}  response.ErrorResponse
// @Failure      401      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "Invalid request data",
			},
		})
	}

	if err := validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: err.Error(),
			},
		})
	}

	authResp, err := h.authService.Login(c.Request().Context(), &req)
	if err != nil {
		if appErr, ok := err.(*util.AppError); ok {
			return c.JSON(appErr.Status, response.ErrorResponse{
				Error: response.ErrorDetail{
					Code:    string(appErr.Code),
					Message: appErr.Message,
					Details: appErr.Details,
				},
			})
		}
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeInternalServer),
				Message: "Internal server error",
			},
		})
	}

	return c.JSON(http.StatusOK, authResp)
}

// Refresh handles token refresh
// @Summary      Refresh access token
// @Description  Get a new access token using a valid refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      request.RefreshRequest  true  "Refresh token request"
// @Success      200      {object}  response.AuthResponse
// @Failure      400      {object}  response.ErrorResponse
// @Failure      401      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/auth/refresh [post]
func (h *AuthHandler) Refresh(c echo.Context) error {
	var req request.RefreshRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "Invalid request data",
			},
		})
	}

	if err := validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: err.Error(),
			},
		})
	}

	authResp, err := h.authService.Refresh(c.Request().Context(), req.RefreshToken)
	if err != nil {
		if appErr, ok := err.(*util.AppError); ok {
			return c.JSON(appErr.Status, response.ErrorResponse{
				Error: response.ErrorDetail{
					Code:    string(appErr.Code),
					Message: appErr.Message,
					Details: appErr.Details,
				},
			})
		}
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeInternalServer),
				Message: "Internal server error",
			},
		})
	}

	return c.JSON(http.StatusOK, authResp)
}

// Logout handles user logout
// @Summary      Logout user
// @Description  Logout the current user and invalidate their session
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Success      200      {object}  map[string]string
// @Failure      401      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/auth/logout [post]
func (h *AuthHandler) Logout(c echo.Context) error {
	userID := c.Get("user_id").(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeUnauthorized),
				Message: "Unauthorized",
			},
		})
	}

	err := h.authService.Logout(c.Request().Context(), userID)
	if err != nil {
		if appErr, ok := err.(*util.AppError); ok {
			return c.JSON(appErr.Status, response.ErrorResponse{
				Error: response.ErrorDetail{
					Code:    string(appErr.Code),
					Message: appErr.Message,
				},
			})
		}
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeInternalServer),
				Message: "Internal server error",
			},
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
