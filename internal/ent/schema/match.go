package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Match holds the schema definition for the Match entity.
type Match struct {
	ent.Schema
}

// Fields of the Match.
func (Match) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			StorageKey("id"),

		field.String("title").
			NotEmpty().
			MaxLen(100).
			Comment("Match title"),

		field.Text("description").
			Optional().
			Comment("Match description"),

		field.UUID("creator_id", uuid.UUID{}).
			Comment("User who created the match"),

		field.UUID("venue_id", uuid.UUID{}).
			Comment("Venue where match will be played"),

		field.Time("match_date").
			Comment("Date of the match"),

		field.Time("start_time").
			Comment("Start time of the match"),

		field.Float("duration_hours").
			Default(1.5).
			Comment("Duration in hours"),

		field.Int("max_players").
			Default(10).
			Positive().
			Comment("Maximum number of players"),

		field.Int("current_players").
			Default(1).
			Min(0).
			Comment("Current number of players"),

		field.Int("cost_per_player").
			Positive().
			Comment("Cost per player in BDT"),

		field.Enum("skill_level_required").
			Values("beginner", "intermediate", "advanced", "professional").
			Optional().
			Comment("Required skill level"),

		field.Enum("match_type").
			Values("casual", "competitive", "tournament").
			Default("casual").
			Comment("Type of match"),

		field.Enum("status").
			Values("open", "full", "confirmed", "ongoing", "completed", "cancelled").
			Default("open").
			Comment("Match status"),

		field.Enum("visibility").
			Values("public", "private", "friends_only").
			Default("public").
			Comment("Match visibility"),

		field.Text("rules_notes").
			Optional().
			Comment("Rules and notes"),

		field.Bool("reminder_sent").
			Default(false).
			Comment("Whether reminder has been sent"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Creation timestamp"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Last update timestamp"),
	}
}

// Edges of the Match.
func (Match) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("creator", User.Type).
			Unique().
			Required().
			Field("creator_id"),
		edge.To("venue", Venue.Type).
			Unique().
			Required().
			Field("venue_id"),
		edge.To("participants", MatchParticipant.Type),
		edge.To("notifications", Notification.Type),
	}
}

// Indexes of the Match.
func (Match) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("creator_id"),
		index.Fields("venue_id"),
		index.Fields("match_date"),
		index.Fields("status"),
		index.Fields("match_date", "start_time", "status"),
	}
}
