package util

import (
	"fmt"
	"net/http"
)

// ErrorCode represents application error codes
type ErrorCode string

const (
	// Auth errors
	ErrorCodeInvalidCredentials ErrorCode = "AUTH_001"
	ErrorCodeTokenExpired       ErrorCode = "AUTH_002"
	ErrorCodeUnauthorized        ErrorCode = "AUTH_003"
	ErrorCodeEmailExists         ErrorCode = "AUTH_004"
	ErrorCodePhoneExists         ErrorCode = "AUTH_005"

	// Match errors
	ErrorCodeMatchNotFound    ErrorCode = "MATCH_001"
	ErrorCodeMatchFull        ErrorCode = "MATCH_002"
	ErrorCodeAlreadyJoined    ErrorCode = "MATCH_003"
	ErrorCodeNotParticipant   ErrorCode = "MATCH_004"
	ErrorCodeMatchPast         ErrorCode = "MATCH_005"
	ErrorCodeMatchCancelled    ErrorCode = "MATCH_006"

	// Venue errors
	ErrorCodeVenueNotFound ErrorCode = "VENUE_001"
	ErrorCodeVenueInactive ErrorCode = "VENUE_002"

	// User errors
	ErrorCodeUserNotFound ErrorCode = "USER_001"
	ErrorCodeUserInactive ErrorCode = "USER_002"

	// Upload errors
	ErrorCodeFileTooLarge  ErrorCode = "UPLOAD_001"
	ErrorCodeInvalidFileType ErrorCode = "UPLOAD_002"
	ErrorCodeUploadFailed  ErrorCode = "UPLOAD_003"

	// Validation errors
	ErrorCodeValidationError ErrorCode = "VALIDATION_001"

	// Server errors
	ErrorCodeInternalServer ErrorCode = "SERVER_001"
	ErrorCodeDatabaseError  ErrorCode = "SERVER_002"
)

// AppError represents an application error
type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Details []string  `json:"details,omitempty"`
	Status  int       `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if len(e.Details) > 0 {
		return fmt.Sprintf("%s: %s - %v", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewAppError creates a new application error
func NewAppError(code ErrorCode, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

// WithDetails adds details to an error
func (e *AppError) WithDetails(details ...string) *AppError {
	e.Details = details
	return e
}

// Common error constructors
func ErrInvalidCredentials() *AppError {
	return NewAppError(ErrorCodeInvalidCredentials, "Invalid email or password", http.StatusUnauthorized)
}

func ErrTokenExpired() *AppError {
	return NewAppError(ErrorCodeTokenExpired, "Token has expired", http.StatusUnauthorized)
}

func ErrUnauthorized() *AppError {
	return NewAppError(ErrorCodeUnauthorized, "Unauthorized access", http.StatusUnauthorized)
}

func ErrEmailExists() *AppError {
	return NewAppError(ErrorCodeEmailExists, "Email already exists", http.StatusConflict)
}

func ErrPhoneExists() *AppError {
	return NewAppError(ErrorCodePhoneExists, "Phone number already exists", http.StatusConflict)
}

func ErrMatchNotFound() *AppError {
	return NewAppError(ErrorCodeMatchNotFound, "Match not found", http.StatusNotFound)
}

func ErrMatchFull() *AppError {
	return NewAppError(ErrorCodeMatchFull, "Match is full", http.StatusBadRequest)
}

func ErrAlreadyJoined() *AppError {
	return NewAppError(ErrorCodeAlreadyJoined, "Already joined this match", http.StatusConflict)
}

func ErrNotParticipant() *AppError {
	return NewAppError(ErrorCodeNotParticipant, "Not a participant of this match", http.StatusForbidden)
}

func ErrVenueNotFound() *AppError {
	return NewAppError(ErrorCodeVenueNotFound, "Venue not found", http.StatusNotFound)
}

func ErrUserNotFound() *AppError {
	return NewAppError(ErrorCodeUserNotFound, "User not found", http.StatusNotFound)
}

func ErrFileTooLarge() *AppError {
	return NewAppError(ErrorCodeFileTooLarge, "File too large", http.StatusBadRequest)
}

func ErrInvalidFileType() *AppError {
	return NewAppError(ErrorCodeInvalidFileType, "Invalid file type", http.StatusBadRequest)
}

func ErrValidationError(message string) *AppError {
	return NewAppError(ErrorCodeValidationError, message, http.StatusBadRequest)
}

func ErrInternalServer() *AppError {
	return NewAppError(ErrorCodeInternalServer, "Internal server error", http.StatusInternalServerError)
}

