package repository

import (
	"context"

	"kicknroar/internal/database"
	"kicknroar/internal/ent"
	"kicknroar/internal/ent/user"
	"kicknroar/internal/ent/userstats"
)

// UserRepository handles user data access
type UserRepository struct {
	client *ent.Client
}

// NewUserRepository creates a new user repository
func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{client: client}
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, email, passwordHash, fullName, phone string) (*ent.User, error) {
	u, err := r.client.User.
		Create().
		SetEmail(email).
		SetPasswordHash(passwordHash).
		SetFullName(fullName).
		SetPhone(phone).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			if database.IsUniqueConstraintError(err) {
				// Check if it's email or phone
				existing, _ := r.FindByEmail(ctx, email)
				if existing != nil {
					return nil, err
				}
			}
		}
		return nil, err
	}

	// Create user stats
	_, err = r.client.UserStats.
		Create().
		SetUserID(u.ID).
		Save(ctx)
	if err != nil {
		// Log error but don't fail user creation
	}

	return u, nil
}

// FindByID finds a user by ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*ent.User, error) {
	uid, err := database.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	return r.client.User.
		Query().
		Where(user.ID(uid)).
		Only(ctx)
}

// FindByEmail finds a user by email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
}

// FindByPhone finds a user by phone
func (r *UserRepository) FindByPhone(ctx context.Context, phone string) (*ent.User, error) {
	return r.client.User.
		Query().
		Where(user.Phone(phone)).
		Only(ctx)
}

// UpdateLastLogin updates the last login timestamp
func (r *UserRepository) UpdateLastLogin(ctx context.Context, id string) error {
	uid, err := database.ParseUUID(id)
	if err != nil {
		return err
	}

	_, err = r.client.User.
		UpdateOneID(uid).
		SetLastLoginAt(database.Now()).
		Save(ctx)
	return err
}

// UpdateProfile updates user profile
func (r *UserRepository) UpdateProfile(ctx context.Context, id string, updates map[string]interface{}) (*ent.User, error) {
	uid, err := database.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	update := r.client.User.UpdateOneID(uid)
	
	if fullName, ok := updates["full_name"].(string); ok {
		update = update.SetFullName(fullName)
	}
	if skillLevel, ok := updates["skill_level"].(string); ok {
		update = update.SetSkillLevel(user.SkillLevel(skillLevel))
	}
	if profileImageURL, ok := updates["profile_image_url"].(string); ok {
		update = update.SetProfileImageURL(profileImageURL)
	}
	if preferredLocations, ok := updates["preferred_locations"].([]string); ok {
		update = update.SetPreferredLocations(preferredLocations)
	}

	return update.Save(ctx)
}

// GetStats gets user statistics
func (r *UserRepository) GetStats(ctx context.Context, userID string) (*ent.UserStats, error) {
	uid, err := database.ParseUUID(userID)
	if err != nil {
		return nil, err
	}

	return r.client.UserStats.
		Query().
		Where(userstats.UserID(uid)).
		Only(ctx)
}

