package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"kicknroar/internal/testutil"
)

func TestVenueRepository_Create(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewVenueRepository(client)
	ctx := context.Background()

	data := map[string]interface{}{
		"name":       "Test Venue",
		"address":    "123 Test St",
		"latitude":   23.8103,
		"longitude":  90.4125,
		"field_type": "futsal",
		"capacity":   10,
	}

	venue, err := repo.Create(ctx, data)
	assert.NoError(t, err)
	assert.NotNil(t, venue)
	assert.Equal(t, "Test Venue", venue.Name)
	assert.Equal(t, 23.8103, venue.Latitude)
	assert.Equal(t, 90.4125, venue.Longitude)
}

func TestVenueRepository_FindByID(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewVenueRepository(client)
	ctx := context.Background()

	// Create a venue
	venue, _ := testutil.CreateTestVenue(ctx, client, "Test Venue", "123 Test St", 23.8103, 90.4125)

	// Find it
	found, err := repo.FindByID(ctx, venue.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, venue.ID, found.ID)
	assert.Equal(t, "Test Venue", found.Name)
}

func TestVenueRepository_FindByID_NotFound(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewVenueRepository(client)
	ctx := context.Background()

	// Try to find non-existent venue
	_, err := repo.FindByID(ctx, "00000000-0000-0000-0000-000000000000")
	assert.Error(t, err)
}

func TestVenueRepository_List(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewVenueRepository(client)
	ctx := context.Background()

	// Create multiple venues
	_, _ = testutil.CreateTestVenue(ctx, client, "Venue 1", "Address 1", 23.8103, 90.4125)
	_, _ = testutil.CreateTestVenue(ctx, client, "Venue 2", "Address 2", 23.8203, 90.4225)

	venues, err := repo.List(ctx, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, venues, 2)
}

func TestVenueRepository_Update(t *testing.T) {
	client := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(client)

	repo := NewVenueRepository(client)
	ctx := context.Background()

	// Create a venue
	venue, _ := testutil.CreateTestVenue(ctx, client, "Test Venue", "123 Test St", 23.8103, 90.4125)

	// Update it
	newName := "Updated Venue"
	updates := map[string]interface{}{
		"name": &newName,
	}

	updated, err := repo.Update(ctx, venue.ID.String(), updates)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Venue", updated.Name)
}

