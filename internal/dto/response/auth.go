package response

import (
	"time"

	"github.com/google/uuid"
)

// UserResponse represents a user in API responses
// @Description User profile information
type UserResponse struct {
	ID             uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email          string    `json:"email" example:"john.doe@example.com"`
	FullName       string    `json:"full_name" example:"John Doe"`
	Phone          string    `json:"phone" example:"+8801712345678"`
	ProfileImageURL *string   `json:"profile_image_url,omitempty" example:"https://example.com/avatar.jpg"`
	SkillLevel     string    `json:"skill_level" example:"intermediate"`
	IsVerified     bool      `json:"is_verified" example:"true"`
	CreatedAt      time.Time `json:"created_at" example:"2024-11-25T19:00:00Z"`
}

// AuthResponse represents an authentication response
// @Description Authentication response with user info and tokens
type AuthResponse struct {
	User         *UserResponse `json:"user"`
	AccessToken  string        `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string        `json:"refresh_token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

