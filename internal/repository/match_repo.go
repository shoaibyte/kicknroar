package repository

import (
	"context"
	"time"

	"kicknroar/internal/database"
	"kicknroar/internal/ent"
	"kicknroar/internal/ent/match"
)

// MatchRepository handles match data access
type MatchRepository struct {
	client *ent.Client
}

// NewMatchRepository creates a new match repository
func NewMatchRepository(client *ent.Client) *MatchRepository {
	return &MatchRepository{client: client}
}

// Create creates a new match
func (r *MatchRepository) Create(ctx context.Context, data map[string]interface{}) (*ent.Match, error) {
	create := r.client.Match.Create()

	if title, ok := data["title"].(string); ok {
		create = create.SetTitle(title)
	}
	if description, ok := data["description"].(*string); ok && description != nil {
		create = create.SetDescription(*description)
	}
	if creatorID, ok := data["creator_id"].(string); ok {
		uid, err := database.ParseUUID(creatorID)
		if err == nil {
			create = create.SetCreatorID(uid)
		}
	}
	if venueID, ok := data["venue_id"].(string); ok {
		uid, err := database.ParseUUID(venueID)
		if err == nil {
			create = create.SetVenueID(uid)
		}
	}
	if matchDate, ok := data["match_date"].(time.Time); ok {
		create = create.SetMatchDate(matchDate)
	}
	if startTime, ok := data["start_time"].(time.Time); ok {
		create = create.SetStartTime(startTime)
	}
	if durationHours, ok := data["duration_hours"].(float64); ok {
		create = create.SetDurationHours(durationHours)
	}
	if maxPlayers, ok := data["max_players"].(int); ok {
		create = create.SetMaxPlayers(maxPlayers)
	}
	if costPerPlayer, ok := data["cost_per_player"].(int); ok {
		create = create.SetCostPerPlayer(costPerPlayer)
	}
	if matchType, ok := data["match_type"].(string); ok {
		create = create.SetMatchType(match.MatchType(matchType))
	}
	if visibility, ok := data["visibility"].(string); ok {
		create = create.SetVisibility(match.Visibility(visibility))
	}
	if rulesNotes, ok := data["rules_notes"].(*string); ok && rulesNotes != nil {
		create = create.SetRulesNotes(*rulesNotes)
	}

	return create.Save(ctx)
}

// FindByID finds a match by ID
func (r *MatchRepository) FindByID(ctx context.Context, id string) (*ent.Match, error) {
	uid, err := database.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	return r.client.Match.
		Query().
		Where(match.ID(uid)).
		WithCreator().
		WithVenue().
		Only(ctx)
}

// List lists matches with filters
func (r *MatchRepository) List(ctx context.Context, filters map[string]interface{}, limit, offset int) ([]*ent.Match, error) {
	query := r.client.Match.Query().
		WithCreator().
		WithVenue()

	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where(match.StatusEQ(match.Status(status)))
	}
	if dateFrom, ok := filters["date_from"].(time.Time); ok {
		query = query.Where(match.MatchDateGTE(dateFrom))
	}
	if dateTo, ok := filters["date_to"].(time.Time); ok {
		query = query.Where(match.MatchDateLTE(dateTo))
	}

	return query.
		Order(match.ByMatchDate()).
		Limit(limit).
		Offset(offset).
		All(ctx)
}

// Update updates a match
func (r *MatchRepository) Update(ctx context.Context, id string, updates map[string]interface{}) (*ent.Match, error) {
	uid, err := database.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	update := r.client.Match.UpdateOneID(uid)

	if title, ok := updates["title"].(*string); ok && title != nil {
		update = update.SetTitle(*title)
	}
	if description, ok := updates["description"].(*string); ok {
		update = update.SetDescription(*description)
	}
	if matchDate, ok := updates["match_date"].(*time.Time); ok && matchDate != nil {
		update = update.SetMatchDate(*matchDate)
	}
	if startTime, ok := updates["start_time"].(*time.Time); ok && startTime != nil {
		update = update.SetStartTime(*startTime)
	}
	if maxPlayers, ok := updates["max_players"].(*int); ok && maxPlayers != nil {
		update = update.SetMaxPlayers(*maxPlayers)
	}
	if costPerPlayer, ok := updates["cost_per_player"].(*int); ok && costPerPlayer != nil {
		update = update.SetCostPerPlayer(*costPerPlayer)
	}
	if status, ok := updates["status"].(string); ok {
		update = update.SetStatus(match.Status(status))
	}

	return update.Save(ctx)
}

// Delete deletes a match
func (r *MatchRepository) Delete(ctx context.Context, id string) error {
	uid, err := database.ParseUUID(id)
	if err != nil {
		return err
	}

	return r.client.Match.DeleteOneID(uid).Exec(ctx)
}
