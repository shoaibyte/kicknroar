package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Venue holds the schema definition for the Venue entity.
type Venue struct {
	ent.Schema
}

// Fields of the Venue.
func (Venue) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			StorageKey("id"),

		field.String("name").
			NotEmpty().
			MaxLen(100).
			Comment("Venue name"),

		field.Text("address").
			NotEmpty().
			Comment("Full address of the venue"),

		field.Float("latitude").
			Comment("Latitude coordinate (WGS 84)"),

		field.Float("longitude").
			Comment("Longitude coordinate (WGS 84)"),

		field.String("google_place_id").
			Optional().
			MaxLen(255).
			Comment("Google Places API ID"),

		field.Enum("field_type").
			Values("futsal", "football", "astro").
			Comment("Type of field"),

		field.Enum("surface_type").
			Values("grass", "artificial", "concrete").
			Optional().
			Comment("Surface type"),

		field.Int("capacity").
			Default(10).
			Positive().
			Comment("Maximum number of players"),

		field.Int("hourly_rate").
			Optional().
			Comment("Hourly rate in BDT"),

		field.JSON("facilities", []string{}).
			Optional().
			Comment("Array of facilities (parking, changing_room, lighting)"),

		field.JSON("images", []string{}).
			Optional().
			Comment("Array of S3 image URLs"),

		field.JSON("contact_info", map[string]interface{}{}).
			Optional().
			Comment("Contact information (JSON)"),

		field.JSON("operating_hours", map[string]interface{}{}).
			Optional().
			Comment("Operating hours (JSON)"),

		field.UUID("owner_id", uuid.UUID{}).
			Optional().
			Nillable().
			Comment("Owner user ID"),

		field.Bool("is_verified").
			Default(false).
			Comment("Verification status"),

		field.Bool("is_active").
			Default(true).
			Comment("Active status"),

		field.Float("rating").
			Default(0.0).
			Min(0.0).
			Max(5.0).
			Comment("Average rating"),

		field.Int("total_ratings").
			Default(0).
			Comment("Total number of ratings"),

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

// Edges of the Venue.
func (Venue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).
			Unique().
			Field("owner_id"),
		edge.To("matches", Match.Type),
	}
}

// Indexes of the Venue.
func (Venue) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("field_type"),
		index.Fields("is_active"),
		// Note: PostGIS spatial index will be created via migration
	}
}
