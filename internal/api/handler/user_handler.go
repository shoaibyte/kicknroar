package handler

import (
	"errors"
	"net/http"

	"kicknroar/internal/api/middleware"
	"kicknroar/internal/dto/request"
	"kicknroar/internal/dto/response"
	"kicknroar/internal/pkg/validator"
	"kicknroar/internal/service"
	"kicknroar/internal/util"

	"github.com/labstack/echo/v4"
)

// UserHandler handles user requests
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetMe gets current user profile
// @Summary      Get current user profile
// @Description  Get the authenticated user's profile information
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Success      200      {object}  response.UserResponse
// @Failure      401      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/users/me [get]
func (h *UserHandler) GetMe(c echo.Context) error {
	userID := middleware.GetUserID(c)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeUnauthorized),
				Message: "Unauthorized",
			},
		})
	}

	user, err := h.userService.GetProfile(c.Request().Context(), userID)
	if err != nil {
		var appErr *util.AppError
		if errors.As(err, &appErr) {
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

	return c.JSON(http.StatusOK, user)
}

// UpdateMe updates current user profile
// @Summary      Update current user profile
// @Description  Update the authenticated user's profile information (full_name, skill_level, profile_image_url, preferred_locations)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        request  body      request.UpdateUserRequest  true  "Profile update request"
// @Success      200      {object}  response.UserResponse
// @Failure      400      {object}  response.ErrorResponse
// @Failure      401      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/users/me [put]
func (h *UserHandler) UpdateMe(c echo.Context) error {
	userID := middleware.GetUserID(c)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeUnauthorized),
				Message: "Unauthorized",
			},
		})
	}

	var req request.UpdateUserRequest
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

	// Convert DTO to map for service layer
	updates := make(map[string]interface{})
	if req.FullName != nil {
		updates["full_name"] = *req.FullName
	}
	if req.SkillLevel != nil {
		updates["skill_level"] = *req.SkillLevel
	}
	if req.ProfileImageURL != nil {
		updates["profile_image_url"] = *req.ProfileImageURL
	}
	if req.PreferredLocations != nil {
		updates["preferred_locations"] = req.PreferredLocations
	}

	user, err := h.userService.UpdateProfile(c.Request().Context(), userID, updates)
	if err != nil {
		var appErr *util.AppError
		if errors.As(err, &appErr) {
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

	return c.JSON(http.StatusOK, user)
}

// GetUser gets a user by ID
// @Summary      Get user by ID
// @Description  Get a user's profile information by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  response.UserResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      401  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "User ID is required",
			},
		})
	}

	user, err := h.userService.GetProfile(c.Request().Context(), userID)
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

	return c.JSON(http.StatusOK, user)
}

// GetUserStats gets user statistics
// @Summary      Get user statistics
// @Description  Get statistics for a specific user (matches played, wins, etc.)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  object
// @Failure      400  {object}  response.ErrorResponse
// @Failure      401  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/v1/users/{id}/stats [get]
func (h *UserHandler) GetUserStats(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "User ID is required",
			},
		})
	}

	stats, err := h.userService.GetUserStats(c.Request().Context(), userID)
	if err != nil {
		var appErr *util.AppError
		if errors.As(err, &appErr) {
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

	return c.JSON(http.StatusOK, stats)
}
