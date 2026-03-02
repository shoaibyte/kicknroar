package service

import (
	"context"

	"kicknroar/internal/dto/request"
	"kicknroar/internal/dto/response"
	"kicknroar/internal/ent"
	"kicknroar/internal/pkg/jwt"
	"kicknroar/internal/repository"
	"kicknroar/internal/util"
)

// AuthService handles authentication business logic
type AuthService struct {
	userRepo  *repository.UserRepository
	jwtManager *jwt.Manager
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo *repository.UserRepository, jwtManager *jwt.Manager) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

// Signup creates a new user account
func (s *AuthService) Signup(ctx context.Context, req *request.SignupRequest) (*response.AuthResponse, error) {
	// Check if email already exists
	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, util.ErrEmailExists()
	}

	// Check if phone already exists
	existing, _ = s.userRepo.FindByPhone(ctx, req.Phone)
	if existing != nil {
		return nil, util.ErrPhoneExists()
	}

	// Hash password
	passwordHash, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	// Create user
	user, err := s.userRepo.Create(ctx, req.Email, passwordHash, req.FullName, req.Phone)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	// Generate tokens
	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	// Update last login
	_ = s.userRepo.UpdateLastLogin(ctx, user.ID.String())

	return &response.AuthResponse{
		User:         toUserResponse(user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Login authenticates a user
func (s *AuthService) Login(ctx context.Context, req *request.LoginRequest) (*response.AuthResponse, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, util.ErrInvalidCredentials()
	}

	// Check password
	if !util.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, util.ErrInvalidCredentials()
	}

	// Check if user is active
	if !user.IsActive {
		return nil, util.ErrUnauthorized()
	}

	// Generate tokens
	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	// Update last login
	_ = s.userRepo.UpdateLastLogin(ctx, user.ID.String())

	return &response.AuthResponse{
		User:         toUserResponse(user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Refresh generates a new access token from a refresh token
func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (*response.AuthResponse, error) {
	// Validate refresh token
	claims, err := s.jwtManager.ValidateToken(refreshToken)
	if err != nil {
		return nil, util.ErrTokenExpired()
	}

	// Find user
	user, err := s.userRepo.FindByID(ctx, claims.UserID)
	if err != nil {
		return nil, util.ErrUserNotFound()
	}

	// Check if user is active
	if !user.IsActive {
		return nil, util.ErrUnauthorized()
	}

	// Generate new access token
	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	// Generate new refresh token
	newRefreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	return &response.AuthResponse{
		User:         toUserResponse(user),
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

// Logout logs out a user (client-side token removal, but we can add token blacklisting here)
func (s *AuthService) Logout(ctx context.Context, userID string) error {
	// In a more advanced implementation, we could blacklist the token
	// For now, logout is handled client-side by removing the token
	return nil
}

// toUserResponse converts an ent.User to UserResponse
func toUserResponse(user *ent.User) *response.UserResponse {
	var profileImageURL *string
	if user.ProfileImageURL != "" {
		profileImageURL = &user.ProfileImageURL
	}

	return &response.UserResponse{
		ID:             user.ID,
		Email:          user.Email,
		FullName:       user.FullName,
		Phone:          user.Phone,
		ProfileImageURL: profileImageURL,
		SkillLevel:     string(user.SkillLevel),
		IsVerified:     user.IsVerified,
		CreatedAt:      user.CreatedAt,
	}
}

