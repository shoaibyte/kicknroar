package service

import (
	"context"
	"fmt"
	"io"

	"kicknroar/internal/pkg/aws"
)

// StorageService handles file storage operations
type StorageService struct {
	s3Client *aws.Client
	bucket   string
}

// NewStorageService creates a new storage service
func NewStorageService(s3Client *aws.Client, bucket string) *StorageService {
	return &StorageService{
		s3Client: s3Client,
		bucket:   bucket,
	}
}

// UploadAvatar uploads a user avatar
func (s *StorageService) UploadAvatar(
	ctx context.Context, userID string, file io.Reader, filename string) (string, error) {
	if s.s3Client == nil {
		return "", fmt.Errorf("S3 client not configured")
	}

	key := aws.GenerateAvatarKey(userID)
	contentType := aws.GetContentType(filename)

	url, err := s.s3Client.UploadFile(ctx, key, file, contentType)
	if err != nil {
		return "", err
	}

	return url, nil
}

// UploadVenueImage uploads a venue image
func (s *StorageService) UploadVenueImage(
	ctx context.Context, venueID string, file io.Reader, filename string) (string, error) {
	if s.s3Client == nil {
		return "", fmt.Errorf("S3 client not configured")
	}

	key := aws.GenerateVenueImageKey(venueID)
	contentType := aws.GetContentType(filename)

	url, err := s.s3Client.UploadFile(ctx, key, file, contentType)
	if err != nil {
		return "", err
	}

	return url, nil
}
