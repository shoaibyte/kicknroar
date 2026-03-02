package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// MatchParticipant holds the schema definition for the MatchParticipant entity.
type MatchParticipant struct {
	ent.Schema
}

// Fields of the MatchParticipant.
func (MatchParticipant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			StorageKey("id"),

		field.UUID("match_id", uuid.UUID{}).
			Comment("Match ID"),

		field.UUID("user_id", uuid.UUID{}).
			Comment("User ID"),

		field.Time("joined_at").
			Default(time.Now).
			Comment("When user joined the match"),

		field.Enum("payment_status").
			Values("pending", "paid", "refunded").
			Default("pending").
			Comment("Payment status"),

		field.String("payment_method").
			Optional().
			MaxLen(50).
			Comment("Payment method"),

		field.String("payment_transaction_id").
			Optional().
			MaxLen(100).
			Comment("Payment transaction ID"),

		field.Enum("attendance_status").
			Values("confirmed", "attended", "no_show", "cancelled").
			Default("confirmed").
			Comment("Attendance status"),

		field.Int("player_rating").
			Optional().
			Min(1).
			Max(5).
			Comment("Rating given to the player (1-5)"),

		field.Int("match_rating").
			Optional().
			Min(1).
			Max(5).
			Comment("Rating given to the match (1-5)"),

		field.Text("feedback").
			Optional().
			Comment("Feedback text"),
	}
}

// Edges of the MatchParticipant.
func (MatchParticipant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("match", Match.Type).
			Unique().
			Required().
			Field("match_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}

// Indexes of the MatchParticipant.
func (MatchParticipant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("match_id"),
		index.Fields("user_id"),
		index.Fields("match_id", "user_id").
			Unique(),
	}
}
