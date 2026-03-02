package service

import (
	"context"

	"kicknroar/internal/dto/request"
	"kicknroar/internal/dto/response"
	"kicknroar/internal/ent"
	"kicknroar/internal/repository"
	"kicknroar/internal/util"
)

// VenueService handles venue business logic
type VenueService struct {
	venueRepo *repository.VenueRepository
}

// NewVenueService creates a new venue service
func NewVenueService(venueRepo *repository.VenueRepository) *VenueService {
	return &VenueService{
		venueRepo: venueRepo,
	}
}

// CreateVenue creates a new venue
func (s *VenueService) CreateVenue(ctx context.Context, req *request.CreateVenueRequest, ownerID string) (*response.VenueResponse, error) {
	data := map[string]interface{}{
		"name":           req.Name,
		"address":        req.Address,
		"latitude":       req.Latitude,
		"longitude":      req.Longitude,
		"field_type":     req.FieldType,
		"capacity":       req.Capacity,
		"facilities":     req.Facilities,
		"contact_info":   req.ContactInfo,
		"operating_hours": req.OperatingHours,
		"owner_id":       ownerID,
	}

	if req.GooglePlaceID != nil {
		data["google_place_id"] = *req.GooglePlaceID
	}
	if req.SurfaceType != nil {
		data["surface_type"] = *req.SurfaceType
	}
	if req.HourlyRate != nil {
		data["hourly_rate"] = *req.HourlyRate
	}

	venue, err := s.venueRepo.Create(ctx, data)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	return toVenueResponse(venue), nil
}

// GetVenue gets a venue by ID
func (s *VenueService) GetVenue(ctx context.Context, id string) (*response.VenueResponse, error) {
	venue, err := s.venueRepo.FindByID(ctx, id)
	if err != nil {
		return nil, util.ErrVenueNotFound()
	}

	return toVenueResponse(venue), nil
}

// UpdateVenue updates a venue
func (s *VenueService) UpdateVenue(ctx context.Context, id string, req *request.UpdateVenueRequest) (*response.VenueResponse, error) {
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = req.Name
	}
	if req.Address != nil {
		updates["address"] = req.Address
	}
	if req.Latitude != nil {
		updates["latitude"] = req.Latitude
	}
	if req.Longitude != nil {
		updates["longitude"] = req.Longitude
	}
	if req.FieldType != nil {
		updates["field_type"] = req.FieldType
	}
	if req.Capacity != nil {
		updates["capacity"] = req.Capacity
	}
	if req.HourlyRate != nil {
		updates["hourly_rate"] = req.HourlyRate
	}
	if req.Facilities != nil {
		updates["facilities"] = req.Facilities
	}
	if req.ContactInfo != nil {
		updates["contact_info"] = req.ContactInfo
	}
	if req.OperatingHours != nil {
		updates["operating_hours"] = req.OperatingHours
	}

	venue, err := s.venueRepo.Update(ctx, id, updates)
	if err != nil {
		return nil, util.ErrVenueNotFound()
	}

	return toVenueResponse(venue), nil
}

// FindNearbyVenues finds nearby venues
func (s *VenueService) FindNearbyVenues(ctx context.Context, lat, lng, radiusKm float64, fieldType *string, limit int) ([]*response.VenueResponse, error) {
	venues, err := s.venueRepo.FindNearby(ctx, lat, lng, radiusKm, fieldType, limit)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	responses := make([]*response.VenueResponse, len(venues))
	for i, venue := range venues {
		responses[i] = toVenueResponse(venue)
	}

	return responses, nil
}

// ListVenues lists all venues
func (s *VenueService) ListVenues(ctx context.Context, limit, offset int) ([]*response.VenueResponse, error) {
	venues, err := s.venueRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, util.ErrInternalServer()
	}

	responses := make([]*response.VenueResponse, len(venues))
	for i, venue := range venues {
		responses[i] = toVenueResponse(venue)
	}

	return responses, nil
}

// toVenueResponse converts an ent.Venue to VenueResponse
func toVenueResponse(venue *ent.Venue) *response.VenueResponse {
	var googlePlaceID *string
	if venue.GooglePlaceID != "" {
		googlePlaceID = &venue.GooglePlaceID
	}

	var surfaceType *string
	if venue.SurfaceType != "" {
		s := string(venue.SurfaceType)
		surfaceType = &s
	}

	var hourlyRate *int
	// HourlyRate is an int (0 means not set)
	if venue.HourlyRate > 0 {
		hourlyRate = &venue.HourlyRate
	}

	return &response.VenueResponse{
		ID:             venue.ID,
		Name:           venue.Name,
		Address:        venue.Address,
		Latitude:       venue.Latitude,
		Longitude:      venue.Longitude,
		GooglePlaceID:  googlePlaceID,
		FieldType:      string(venue.FieldType),
		SurfaceType:    surfaceType,
		Capacity:       venue.Capacity,
		HourlyRate:     hourlyRate,
		Facilities:      venue.Facilities,
		Images:         venue.Images,
		ContactInfo:    venue.ContactInfo,
		OperatingHours: venue.OperatingHours,
		Rating:         venue.Rating,
		TotalRatings:   venue.TotalRatings,
		IsVerified:     venue.IsVerified,
		CreatedAt:      venue.CreatedAt,
	}
}

