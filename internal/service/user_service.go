package service

import (
	"context"

	"kicknroar/internal/dto/response"
	"kicknroar/internal/ent"
	"kicknroar/internal/repository"
	"kicknroar/internal/util"
)

// UserService handles user business logic
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetProfile gets user profile
func (s *UserService) GetProfile(ctx context.Context, userID string) (*response.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, util.ErrUserNotFound()
	}

	return toUserResponse(user), nil
}

// UpdateProfile updates user profile
func (s *UserService) UpdateProfile(ctx context.Context, userID string, updates map[string]interface{}) (*response.UserResponse, error) {
	user, err := s.userRepo.UpdateProfile(ctx, userID, updates)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	return toUserResponse(user), nil
}

// GetUserStats gets user statistics
func (s *UserService) GetUserStats(ctx context.Context, userID string) (*ent.UserStats, error) {
	stats, err := s.userRepo.GetStats(ctx, userID)
	if err != nil {
		return nil, util.ErrUserNotFound()
	}

	return stats, nil
}

