package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// UserStats holds the schema definition for the UserStats entity.
type UserStats struct {
	ent.Schema
}

// Fields of the UserStats.
func (UserStats) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).
			Comment("User ID (primary key)"),

		field.Int("total_matches").
			Default(0).
			Comment("Total matches participated"),

		field.Int("matches_attended").
			Default(0).
			Comment("Matches actually attended"),

		field.Int("matches_organized").
			Default(0).
			Comment("Matches organized/created"),

		field.Int("no_shows").
			Default(0).
			Comment("Number of no-shows"),

		field.Float("average_rating").
			Default(0.0).
			Min(0.0).
			Max(5.0).
			Comment("Average rating received"),

		field.Int("total_ratings_received").
			Default(0).
			Comment("Total number of ratings received"),

		field.Float("reliability_score").
			Default(5.0).
			Min(0.0).
			Max(5.0).
			Comment("Reliability score (0-5)"),

		field.Time("last_match_at").
			Optional().
			Nillable().
			Comment("Last match timestamp"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Last update timestamp"),
	}
}

// Edges of the UserStats.
func (UserStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Unique().
			Required().
			Field("user_id").
			Ref("stats"),
	}
}

// Indexes of the UserStats.
func (UserStats) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("average_rating"),
		index.Fields("reliability_score"),
	}
}
