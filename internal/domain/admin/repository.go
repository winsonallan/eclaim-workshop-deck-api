package admin

import (
	"eclaim-workshop-deck-api/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUserProfile(userProfile *models.UserProfile) error {
	return r.db.Create(userProfile).Error
}

func (r *Repository) FindUserProfileByID(id uint) (*models.UserProfile, error) {
	var userProfile models.UserProfile
	err := r.db.First(&userProfile, id).Error
	return &userProfile, err
}
