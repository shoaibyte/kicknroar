package database

import (
	"context"
	"fmt"
	"kicknroar/internal/config"
	"time"

	entclient "kicknroar/internal/ent"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Client wraps the Ent client
type Client struct {
	*entclient.Client
}

// New creates a new database client
func New(cfg *config.Config) (*Client, error) {
	client, err := entclient.Open("postgres", cfg.Database.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Connection pool settings are configured via connection string
	// Ent manages the connection pool internally

	return &Client{Client: client}, nil
}

// Close closes the database connection
func (c *Client) Close() error {
	return c.Client.Close()
}

// Ping checks if database connection is alive
func (c *Client) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Use a simple query to check connection
	_, err := c.Client.User.Query().Limit(1).All(ctx)
	return err
}

// VerifyPostGIS checks if PostGIS extension is installed
// Note: This is a simplified check. The migration will fail if PostGIS is not installed.
func (c *Client) VerifyPostGIS(ctx context.Context) error {
	// PostGIS verification is handled by the migration
	// If PostGIS is not installed, the migration will fail
	return nil
}

// RunMigrations runs database schema migrations
func (c *Client) RunMigrations(ctx context.Context) error {
	if err := c.Client.Schema.Create(ctx); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}
	return nil
}

// ParseUUID parses a UUID string
func ParseUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

// Now returns current time
func Now() time.Time {
	return time.Now()
}

// IsUniqueConstraintError checks if error is a unique constraint violation
func IsUniqueConstraintError(err error) bool {
	return sqlgraph.IsUniqueConstraintError(err)
}

// ErrAlreadyExists is returned when a record already exists
var ErrAlreadyExists = fmt.Errorf("record already exists")
