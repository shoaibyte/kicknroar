package request

import "time"

// CreateMatchRequest represents a match creation request
// @Description Match creation request with venue, date, time, and player details
type CreateMatchRequest struct {
	Title             string    `json:"title" validate:"required,min=3,max=100" example:"Weekend Futsal Match"`
	Description       *string   `json:"description,omitempty" example:"Casual futsal match for intermediate players"`
	VenueID           string    `json:"venue_id" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
	MatchDate         time.Time `json:"match_date" validate:"required" example:"2024-12-01T00:00:00Z"`
	StartTime         time.Time `json:"start_time" validate:"required" example:"2024-12-01T18:00:00Z"`
	DurationHours     float64   `json:"duration_hours" validate:"required,min=0.5,max=4" example:"2"`
	MaxPlayers        int       `json:"max_players" validate:"required,min=2,max=22" example:"10"`
	CostPerPlayer     int       `json:"cost_per_player" validate:"required,min=1" example:"500"`
	SkillLevelRequired *string   `json:"skill_level_required,omitempty" validate:"omitempty,oneof=beginner intermediate advanced professional" example:"intermediate"`
	MatchType         string    `json:"match_type" validate:"required,oneof=casual competitive tournament" example:"casual"`
	Visibility        string    `json:"visibility" validate:"required,oneof=public private friends_only" example:"public"`
	RulesNotes        *string   `json:"rules_notes,omitempty" example:"Bring your own water bottle"`
}

// UpdateMatchRequest represents a match update request
// @Description Match update request - all fields are optional
type UpdateMatchRequest struct {
	Title             *string   `json:"title,omitempty" validate:"omitempty,min=3,max=100" example:"Weekend Futsal Match - Updated"`
	Description       *string   `json:"description,omitempty" example:"Updated match description"`
	MatchDate         *time.Time `json:"match_date,omitempty" example:"2024-12-01T00:00:00Z"`
	StartTime         *time.Time `json:"start_time,omitempty" example:"2024-12-01T19:00:00Z"`
	DurationHours     *float64  `json:"duration_hours,omitempty" validate:"omitempty,min=0.5,max=4" example:"2.5"`
	MaxPlayers        *int      `json:"max_players,omitempty" validate:"omitempty,min=2,max=22" example:"12"`
	CostPerPlayer     *int      `json:"cost_per_player,omitempty" validate:"omitempty,min=1" example:"600"`
	SkillLevelRequired *string   `json:"skill_level_required,omitempty" validate:"omitempty,oneof=beginner intermediate advanced professional" example:"advanced"`
	MatchType         *string   `json:"match_type,omitempty" validate:"omitempty,oneof=casual competitive tournament" example:"competitive"`
	Visibility        *string   `json:"visibility,omitempty" validate:"omitempty,oneof=public private friends_only" example:"public"`
	RulesNotes        *string   `json:"rules_notes,omitempty" example:"Updated rules and notes"`
}

