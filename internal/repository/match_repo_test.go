package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"kicknroar/internal/testutil"
)

func TestMatchRepository_Create(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewMatchRepository(client)
	ctx := context.Background()

	// Create test user and venue
	user, _ := testutil.CreateTestUser(ctx, client, "creator@example.com", "hash", "Creator", "+8801712345678")
	venue, _ := testutil.CreateTestVenue(ctx, client, "Test Venue", "123 St", 23.8103, 90.4125)

	matchDate := time.Now().Add(24 * time.Hour)
	startTime := time.Date(2025, 1, 1, 18, 0, 0, 0, time.UTC)

	data := map[string]interface{}{
		"title":          "Test Match",
		"creator_id":     user.ID.String(),
		"venue_id":       venue.ID.String(),
		"match_date":     matchDate,
		"start_time":     startTime,
		"max_players":    10,
		"cost_per_player": 150,
		"match_type":     "casual",
		"visibility":     "public",
	}

	match, err := repo.Create(ctx, data)
	assert.NoError(t, err)
	assert.NotNil(t, match)
	assert.Equal(t, "Test Match", match.Title)
	assert.Equal(t, 10, match.MaxPlayers)
	assert.Equal(t, "open", string(match.Status))
}

func TestMatchRepository_FindByID(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewMatchRepository(client)
	ctx := context.Background()

	// Create test data
	user, _ := testutil.CreateTestUser(ctx, client, "creator@example.com", "hash", "Creator", "+8801712345678")
	venue, _ := testutil.CreateTestVenue(ctx, client, "Test Venue", "123 St", 23.8103, 90.4125)
	matchDate := time.Now().Add(24 * time.Hour)
	startTime := time.Date(2025, 1, 1, 18, 0, 0, 0, time.UTC)

	match, _ := testutil.CreateTestMatch(ctx, client, user.ID.String(), venue.ID.String(), matchDate, startTime)

	// Find it
	found, err := repo.FindByID(ctx, match.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, match.ID, found.ID)
}

func TestMatchRepository_List(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewMatchRepository(client)
	ctx := context.Background()

	// Create test data
	user, _ := testutil.CreateTestUser(ctx, client, "creator@example.com", "hash", "Creator", "+8801712345678")
	venue, _ := testutil.CreateTestVenue(ctx, client, "Test Venue", "123 St", 23.8103, 90.4125)
	matchDate := time.Now().Add(24 * time.Hour)
	startTime := time.Date(2025, 1, 1, 18, 0, 0, 0, time.UTC)

	_, _ = testutil.CreateTestMatch(ctx, client, user.ID.String(), venue.ID.String(), matchDate, startTime)

	// List matches
	filters := map[string]interface{}{
		"status": "open",
	}
	matches, err := repo.List(ctx, filters, 10, 0)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(matches), 1)
}

