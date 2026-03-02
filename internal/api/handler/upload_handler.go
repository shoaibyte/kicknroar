package handler

import (
	"fmt"
	"net/http"

	"kicknroar/internal/api/middleware"
	"kicknroar/internal/dto/response"
	"kicknroar/internal/pkg/aws"
	"kicknroar/internal/service"
	"kicknroar/internal/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// UploadHandler handles file upload requests
type UploadHandler struct {
	storageService *service.StorageService
}

// NewUploadHandler creates a new upload handler
func NewUploadHandler(storageService *service.StorageService) *UploadHandler {
	return &UploadHandler{
		storageService: storageService,
	}
}

// UploadAvatar handles avatar upload
// @Summary      Upload user avatar
// @Description  Upload a profile image for the authenticated user (max 5MB, supports: jpg, jpeg, png, gif, webp)
// @Tags         upload
// @Accept       multipart/form-data
// @Produce      json
// @Security     Bearer
// @Param        file  formData  file  true  "Image file (max 5MB)"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  response.ErrorResponse
// @Failure      401   {object}  response.ErrorResponse
// @Failure      500   {object}  response.ErrorResponse
// @Router       /api/v1/upload/avatar [post]
func (h *UploadHandler) UploadAvatar(c echo.Context) error {
	userID := middleware.GetUserID(c)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeUnauthorized),
				Message: "Unauthorized",
			},
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeValidationError),
				Message: "File is required",
			},
		})
	}

	// Validate file type
	if !aws.ValidateImageType(file.Filename) {
		return c.JSON(http.StatusBadRequest, util.ErrInvalidFileType())
	}

	// Validate file size (5MB max)
	if file.Size > 5*1024*1024 {
		return c.JSON(http.StatusBadRequest, util.ErrFileTooLarge())
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeInternalServer),
				Message: "Failed to open file",
			},
		})
	}
	defer src.Close()

	// Upload to S3
	url, err := h.storageService.UploadAvatar(c.Request().Context(), userID, src, file.Filename)
	log.Info(fmt.Sprintf("err: %s", err))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: response.ErrorDetail{
				Code:    string(util.ErrorCodeUploadFailed),
				Message: "Failed to upload file",
			},
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"url": url,
	})
}
