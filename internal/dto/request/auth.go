package request

// SignupRequest represents a signup request
// @Description User registration request
type SignupRequest struct {
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"SecurePass123!"`
	FullName string `json:"full_name" validate:"required,min=2,max=100" example:"John Doe"`
	Phone    string `json:"phone" validate:"required,min=10,max=20" example:"+8801712345678"`
}

// LoginRequest represents a login request
// @Description User login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required" example:"SecurePass123!"`
}

// RefreshRequest represents a refresh token request
// @Description Token refresh request
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

