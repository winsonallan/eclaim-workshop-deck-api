package settings

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

func (r *Repository) GetAccount(id uint) ([]models.User, error) {
	var account []models.User

	err := r.db.Where("is_locked = ?", 0).Where("user_no", id).Find(&account).Error

	return account, err
}

func (r *Repository) GetProfileDetails(id uint) ([]models.UserProfile, error) {
	var profile []models.UserProfile

	err := r.db.Where("is_locked = ?", 0).Where("user_profile_no", 1).Find(&profile).Error

	return profile, err
}

func (r *Repository) GetWorkshopDetails(id uint) ([]models.WorkshopDetails, error) {
	var workshopDetails []models.WorkshopDetails

	err := r.db.Where("is_locked = ?", 0).Where("workshop_details_no", 1).Find(&workshopDetails).Error

	return workshopDetails, err
}
