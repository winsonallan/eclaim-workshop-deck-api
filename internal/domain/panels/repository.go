package panels

import (
	"eclaim-workshop-deck-api/internal/models"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllPanels() ([]models.Panel, error) {
	var panels []models.Panel

	err := r.db.Where("is_locked = ?", 0).Order("panel_name asc").Find(&panels).Error

	return panels, err
}

func (r *Repository) GetAllWorkshopPanels(id uint) ([]models.WorkshopPanels, error) {
	var workshopPanels []models.WorkshopPanels

	err := r.db.
		Where("is_locked = ?", 0).
		Where("(workshop_no IS NULL OR workshop_no = ?)", id).
		Find(&workshopPanels).Error

	return workshopPanels, err
}

func (r *Repository) GetMOUs(insID, woID, mouID uint, activeOnly bool) ([]models.MOU, error) {
	var mous []models.MOU
	query := r.db.Preload("InsurerUserProfile").
		Preload("WorkshopUserProfile").
		Preload("CreatedByUser").
		Where("is_locked = ?", 0)

	// Add filters only if the ID is provided (not 0)
	if insID != 0 {
		query = query.Where("insurer_no = ?", insID)
	}
	if woID != 0 {
		query = query.Where("workshop_no = ?", woID)
	}
	if mouID != 0 {
		query = query.Where("mou_no = ?", mouID)
	}

	// Add expiry filter only if requested
	if activeOnly {
		query = query.Where("(mou_expiry_date > ? OR mou_expiry_date IS NULL)", time.Now())
	}

	err := query.Find(&mous).Error
	return mous, err
}

func (r *Repository) GetPanelPricings(insID, woID, mouID uint) ([]models.PanelPricing, error) {
	var panelPricings []models.PanelPricing
	query := r.db.Preload("Insurer").
		Preload("Workshop").
		Preload("WorkshopPanels").
		Preload("Mou").
		Preload("CreatedByUser").
		Preload("LastModifiedByUser")

	// Add filters only if the ID is provided (not 0)
	if insID != 0 {
		query = query.Where("insurer_no = ?", insID)
	}
	if woID != 0 {
		query = query.Where("workshop_no = ?", woID)
	}
	if mouID != 0 {
		query = query.Where("mou_no = ?", mouID)
	}

	err := query.Find(&panelPricings).Error
	return panelPricings, err
}

func (r *Repository) FindMOUByID(id uint) (*models.MOU, error) {
	var mou models.MOU

	err := r.db.
		Preload("InsurerUserProfile").
		Preload("WorkshopUserProfile").
		Preload("CreatedByUser").
		Where("is_locked = ?", 0).
		Where("mou_no", id).
		First(&mou).Error

	return &mou, err
}

func (r *Repository) CreateMOU(mou *models.MOU) error {
	return r.db.Create(mou).Error
}
