package repository

import (
	"context"
	"fmt"

	"kicknroar/internal/database"
	"kicknroar/internal/ent"
	"kicknroar/internal/ent/matchparticipant"
)

// ParticipantRepository handles match participant data access
type ParticipantRepository struct {
	client *ent.Client
}

// NewParticipantRepository creates a new participant repository
func NewParticipantRepository(client *ent.Client) *ParticipantRepository {
	return &ParticipantRepository{client: client}
}

// Join adds a user to a match
func (r *ParticipantRepository) Join(ctx context.Context, matchID, userID string) (*ent.MatchParticipant, error) {
	matchUID, err := database.ParseUUID(matchID)
	if err != nil {
		return nil, err
	}

	userUID, err := database.ParseUUID(userID)
	if err != nil {
		return nil, err
	}

	// Check if already joined
	existing, _ := r.client.MatchParticipant.
		Query().
		Where(
			matchparticipant.MatchID(matchUID),
			matchparticipant.UserID(userUID),
		).
		Only(ctx)

	if existing != nil {
		return nil, fmt.Errorf("user already joined this match")
	}

	return r.client.MatchParticipant.
		Create().
		SetMatchID(matchUID).
		SetUserID(userUID).
		Save(ctx)
}

// Leave removes a user from a match
func (r *ParticipantRepository) Leave(ctx context.Context, matchID, userID string) error {
	matchUID, err := database.ParseUUID(matchID)
	if err != nil {
		return err
	}

	userUID, err := database.ParseUUID(userID)
	if err != nil {
		return err
	}

	_, err = r.client.MatchParticipant.
		Delete().
		Where(
			matchparticipant.MatchID(matchUID),
			matchparticipant.UserID(userUID),
		).
		Exec(ctx)

	return err
}

// GetParticipants gets all participants for a match
func (r *ParticipantRepository) GetParticipants(ctx context.Context, matchID string) ([]*ent.MatchParticipant, error) {
	matchUID, err := database.ParseUUID(matchID)
	if err != nil {
		return nil, err
	}

	return r.client.MatchParticipant.
		Query().
		Where(matchparticipant.MatchID(matchUID)).
		WithUser().
		All(ctx)
}

