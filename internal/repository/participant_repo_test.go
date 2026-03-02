package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"kicknroar/internal/testutil"
)

func TestParticipantRepository_Join(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewParticipantRepository(client)
	ctx := context.Background()

	// Create test data
	user, _ := testutil.CreateTestUser(ctx, client, "user@example.com", "hash", "User", "+8801712345678")
	venue, _ := testutil.CreateTestVenue(ctx, client, "Test Venue", "123 St", 23.8103, 90.4125)
	matchDate := time.Now().Add(24 * time.Hour)
	startTime := time.Date(2025, 1, 1, 18, 0, 0, 0, time.UTC)
	match, _ := testutil.CreateTestMatch(ctx, client, user.ID.String(), venue.ID.String(), matchDate, startTime)

	// Join match
	participant, err := repo.Join(ctx, match.ID.String(), user.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, participant)
	assert.Equal(t, match.ID, participant.MatchID)
	assert.Equal(t, user.ID, participant.UserID)
}

func TestParticipantRepository_Join_Duplicate(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewParticipantRepository(client)
	ctx := context.Background()

	// Create test data
	user, _ := testutil.CreateTestUser(ctx, client, "user@example.com", "hash", "User", "+8801712345678")
	venue, _ := testutil.CreateTestVenue(ctx, client, "Test Venue", "123 St", 23.8103, 90.4125)
	matchDate := time.Now().Add(24 * time.Hour)
	startTime := time.Date(2025, 1, 1, 18, 0, 0, 0, time.UTC)
	match, _ := testutil.CreateTestMatch(ctx, client, user.ID.String(), venue.ID.String(), matchDate, startTime)

	// Join first time
	_, err := repo.Join(ctx, match.ID.String(), user.ID.String())
	assert.NoError(t, err)

	// Try to join again (should fail)
	_, err = repo.Join(ctx, match.ID.String(), user.ID.String())
	assert.Error(t, err)
}

func TestParticipantRepository_Leave(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewParticipantRepository(client)
	ctx := context.Background()

	// Create test data
	user, _ := testutil.CreateTestUser(ctx, client, "user@example.com", "hash", "User", "+8801712345678")
	venue, _ := testutil.CreateTestVenue(ctx, client, "Test Venue", "123 St", 23.8103, 90.4125)
	matchDate := time.Now().Add(24 * time.Hour)
	startTime := time.Date(2025, 1, 1, 18, 0, 0, 0, time.UTC)
	match, _ := testutil.CreateTestMatch(ctx, client, user.ID.String(), venue.ID.String(), matchDate, startTime)

	// Join match
	_, _ = repo.Join(ctx, match.ID.String(), user.ID.String())

	// Leave match
	err := repo.Leave(ctx, match.ID.String(), user.ID.String())
	assert.NoError(t, err)

	// Verify participant is removed
	participants, _ := repo.GetParticipants(ctx, match.ID.String())
	assert.Len(t, participants, 0)
}

func TestParticipantRepository_GetParticipants(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewParticipantRepository(client)
	ctx := context.Background()

	// Create test data
	user1, _ := testutil.CreateTestUser(ctx, client, "user1@example.com", "hash", "User 1", "+8801711111111")
	user2, _ := testutil.CreateTestUser(ctx, client, "user2@example.com", "hash", "User 2", "+8801722222222")
	venue, _ := testutil.CreateTestVenue(ctx, client, "Test Venue", "123 St", 23.8103, 90.4125)
	matchDate := time.Now().Add(24 * time.Hour)
	startTime := time.Date(2025, 1, 1, 18, 0, 0, 0, time.UTC)
	match, _ := testutil.CreateTestMatch(ctx, client, user1.ID.String(), venue.ID.String(), matchDate, startTime)

	// Join multiple users
	_, _ = repo.Join(ctx, match.ID.String(), user1.ID.String())
	_, _ = repo.Join(ctx, match.ID.String(), user2.ID.String())

	// Get participants
	participants, err := repo.GetParticipants(ctx, match.ID.String())
	assert.NoError(t, err)
	assert.Len(t, participants, 2)
}

