// @title           Kick&Roar API
// @version         1.0
// @description     Football match coordination platform for Dhaka, Bangladesh
// @contact.name    API Support
// @contact.email   support@kickandroar.com
// @license.name    Proprietary
// @host            localhost:8000
// @BasePath        /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kicknroar/internal/api"
	"kicknroar/internal/api/handler"
	"kicknroar/internal/config"
	"kicknroar/internal/database"
	"kicknroar/internal/pkg/aws"
	"kicknroar/internal/pkg/jwt"
	"kicknroar/internal/repository"
	"kicknroar/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("AWS_REGION: %s", cfg.AWS.Region)
	log.Printf("AWS_ACCESS_KEY_ID: %s", cfg.AWS.AccessKeyID)
	if len(cfg.AWS.SecretAccessKey) >= 10 {
		log.Printf("AWS_SECRET_ACCESS_KEY: %s...", cfg.AWS.SecretAccessKey[:10])
	} else {
		log.Printf("AWS_SECRET_ACCESS_KEY: [set]")
	}
	log.Printf("S3_BUCKET_NAME: %s", cfg.AWS.S3Bucket)

	// Initialize database
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Verify PostGIS extension
	ctx := context.Background()
	if err := db.VerifyPostGIS(ctx); err != nil {
		log.Fatalf("PostGIS extension not available: %v", err)
	}

	// Run migrations
	if err := db.RunMigrations(ctx); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize JWT manager
	jwtManager := jwt.NewManager(&cfg.JWT)

	// Initialize AWS S3 client (optional, only if AWS credentials are provided)
	var s3Client *aws.Client
	if cfg.AWS.AccessKeyID != "" && cfg.AWS.SecretAccessKey != "" {
		s3Client, err = aws.NewClient(&cfg.AWS)
		if err != nil {
			log.Printf("Warning: Failed to initialize S3 client: %v", err)
		}
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db.Client)
	matchRepo := repository.NewMatchRepository(db.Client)
	participantRepo := repository.NewParticipantRepository(db.Client)
	venueRepo := repository.NewVenueRepository(db.Client)

	// Initialize services
	authService := service.NewAuthService(userRepo, jwtManager)
	userService := service.NewUserService(userRepo)
	matchService := service.NewMatchService(matchRepo, participantRepo, venueRepo)
	venueService := service.NewVenueService(venueRepo)
	storageService := service.NewStorageService(s3Client, cfg.AWS.S3Bucket)
	_ = service.NewNotificationService(userRepo) // Reserved for future use

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	matchHandler := handler.NewMatchHandler(matchService)
	venueHandler := handler.NewVenueHandler(venueService)
	uploadHandler := handler.NewUploadHandler(storageService)
	docsHandler := handler.NewDocsHandler()

	// Setup router
	e := api.SetupRouter(
		cfg,
		jwtManager,
		authHandler,
		userHandler,
		matchHandler,
		venueHandler,
		uploadHandler,
		docsHandler,
	)

	// Start server
	port := cfg.Server.Port
	if port == "" {
		port = "8000"
	}

	// Start server in a goroutine
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
