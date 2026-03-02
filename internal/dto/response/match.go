package response

import (
	"time"

	"github.com/google/uuid"
)

// VenueResponse represents a venue in API responses
// @Description Venue information with location and facilities
type VenueResponse struct {
	ID            uuid.UUID              `json:"id"`
	Name          string                 `json:"name"`
	Address       string                 `json:"address"`
	Latitude      float64                `json:"latitude"`
	Longitude     float64                `json:"longitude"`
	GooglePlaceID *string                `json:"google_place_id,omitempty"`
	FieldType     string                 `json:"field_type"`
	SurfaceType   *string                `json:"surface_type,omitempty"`
	Capacity      int                    `json:"capacity"`
	HourlyRate    *int                   `json:"hourly_rate,omitempty"`
	Facilities    []string               `json:"facilities,omitempty"`
	Images        []string               `json:"images,omitempty"`
	ContactInfo   map[string]interface{} `json:"contact_info,omitempty"`
	OperatingHours map[string]interface{} `json:"operating_hours,omitempty"`
	Rating        float64                `json:"rating"`
	TotalRatings  int                    `json:"total_ratings"`
	IsVerified    bool                   `json:"is_verified"`
	CreatedAt     time.Time              `json:"created_at"`
}

// MatchResponse represents a match in API responses
// @Description Match information with venue and creator details
type MatchResponse struct {
	ID                uuid.UUID      `json:"id"`
	Title             string         `json:"title"`
	Description       *string        `json:"description,omitempty"`
	Venue             *VenueResponse `json:"venue"`
	MatchDate         time.Time      `json:"match_date"`
	StartTime         time.Time      `json:"start_time"`
	DurationHours     float64         `json:"duration_hours"`
	MaxPlayers        int            `json:"max_players"`
	CurrentPlayers    int            `json:"current_players"`
	CostPerPlayer     int            `json:"cost_per_player"`
	SkillLevelRequired *string        `json:"skill_level_required,omitempty"`
	MatchType         string         `json:"match_type"`
	Status            string         `json:"status"`
	Visibility        string         `json:"visibility"`
	RulesNotes        *string        `json:"rules_notes,omitempty"`
	Creator           *UserResponse  `json:"creator,omitempty"`
	CreatedAt         time.Time      `json:"created_at"`
}

