package panels

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

func (r *Repository) GetAllPanels() ([]models.Panel, error) {
	var panels []models.Panel

	err := r.db.Where("is_locked = ?", 0).Order("panel_name asc").Find(&panels).Error

	return panels, err
}
