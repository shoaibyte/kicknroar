package util

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppError(t *testing.T) {
	err := NewAppError(ErrorCodeInvalidCredentials, "Invalid credentials", http.StatusUnauthorized)
	assert.Equal(t, ErrorCodeInvalidCredentials, err.Code)
	assert.Equal(t, "Invalid credentials", err.Message)
	assert.Equal(t, http.StatusUnauthorized, err.Status)
}

func TestAppError_WithDetails(t *testing.T) {
	err := NewAppError(ErrorCodeValidationError, "Validation failed", http.StatusBadRequest).
		WithDetails("Email is required", "Password must be at least 8 characters")
	
	assert.Len(t, err.Details, 2)
	assert.Equal(t, "Email is required", err.Details[0])
}

func TestErrorConstructors(t *testing.T) {
	tests := []struct {
		name     string
		constructor func() *AppError
		expectedCode ErrorCode
		expectedStatus int
	}{
		{"InvalidCredentials", ErrInvalidCredentials, ErrorCodeInvalidCredentials, http.StatusUnauthorized},
		{"TokenExpired", ErrTokenExpired, ErrorCodeTokenExpired, http.StatusUnauthorized},
		{"Unauthorized", ErrUnauthorized, ErrorCodeUnauthorized, http.StatusUnauthorized},
		{"EmailExists", ErrEmailExists, ErrorCodeEmailExists, http.StatusConflict},
		{"MatchNotFound", ErrMatchNotFound, ErrorCodeMatchNotFound, http.StatusNotFound},
		{"MatchFull", ErrMatchFull, ErrorCodeMatchFull, http.StatusBadRequest},
		{"VenueNotFound", ErrVenueNotFound, ErrorCodeVenueNotFound, http.StatusNotFound},
		{"UserNotFound", ErrUserNotFound, ErrorCodeUserNotFound, http.StatusNotFound},
		{"FileTooLarge", ErrFileTooLarge, ErrorCodeFileTooLarge, http.StatusBadRequest},
		{"InvalidFileType", ErrInvalidFileType, ErrorCodeInvalidFileType, http.StatusBadRequest},
		{"InternalServer", ErrInternalServer, ErrorCodeInternalServer, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.constructor()
			assert.Equal(t, tt.expectedCode, err.Code)
			assert.Equal(t, tt.expectedStatus, err.Status)
			assert.NotEmpty(t, err.Message)
		})
	}
}

func TestErrValidationError(t *testing.T) {
	err := ErrValidationError("Custom validation message")
	assert.Equal(t, ErrorCodeValidationError, err.Code)
	assert.Equal(t, "Custom validation message", err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
}

