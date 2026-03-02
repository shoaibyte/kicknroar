package response

// ErrorResponse represents an error response
// @Description Standard error response format
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error information
// @Description Detailed error information
type ErrorDetail struct {
	Code    string   `json:"code" example:"VALIDATION_ERROR"`
	Message string   `json:"message" example:"Invalid request data"`
	Details []string  `json:"details,omitempty" example:"email is required,password must be at least 8 characters"`
}

