package location

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

func (r *Repository) GetCities() ([]models.City, error) {
	var cities []models.City

	err := r.db.Where("is_locked = ?", 0).Order("city_name").Find(&cities).Error

	return cities, err
}

func (r *Repository) GetProvinces() ([]models.Province, error) {
	var provinces []models.Province

	err := r.db.Where("is_locked = ?", 0).Order("province_name").Find(&provinces).Error

	return provinces, err
}
