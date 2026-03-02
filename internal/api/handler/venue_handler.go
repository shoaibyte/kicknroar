package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"kicknroar/internal/api/middleware"
	"kicknroar/internal/dto/request"
	"kicknroar/internal/dto/response"
	"kicknroar/internal/pkg/validator"
	"kicknroar/internal/service"
	"kicknroar/internal/util"
)

// VenueHandler handles venue requests
type VenueHandler struct {
	venueService *service.VenueService
}

// NewVenueHandler creates a new venue handler
func NewVenueHandler(venueService *service.VenueService) *VenueHandler {
	return &VenueHandler{
		venueService: venueService,
	}
}

// CreateVenue creates a new venue
// @Summary      Create a new venue
// @Description  Create a new football venue with location, facilities, and other details
// @Tags         venues
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        request  body      request.CreateVenueRequest  true  "Venue creation request"
// @Success      201      {object}  object
// @Failure      400      {object}  response.ErrorResponse
// @Failure      401      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/venues [post]
func (h *VenueHandler) CreateVenue(c echo.Context) error {
	var req request.CreateVenueRequest
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

	ownerID := middleware.GetUserID(c)
	venue, err := h.venueService.CreateVenue(c.Request().Context(), &req, ownerID)
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

	return c.JSON(http.StatusCreated, venue)
}

// GetVenue gets a venue by ID
// @Summary      Get venue by ID
// @Description  Get detailed information about a specific venue
// @Tags         venues
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string  true  "Venue ID"
// @Success      200  {object}  object
// @Failure      400  {object}  response.ErrorResponse
// @Failure      401  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/v1/venues/{id} [get]
func (h *VenueHandler) GetVenue(c echo.Context) error {
	venueID := c.Param("id")
	if venueID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "Venue ID is required",
			},
		})
	}

	venue, err := h.venueService.GetVenue(c.Request().Context(), venueID)
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

	return c.JSON(http.StatusOK, venue)
}

// UpdateVenue updates a venue
// @Summary      Update venue
// @Description  Update venue details (only owner can update)
// @Tags         venues
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id       path      string                    true  "Venue ID"
// @Param        request  body      request.UpdateVenueRequest  true  "Venue update request"
// @Success      200      {object}  object
// @Failure      400      {object}  response.ErrorResponse
// @Failure      401      {object}  response.ErrorResponse
// @Failure      403      {object}  response.ErrorResponse
// @Failure      404      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/v1/venues/{id} [put]
func (h *VenueHandler) UpdateVenue(c echo.Context) error {
	venueID := c.Param("id")
	if venueID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "Venue ID is required",
			},
		})
	}

	var req request.UpdateVenueRequest
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

	venue, err := h.venueService.UpdateVenue(c.Request().Context(), venueID, &req)
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

	return c.JSON(http.StatusOK, venue)
}

// FindNearby finds nearby venues
// @Summary      Find nearby venues
// @Description  Find venues within a specified radius from given coordinates
// @Tags         venues
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        lat        query     float64  true   "Latitude"
// @Param        lng        query     float64  true   "Longitude"
// @Param        radius     query     float64  false  "Search radius in kilometers (default: 5)"
// @Param        field_type  query     string   false  "Filter by field type (futsal, football, astro)"
// @Param        limit      query     int      false  "Maximum number of results (default: 20)"
// @Success      200        {object}  object
// @Failure      400        {object}  response.ErrorResponse
// @Failure      401        {object}  response.ErrorResponse
// @Failure      500        {object}  response.ErrorResponse
// @Router       /api/v1/venues/nearby [get]
func (h *VenueHandler) FindNearby(c echo.Context) error {
	latStr := c.QueryParam("lat")
	lngStr := c.QueryParam("lng")
	radiusStr := c.QueryParam("radius")
	fieldType := c.QueryParam("field_type")
	limitStr := c.QueryParam("limit")

	if latStr == "" || lngStr == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "Latitude and longitude are required",
			},
		})
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "Invalid latitude",
			},
		})
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "Invalid longitude",
			},
		})
	}

	radius := 5.0
	if radiusStr != "" {
		radius, err = strconv.ParseFloat(radiusStr, 64)
		if err != nil {
			radius = 5.0
		}
	}

	limit := 20
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			limit = 20
		}
	}

	var fieldTypePtr *string
	if fieldType != "" {
		fieldTypePtr = &fieldType
	}

	venues, err := h.venueService.FindNearbyVenues(c.Request().Context(), lat, lng, radius, fieldTypePtr, limit)
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
		"data": venues,
	})
}

// ListVenues lists all venues
// @Summary      List venues
// @Description  Get a list of all venues with pagination
// @Tags         venues
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        limit   query     int  false  "Number of results per page (default: 20)"
// @Param        offset  query     int  false  "Number of results to skip (default: 0)"
// @Success      200     {object}  object
// @Failure      401     {object}  response.ErrorResponse
// @Failure      500     {object}  response.ErrorResponse
// @Router       /api/v1/venues [get]
func (h *VenueHandler) ListVenues(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")

	limit := 20
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	offset := 0
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	venues, err := h.venueService.ListVenues(c.Request().Context(), limit, offset)
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
		"data": venues,
	})
}

// UploadVenueImage handles venue image upload
// @Summary      Upload venue image
// @Description  Upload an image for a venue (only owner can upload)
// @Tags         venues
// @Accept       multipart/form-data
// @Produce      json
// @Security     Bearer
// @Param        id    path      string  true   "Venue ID"
// @Param        file  formData  file    true   "Image file (max 5MB)"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  response.ErrorResponse
// @Failure      401   {object}  response.ErrorResponse
// @Failure      403   {object}  response.ErrorResponse
// @Failure      500   {object}  response.ErrorResponse
// @Router       /api/v1/venues/{id}/upload [post]
func (h *VenueHandler) UploadVenueImage(c echo.Context) error {
	// This will be implemented in the upload handler
	return c.JSON(http.StatusNotImplemented, response.ErrorResponse{
		Error: response.ErrorDetail{
			Code:    string(util.ErrorCodeInternalServer),
			Message: "Not implemented",
		},
	})
}

