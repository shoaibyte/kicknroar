package request

// UpdateUserRequest represents a user profile update request
// @Description User profile update request - all fields are optional. Example: {"full_name":"John Doe Updated","skill_level":"advanced","profile_image_url":"https://example.com/avatar.jpg","preferred_locations":["Dhanmondi","Gulshan","Banani"]}
type UpdateUserRequest struct {
	FullName          *string   `json:"full_name,omitempty" example:"John Doe Updated" validate:"omitempty,min=2,max=100"`
	SkillLevel        *string   `json:"skill_level,omitempty" example:"advanced" validate:"omitempty,oneof=beginner intermediate advanced professional"`
	ProfileImageURL   *string   `json:"profile_image_url,omitempty" example:"https://example.com/avatar.jpg"`
	PreferredLocations []string `json:"preferred_locations,omitempty" example:"Dhanmondi,Gulshan,Banani"`
}

