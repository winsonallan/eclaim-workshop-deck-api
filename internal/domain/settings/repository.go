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

// READ
func (r *Repository) GetAccount(id uint) ([]models.User, error) {
	var account []models.User

	err := r.db.Where("is_locked = ?", 0).Where("user_no", id).Find(&account).Error

	return account, err
}

func (r *Repository) GetProfileDetails(id uint) (*models.UserProfile, error) {
	var profile *models.UserProfile

	err := r.db.Where("is_locked = ?", 0).Where("user_profile_no", id).First(&profile).Error

	return profile, err
}

func (r *Repository) GetWorkshopDetails(id uint) ([]models.WorkshopDetails, error) {
	var workshopDetails []models.WorkshopDetails

	err := r.db.Where("is_locked = ?", 0).Where("workshop_details_no", id).Find(&workshopDetails).Error

	return workshopDetails, err
}

func (r *Repository) GetWorkshopPICs(id uint) ([]models.WorkshopPics, error) {
	var workshopPics []models.WorkshopPics

	err := r.db.Preload("WorkshopDetails").Where("is_locked = ?", 0).Where("workshop_details_no", id).Find(&workshopPics).Error

	return workshopPics, err
}

func (r *Repository) FindWorkshopDetailsByID(id uint) (*models.WorkshopDetails, error) {
	var details models.WorkshopDetails
	if err := r.db.
		Preload("UserProfile").
		Preload("CreatedByUser").
		Preload("LastModifiedByUser").
		First(&details, "workshop_details_no = ?", id).Error; err != nil {
		return nil, err
	}
	return &details, nil
}

func (r *Repository) FindWorkshopPICByID(id uint) (*models.WorkshopPics, error) {
	var pic models.WorkshopPics
	if err := r.db.
		Preload("WorkshopDetails").
		Preload("CreatedByUser").
		Preload("LastModifiedByUser").
		Where("is_locked = ?", 0).
		First(&pic, "workshop_pic_no = ?", id).Error; err != nil {
		return nil, err
	}
	return &pic, nil
}

// CREATE
func (r *Repository) CreateWorkshopDetails(details *models.WorkshopDetails) error {
	return r.db.Create(details).Error
}

func (r *Repository) CreateWorkshopPICs(workshopPICs *models.WorkshopPics) error {
	return r.db.Create(workshopPICs).Error
}

// UPDATE
func (r *Repository) UpdateWorkshopDetails(workshopDetails *models.WorkshopDetails) error {
	return r.db.Save(workshopDetails).Error
}

func (r *Repository) UpdateUserProfile(userProfile *models.UserProfile) error {
	return r.db.Save(userProfile).Error
}
