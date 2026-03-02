package repository

import (
	"context"

	"kicknroar/internal/database"
	"kicknroar/internal/ent"
	"kicknroar/internal/ent/venue"
)

// VenueRepository handles venue data access
type VenueRepository struct {
	client *ent.Client
}

// NewVenueRepository creates a new venue repository
func NewVenueRepository(client *ent.Client) *VenueRepository {
	return &VenueRepository{client: client}
}

// Create creates a new venue
func (r *VenueRepository) Create(ctx context.Context, data map[string]interface{}) (*ent.Venue, error) {
	create := r.client.Venue.Create()

	if name, ok := data["name"].(string); ok {
		create = create.SetName(name)
	}
	if address, ok := data["address"].(string); ok {
		create = create.SetAddress(address)
	}
	if lat, ok := data["latitude"].(float64); ok {
		create = create.SetLatitude(lat)
	}
	if lng, ok := data["longitude"].(float64); ok {
		create = create.SetLongitude(lng)
	}
	if fieldType, ok := data["field_type"].(string); ok {
		create = create.SetFieldType(venue.FieldType(fieldType))
	}
	if capacity, ok := data["capacity"].(int); ok {
		create = create.SetCapacity(capacity)
	}
	if hourlyRate, ok := data["hourly_rate"].(*int); ok && hourlyRate != nil {
		create = create.SetHourlyRate(*hourlyRate)
	}
	if facilities, ok := data["facilities"].([]string); ok {
		create = create.SetFacilities(facilities)
	}
	if contactInfo, ok := data["contact_info"].(map[string]interface{}); ok {
		create = create.SetContactInfo(contactInfo)
	}
	if operatingHours, ok := data["operating_hours"].(map[string]interface{}); ok {
		create = create.SetOperatingHours(operatingHours)
	}
	if ownerID, ok := data["owner_id"].(string); ok && ownerID != "" {
		uid, err := database.ParseUUID(ownerID)
		if err == nil {
			create = create.SetOwnerID(uid)
		}
	}

	return create.Save(ctx)
}

// FindByID finds a venue by ID
func (r *VenueRepository) FindByID(ctx context.Context, id string) (*ent.Venue, error) {
	uid, err := database.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	return r.client.Venue.
		Query().
		Where(venue.ID(uid)).
		Only(ctx)
}

// FindNearby finds nearby venues using PostGIS
// Note: This implementation uses raw SQL. For production, consider using Ent's SQL builder
func (r *VenueRepository) FindNearby(ctx context.Context, lat, lng, radiusKm float64, fieldType *string, limit int) ([]*ent.Venue, error) {
	// For now, return all venues - PostGIS query will be implemented when database connection is available
	// This is a placeholder implementation
	return r.client.Venue.
		Query().
		Where(venue.IsActive(true)).
		Limit(limit).
		All(ctx)

	// Simplified implementation - returns all active venues
	// Full PostGIS implementation will be added when raw SQL access is available
	query := r.client.Venue.Query().Where(venue.IsActive(true))
	
	if fieldType != nil {
		query = query.Where(venue.FieldTypeEQ(venue.FieldType(*fieldType)))
	}
	
	return query.Limit(limit).All(ctx)
}

// List lists all active venues
func (r *VenueRepository) List(ctx context.Context, limit, offset int) ([]*ent.Venue, error) {
	return r.client.Venue.
		Query().
		Where(venue.IsActive(true)).
		Limit(limit).
		Offset(offset).
		All(ctx)
}

// Update updates a venue
func (r *VenueRepository) Update(ctx context.Context, id string, updates map[string]interface{}) (*ent.Venue, error) {
	uid, err := database.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	update := r.client.Venue.UpdateOneID(uid)

	if name, ok := updates["name"].(*string); ok && name != nil {
		update = update.SetName(*name)
	}
	if address, ok := updates["address"].(*string); ok && address != nil {
		update = update.SetAddress(*address)
	}
	if lat, ok := updates["latitude"].(*float64); ok && lat != nil {
		update = update.SetLatitude(*lat)
	}
	if lng, ok := updates["longitude"].(*float64); ok && lng != nil {
		update = update.SetLongitude(*lng)
	}
	if fieldType, ok := updates["field_type"].(*string); ok && fieldType != nil {
		update = update.SetFieldType(venue.FieldType(*fieldType))
	}
	if capacity, ok := updates["capacity"].(*int); ok && capacity != nil {
		update = update.SetCapacity(*capacity)
	}
	if hourlyRate, ok := updates["hourly_rate"].(*int); ok {
		update = update.SetHourlyRate(*hourlyRate)
	}
	if facilities, ok := updates["facilities"].([]string); ok {
		update = update.SetFacilities(facilities)
	}
	if contactInfo, ok := updates["contact_info"].(map[string]interface{}); ok {
		update = update.SetContactInfo(contactInfo)
	}
	if operatingHours, ok := updates["operating_hours"].(map[string]interface{}); ok {
		update = update.SetOperatingHours(operatingHours)
	}

	return update.Save(ctx)
}

