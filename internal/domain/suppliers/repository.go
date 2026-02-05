package suppliers

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

func (r *Repository) GetSuppliers() ([]models.Supplier, error) {
	var supplier []models.Supplier

	err := r.db.
		Preload("Workshop").
		Preload("CreatedByUser").
		Preload("LastModifiedByUser").
		Preload("City").
		Preload("Province").
		Where("is_locked = ?", 0).
		Order("supplier_name asc").
		Find(&supplier).Error

	return supplier, err
}

func (r *Repository) GetWorkshopSuppliers(id uint) ([]models.Supplier, error) {
	var suppliers []models.Supplier

	err := r.db.
		Preload("Workshop").
		Preload("CreatedByUser").
		Preload("LastModifiedByUser").
		Preload("City").
		Preload("Province").
		Where("is_locked = ?", 0).
		Where("workshop_no = ?", id).
		Find(&suppliers).Error

	return suppliers, err
}

func (r *Repository) FindSupplierByID(id uint) (*models.Supplier, error) {
	var supplier models.Supplier

	err := r.db.
		Preload("Workshop").
		Preload("CreatedByUser").
		Preload("LastModifiedByUser").
		Preload("City").
		Preload("Province").
		Where("is_locked = ?", 0).
		Where("supplier_no", id).
		First(&supplier).Error

	return &supplier, err
}

func (r *Repository) AddSupplier(supplier *models.Supplier) error {
	return r.db.Create(supplier).Error
}

// Update
func (r *Repository) UpdateSupplier(supplier *models.Supplier) error {
	return r.db.Save(supplier).Error
}
