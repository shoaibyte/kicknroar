package aws

import (
	"context"
	"fmt"
	"io"
	"mime"
	"path/filepath"
	"strings"
	"time"

	"kicknroar/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

// Client wraps AWS S3 client
type Client struct {
	s3Client *s3.Client
	uploader *manager.Uploader
	bucket   string
}

// NewClient creates a new S3 client
func NewClient(cfg *config.AWSConfig) (*Client, error) {
	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(cfg.Region),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKeyID,
			cfg.SecretAccessKey,
			"",
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	s3Client := s3.NewFromConfig(awsCfg)
	uploader := manager.NewUploader(s3Client)

	return &Client{
		s3Client: s3Client,
		uploader: uploader,
		bucket:   cfg.S3Bucket,
	}, nil
}

// UploadFile uploads a file to S3
func (c *Client) UploadFile(ctx context.Context, key string, body io.Reader, contentType string) (string, error) {
	_, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(c.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Return the S3 URL
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", c.bucket, "ap-southeast-1", key)
	return url, nil
}

// DeleteFile deletes a file from S3
func (c *Client) DeleteFile(ctx context.Context, key string) error {
	_, err := c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	return err
}

// GenerateAvatarKey generates a unique key for avatar uploads
func GenerateAvatarKey(userID string) string {
	filename := fmt.Sprintf("%s-%s.jpg", userID, uuid.New().String()[:8])
	return fmt.Sprintf("avatars/%s/%s", time.Now().Format("2006/01"), filename)
}

// GenerateVenueImageKey generates a unique key for venue image uploads
func GenerateVenueImageKey(venueID string) string {
	filename := fmt.Sprintf("%s-%s.jpg", venueID, uuid.New().String()[:8])
	return fmt.Sprintf("venues/%s/%s", time.Now().Format("2006/01"), filename)
}

// ValidateImageType validates if the file is an image
func ValidateImageType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validExts := []string{".jpg", ".jpeg", ".png", ".webp"}
	for _, validExt := range validExts {
		if ext == validExt {
			return true
		}
	}
	return false
}

// GetContentType returns the content type for a file
func GetContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return "application/octet-stream"
	}
	return mimeType
}
