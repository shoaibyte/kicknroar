package service

import (
	"context"

	"kicknroar/internal/dto/request"
	"kicknroar/internal/dto/response"
	"kicknroar/internal/ent"
	"kicknroar/internal/repository"
	"kicknroar/internal/util"
)

// MatchService handles match business logic
type MatchService struct {
	matchRepo      *repository.MatchRepository
	participantRepo *repository.ParticipantRepository
	venueRepo      *repository.VenueRepository
}

// NewMatchService creates a new match service
func NewMatchService(matchRepo *repository.MatchRepository, participantRepo *repository.ParticipantRepository, venueRepo *repository.VenueRepository) *MatchService {
	return &MatchService{
		matchRepo:      matchRepo,
		participantRepo: participantRepo,
		venueRepo:      venueRepo,
	}
}

// CreateMatch creates a new match
func (s *MatchService) CreateMatch(ctx context.Context, req *request.CreateMatchRequest, creatorID string) (*response.MatchResponse, error) {
	// Verify venue exists
	_, err := s.venueRepo.FindByID(ctx, req.VenueID)
	if err != nil {
		return nil, util.ErrVenueNotFound()
	}

	data := map[string]interface{}{
		"title":          req.Title,
		"description":    req.Description,
		"creator_id":     creatorID,
		"venue_id":       req.VenueID,
		"match_date":     req.MatchDate,
		"start_time":     req.StartTime,
		"duration_hours": req.DurationHours,
		"max_players":    req.MaxPlayers,
		"cost_per_player": req.CostPerPlayer,
		"match_type":     req.MatchType,
		"visibility":     req.Visibility,
		"rules_notes":    req.RulesNotes,
	}

	match, err := s.matchRepo.Create(ctx, data)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	// Auto-join creator
	_, _ = s.participantRepo.Join(ctx, match.ID.String(), creatorID)

	return toMatchResponse(match), nil
}

// GetMatch gets a match by ID
func (s *MatchService) GetMatch(ctx context.Context, id string) (*response.MatchResponse, error) {
	match, err := s.matchRepo.FindByID(ctx, id)
	if err != nil {
		return nil, util.ErrMatchNotFound()
	}

	return toMatchResponse(match), nil
}

// ListMatches lists matches with filters
func (s *MatchService) ListMatches(ctx context.Context, filters map[string]interface{}, limit, offset int) ([]*response.MatchResponse, error) {
	matches, err := s.matchRepo.List(ctx, filters, limit, offset)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	responses := make([]*response.MatchResponse, len(matches))
	for i, match := range matches {
		responses[i] = toMatchResponse(match)
	}

	return responses, nil
}

// UpdateMatch updates a match
func (s *MatchService) UpdateMatch(ctx context.Context, id string, req *request.UpdateMatchRequest) (*response.MatchResponse, error) {
	updates := make(map[string]interface{})

	if req.Title != nil {
		updates["title"] = req.Title
	}
	if req.Description != nil {
		updates["description"] = req.Description
	}
	if req.MatchDate != nil {
		updates["match_date"] = req.MatchDate
	}
	if req.StartTime != nil {
		updates["start_time"] = req.StartTime
	}
	if req.MaxPlayers != nil {
		updates["max_players"] = req.MaxPlayers
	}
	if req.CostPerPlayer != nil {
		updates["cost_per_player"] = req.CostPerPlayer
	}

	match, err := s.matchRepo.Update(ctx, id, updates)
	if err != nil {
		return nil, util.ErrMatchNotFound()
	}

	return toMatchResponse(match), nil
}

// DeleteMatch deletes a match
func (s *MatchService) DeleteMatch(ctx context.Context, id string) error {
	return s.matchRepo.Delete(ctx, id)
}

// JoinMatch adds a user to a match
func (s *MatchService) JoinMatch(ctx context.Context, matchID, userID string) error {
	// Check if match exists and is open
	match, err := s.matchRepo.FindByID(ctx, matchID)
	if err != nil {
		return util.ErrMatchNotFound()
	}

	if match.Status != "open" {
		return util.ErrMatchFull()
	}

	if match.CurrentPlayers >= match.MaxPlayers {
		return util.ErrMatchFull()
	}

	_, err = s.participantRepo.Join(ctx, matchID, userID)
	if err != nil {
		if err.Error() == "user already joined this match" {
			return util.ErrAlreadyJoined()
		}
		return util.ErrInternalServer()
	}

	// Update match status if full
	if match.CurrentPlayers+1 >= match.MaxPlayers {
		_, _ = s.matchRepo.Update(ctx, matchID, map[string]interface{}{
			"status": "full",
		})
	}

	return nil
}

// LeaveMatch removes a user from a match
func (s *MatchService) LeaveMatch(ctx context.Context, matchID, userID string) error {
	err := s.participantRepo.Leave(ctx, matchID, userID)
	if err != nil {
		return util.ErrInternalServer()
	}

	// Update match status back to open if needed
	match, _ := s.matchRepo.FindByID(ctx, matchID)
	if match != nil && match.Status == "full" && match.CurrentPlayers-1 < match.MaxPlayers {
		_, _ = s.matchRepo.Update(ctx, matchID, map[string]interface{}{
			"status": "open",
		})
	}

	return nil
}

// GetParticipants gets all participants for a match
func (s *MatchService) GetParticipants(ctx context.Context, matchID string) ([]*ent.MatchParticipant, error) {
	return s.participantRepo.GetParticipants(ctx, matchID)
}

// toMatchResponse converts an ent.Match to MatchResponse
func toMatchResponse(match *ent.Match) *response.MatchResponse {
	var venue *response.VenueResponse
	if match.Edges.Venue != nil {
		venue = toVenueResponse(match.Edges.Venue)
	}

	var creator *response.UserResponse
	if match.Edges.Creator != nil {
		creator = toUserResponse(match.Edges.Creator)
	}

	var description *string
	if match.Description != "" {
		description = &match.Description
	}

	var rulesNotes *string
	if match.RulesNotes != "" {
		rulesNotes = &match.RulesNotes
	}

	var skillLevel *string
	if match.SkillLevelRequired != "" {
		s := string(match.SkillLevelRequired)
		skillLevel = &s
	}

	return &response.MatchResponse{
		ID:                match.ID,
		Title:             match.Title,
		Description:       description,
		Venue:             venue,
		MatchDate:         match.MatchDate,
		StartTime:         match.StartTime,
		DurationHours:     match.DurationHours,
		MaxPlayers:        match.MaxPlayers,
		CurrentPlayers:    match.CurrentPlayers,
		CostPerPlayer:     match.CostPerPlayer,
		SkillLevelRequired: skillLevel,
		MatchType:         string(match.MatchType),
		Status:            string(match.Status),
		Visibility:        string(match.Visibility),
		RulesNotes:        rulesNotes,
		Creator:           creator,
		CreatedAt:         match.CreatedAt,
	}
}

