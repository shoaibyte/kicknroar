package request

// CreateVenueRequest represents a venue creation request
// @Description Venue creation request with location, facilities, and contact information
type CreateVenueRequest struct {
	Name          string                 `json:"name" validate:"required,min=2,max=100" example:"Dhanmondi Futsal Arena"`
	Address       string                 `json:"address" validate:"required" example:"Road 27, Dhanmondi, Dhaka"`
	Latitude      float64                `json:"latitude" validate:"required" example:"23.7465"`
	Longitude     float64                `json:"longitude" validate:"required" example:"90.3760"`
	GooglePlaceID *string                `json:"google_place_id,omitempty" example:"ChIJN1t_tDeuEmsRUsoyG83frY4"`
	FieldType     string                 `json:"field_type" validate:"required,oneof=futsal football astro" example:"futsal"`
	SurfaceType   *string                `json:"surface_type,omitempty" validate:"omitempty,oneof=grass artificial concrete" example:"artificial"`
	Capacity      int                    `json:"capacity" validate:"required,min=1" example:"10"`
	HourlyRate    *int                   `json:"hourly_rate,omitempty" example:"2000"`
	Facilities    []string               `json:"facilities,omitempty" example:"parking,changing_room,water_dispenser"`
	ContactInfo   map[string]interface{} `json:"contact_info,omitempty"`
	OperatingHours map[string]interface{} `json:"operating_hours,omitempty"`
}

// UpdateVenueRequest represents a venue update request
// @Description Venue update request - all fields are optional
type UpdateVenueRequest struct {
	Name          *string                `json:"name,omitempty" validate:"omitempty,min=2,max=100" example:"Dhanmondi Futsal Arena - Updated"`
	Address       *string                `json:"address,omitempty" example:"Road 27, Dhanmondi, Dhaka - Updated"`
	Latitude      *float64               `json:"latitude,omitempty" example:"23.7465"`
	Longitude     *float64               `json:"longitude,omitempty" example:"90.3760"`
	FieldType     *string                `json:"field_type,omitempty" validate:"omitempty,oneof=futsal football astro" example:"futsal"`
	SurfaceType   *string                `json:"surface_type,omitempty" validate:"omitempty,oneof=grass artificial concrete" example:"artificial"`
	Capacity      *int                   `json:"capacity,omitempty" validate:"omitempty,min=1" example:"12"`
	HourlyRate    *int                   `json:"hourly_rate,omitempty" example:"2500"`
	Facilities    []string               `json:"facilities,omitempty" example:"parking,changing_room,water_dispenser,shower"`
	ContactInfo   map[string]interface{} `json:"contact_info,omitempty"`
	OperatingHours map[string]interface{} `json:"operating_hours,omitempty"`
}

