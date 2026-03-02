package testutil

import (
	"context"
	"kicknroar/internal/config"
	"kicknroar/internal/database"
	"kicknroar/internal/ent"
	"kicknroar/internal/ent/enttest"
	"kicknroar/internal/pkg/jwt"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3" // SQLite driver for testing
)

// TestConfig returns a test configuration
func TestConfig() *config.Config {
	return &config.Config{
		Server: config.ServerConfig{
			Port:           "8000",
			Env:            "test",
			AllowedOrigins: []string{"http://localhost:3000"},
		},
		Database: config.DatabaseConfig{
			URL:             "postgres://test:test@localhost:5432/test?sslmode=disable",
			MaxOpenConns:    10,
			MaxIdleConns:    5,
			ConnMaxLifetime: 5 * time.Minute,
		},
		JWT: config.JWTConfig{
			Secret:        "test-secret-key-for-testing-only",
			Expiry:        15 * time.Minute,
			RefreshExpiry: 7 * 24 * time.Hour,
		},
		AWS: config.AWSConfig{
			AccessKeyID:     "test-key",
			SecretAccessKey: "test-secret",
			Region:          "ap-south-1",
			S3Bucket:        "test-bucket",
		},
		RateLimit: config.RateLimitConfig{
			RequestsPerMinute:       100,
			AuthRequestsPerMinute:   5,
			UploadRequestsPerMinute: 10,
		},
	}
}

// TestingT is the interface for testing
type TestingT interface {
	Error(...interface{})
	FailNow()
}

// SetupTestDB creates a test database client using SQLite for fast tests
func SetupTestDB(t TestingT) *ent.Client {
	// Use SQLite for testing (faster, no external dependencies)
	// Note: SQLite doesn't support PostGIS, use SetupTestDBWithPostgres for geospatial tests
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	return client
}

// SetupTestDBWithPostgres creates a test database client using PostgreSQL
// Use this for integration tests that need PostGIS
func SetupTestDBWithPostgres(t TestingT, dsn string) *ent.Client {
	client := enttest.Open(t, "postgres", dsn)
	return client
}

// CleanupTestDB closes the test database
func CleanupTestDB(client *ent.Client) {
	if client != nil {
		client.Close()
	}
}

// TestJWTManager returns a JWT manager for testing
func TestJWTManager() *jwt.Manager {
	cfg := TestConfig()
	return jwt.NewManager(&cfg.JWT)
}

// CreateTestUser creates a test user in the database
func CreateTestUser(ctx context.Context, client *ent.Client, email, passwordHash, name, phone string) (*ent.User, error) {
	user, err := client.User.
		Create().
		SetEmail(email).
		SetPasswordHash(passwordHash).
		SetFullName(name).
		SetPhone(phone).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Create user stats
	_, _ = client.UserStats.
		Create().
		SetUserID(user.ID).
		Save(ctx)

	return user, nil
}

// CreateTestVenue creates a test venue in the database
func CreateTestVenue(ctx context.Context, client *ent.Client, name, address string, lat, lng float64) (*ent.Venue, error) {
	return client.Venue.
		Create().
		SetName(name).
		SetAddress(address).
		SetLatitude(lat).
		SetLongitude(lng).
		SetFieldType("futsal").
		SetCapacity(10).
		Save(ctx)
}

// CreateTestMatch creates a test match in the database
func CreateTestMatch(ctx context.Context, client *ent.Client, creatorID, venueID string, matchDate time.Time, startTime time.Time) (*ent.Match, error) {
	var creatorUID, venueUID interface{}
	var err error

	// Try to parse as UUID first
	if creatorID != "" {
		creatorUID, err = database.ParseUUID(creatorID)
		if err != nil {
			// Fallback: get first user
			creatorUID, _ = client.User.Query().FirstID(ctx)
		}
	} else {
		creatorUID, _ = client.User.Query().FirstID(ctx)
	}

	if venueID != "" {
		venueUID, err = database.ParseUUID(venueID)
		if err != nil {
			// Fallback: get first venue
			venueUID, _ = client.Venue.Query().FirstID(ctx)
		}
	} else {
		venueUID, _ = client.Venue.Query().FirstID(ctx)
	}

	return client.Match.
		Create().
		SetTitle("Test Match").
		SetCreatorID(creatorUID.(uuid.UUID)).
		SetVenueID(venueUID.(uuid.UUID)).
		SetMatchDate(matchDate).
		SetStartTime(startTime).
		SetMaxPlayers(10).
		SetCostPerPlayer(150).
		SetMatchType("casual").
		SetStatus("open").
		SetVisibility("public").
		Save(ctx)
}
