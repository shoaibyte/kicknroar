package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			StorageKey("id"),

		field.String("email").
			Unique().
			NotEmpty().
			MaxLen(255).
			Match(regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")).
			Comment("User's email address"),

		field.String("password_hash").
			Sensitive().
			NotEmpty().
			MaxLen(255).
			Comment("Bcrypt hashed password"),

		field.String("full_name").
			NotEmpty().
			MaxLen(100).
			Comment("User's full name"),

		field.String("phone").
			Unique().
			NotEmpty().
			MaxLen(20).
			Match(regexp.MustCompile("^\\+?[0-9]{10,15}$")).
			Comment("User's phone number with optional country code"),

		field.String("profile_image_url").
			Optional().
			MaxLen(500).
			Comment("S3 URL for profile image"),

		field.Enum("skill_level").
			Values("beginner", "intermediate", "advanced", "professional").
			Default("intermediate").
			Comment("Player's skill level"),

		field.JSON("preferred_locations", []string{}).
			Optional().
			Comment("Array of preferred location names"),

		field.Bool("is_verified").
			Default(false).
			Comment("Email verification status"),

		field.Bool("is_active").
			Default(true).
			Comment("Soft delete flag"),

		field.Time("last_login_at").
			Optional().
			Nillable().
			Comment("Last login timestamp"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Account creation timestamp"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Last update timestamp"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stats", UserStats.Type).
			Unique(),
		edge.To("created_matches", Match.Type),
		edge.To("participations", MatchParticipant.Type),
		edge.To("notifications", Notification.Type),
		edge.To("owned_venues", Venue.Type),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		// Index on email for fast lookups
		index.Fields("email"),

		// Index on phone for fast lookups
		index.Fields("phone"),

		// Composite index for active users (common query)
		index.Fields("is_active"),
	}
}
