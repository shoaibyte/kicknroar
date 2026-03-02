package handler

import (
	"net/http"
	"strconv"
	"time"

	"kicknroar/internal/api/middleware"
	"kicknroar/internal/dto/request"
	"kicknroar/internal/dto/response"
	"kicknroar/internal/pkg/validator"
	"kicknroar/internal/service"
	"kicknroar/internal/util"

	"github.com/labstack/echo/v4"
)

// MatchHandler handles match requests
type MatchHandler struct {
	matchService *service.MatchService
}

// NewMatchHandler creates a new match handler
func NewMatchHandler(matchService *service.MatchService) *MatchHandler {
	return &MatchHandler{
		matchService: matchService,
	}
}

// CreateMatch creates a new match
// @Summary      Create a new match
// @Description  Create a new football match with venue, date, time, and other details
// @Tags         matches
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        request  body      request.CreateMatchRequest  true  "Match creation request"
// @Success      201      {object}  object
// @Failure      400      {object}  response.ErrorResponse
// @Failure      401      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/matches [post]
func (h *MatchHandler) CreateMatch(c echo.Context) error {
	var req request.CreateMatchRequest
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

	creatorID := middleware.GetUserID(c)
	match, err := h.matchService.CreateMatch(c.Request().Context(), &req, creatorID)
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

	return c.JSON(http.StatusCreated, match)
}

// GetMatch gets a match by ID
// @Summary      Get match by ID
// @Description  Get detailed information about a specific match
// @Tags         matches
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string  true  "Match ID"
// @Success      200  {object}  object
// @Failure      401  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/v1/matches/{id} [get]
func (h *MatchHandler) GetMatch(c echo.Context) error {
	matchID := c.Param("id")
	match, err := h.matchService.GetMatch(c.Request().Context(), matchID)
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

	return c.JSON(http.StatusOK, match)
}

// ListMatches lists matches with filters
// @Summary      List matches
// @Description  Get a list of matches with optional filters (status, date range, pagination)
// @Tags         matches
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        status     query     string  false  "Filter by status (open, full, cancelled, completed)"
// @Param        date_from   query     string  false  "Filter matches from date (YYYY-MM-DD)"
// @Param        date_to     query     string  false  "Filter matches to date (YYYY-MM-DD)"
// @Param        limit       query     int     false  "Number of results per page (default: 20)"
// @Param        offset      query     int     false  "Number of results to skip (default: 0)"
// @Success      200         {object}  object
// @Failure      401         {object}  response.ErrorResponse
// @Failure      500         {object}  response.ErrorResponse
// @Router       /api/v1/matches [get]
func (h *MatchHandler) ListMatches(c echo.Context) error {
	filters := make(map[string]interface{})

	if status := c.QueryParam("status"); status != "" {
		filters["status"] = status
	}
	if dateFromStr := c.QueryParam("date_from"); dateFromStr != "" {
		if dateFrom, err := time.Parse("2006-01-02", dateFromStr); err == nil {
			filters["date_from"] = dateFrom
		}
	}
	if dateToStr := c.QueryParam("date_to"); dateToStr != "" {
		if dateTo, err := time.Parse("2006-01-02", dateToStr); err == nil {
			filters["date_to"] = dateTo
		}
	}

	limit := 20
	if limitStr := c.QueryParam("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	offset := 0
	if offsetStr := c.QueryParam("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	matches, err := h.matchService.ListMatches(c.Request().Context(), filters, limit, offset)
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": matches,
	})
}

// UpdateMatch updates a match
// @Summary      Update match
// @Description  Update match details (only creator can update)
// @Tags         matches
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id       path      string                    true  "Match ID"
// @Param        request  body      request.UpdateMatchRequest  true  "Match update request"
// @Success      200      {object}  object
// @Failure      400      {object}  response.ErrorResponse
// @Failure      401      {object}  response.ErrorResponse
// @Failure      403      {object}  response.ErrorResponse
// @Failure      404      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/matches/{id} [put]
func (h *MatchHandler) UpdateMatch(c echo.Context) error {
	matchID := c.Param("id")
	var req request.UpdateMatchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "Invalid request data",
			},
		})
	}

	match, err := h.matchService.UpdateMatch(c.Request().Context(), matchID, &req)
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

	return c.JSON(http.StatusOK, match)
}

// DeleteMatch deletes a match
// @Summary      Delete match
// @Description  Delete a match (only creator can delete)
// @Tags         matches
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string  true  "Match ID"
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  response.ErrorResponse
// @Failure      403  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/v1/matches/{id} [delete]
func (h *MatchHandler) DeleteMatch(c echo.Context) error {
	matchID := c.Param("id")
	err := h.matchService.DeleteMatch(c.Request().Context(), matchID)
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

	return c.JSON(http.StatusOK, map[string]string{"message": "Match deleted successfully"})
}

// JoinMatch joins a match
// @Summary      Join a match
// @Description  Join a match as a participant
// @Tags         matches
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string  true  "Match ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  response.ErrorResponse
// @Failure      401  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      409  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/v1/matches/{id}/join [post]
func (h *MatchHandler) JoinMatch(c echo.Context) error {
	matchID := c.Param("id")
	userID := middleware.GetUserID(c)

	err := h.matchService.JoinMatch(c.Request().Context(), matchID, userID)
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

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully joined match"})
}

// LeaveMatch leaves a match
// @Summary      Leave a match
// @Description  Leave a match that you previously joined
// @Tags         matches
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string  true  "Match ID"
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/v1/matches/{id}/leave [post]
func (h *MatchHandler) LeaveMatch(c echo.Context) error {
	matchID := c.Param("id")
	userID := middleware.GetUserID(c)

	err := h.matchService.LeaveMatch(c.Request().Context(), matchID, userID)
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

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully left match"})
}

// GetParticipants gets match participants
// @Summary      Get match participants
// @Description  Get list of all participants for a specific match
// @Tags         matches
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string  true  "Match ID"
// @Success      200  {object}  object
// @Failure      401  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/v1/matches/{id}/participants [get]
func (h *MatchHandler) GetParticipants(c echo.Context) error {
	matchID := c.Param("id")
	participants, err := h.matchService.GetParticipants(c.Request().Context(), matchID)
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": participants,
	})
}
