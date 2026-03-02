package service

import (
	"context"

	"kicknroar/internal/ent"
	"kicknroar/internal/repository"
)

// NotificationService handles notification business logic
type NotificationService struct {
	userRepo *repository.UserRepository
}

// NewNotificationService creates a new notification service
func NewNotificationService(userRepo *repository.UserRepository) *NotificationService {
	return &NotificationService{
		userRepo: userRepo,
	}
}

// GetNotifications gets user notifications
func (s *NotificationService) GetNotifications(ctx context.Context, userID string, limit, offset int) ([]*ent.Notification, error) {
	// This would need a notification repository
	// For now, return empty
	return []*ent.Notification{}, nil
}

// MarkAsRead marks a notification as read
func (s *NotificationService) MarkAsRead(ctx context.Context, notificationID, userID string) error {
	// Implementation needed
	return nil
}

// MarkAllAsRead marks all notifications as read
func (s *NotificationService) MarkAllAsRead(ctx context.Context, userID string) error {
	// Implementation needed
	return nil
}

