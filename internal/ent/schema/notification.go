package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			StorageKey("id"),

		field.UUID("user_id", uuid.UUID{}).
			Comment("User who receives the notification"),

		field.UUID("match_id", uuid.UUID{}).
			Optional().
			Nillable().
			Comment("Related match ID (if applicable)"),

		field.Enum("type").
			Values(
				"match_reminder",
				"new_match",
				"match_update",
				"player_joined",
				"player_left",
				"match_full",
				"payment_due",
				"match_cancelled",
				"rating_request",
			).
			Comment("Notification type"),

		field.String("title").
			NotEmpty().
			MaxLen(100).
			Comment("Notification title"),

		field.Text("message").
			NotEmpty().
			Comment("Notification message"),

		field.String("action_url").
			Optional().
			Comment("URL for action (if applicable)"),

		field.Bool("is_read").
			Default(false).
			Comment("Read status"),

		field.Bool("is_sent").
			Default(false).
			Comment("Sent status"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Creation timestamp"),

		field.Time("read_at").
			Optional().
			Nillable().
			Comment("When notification was read"),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("match", Match.Type).
			Unique().
			Field("match_id"),
	}
}

// Indexes of the Notification.
func (Notification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("user_id", "is_read"),
	}
}
