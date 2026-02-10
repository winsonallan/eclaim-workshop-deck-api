package orders

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

func (r *Repository) GetOrders() ([]models.Order, error) {
	var orders []models.Order

	err := r.db.
		Preload("Workshop").
		Preload("Insurance").
		Preload("Invoice").
		Preload("Client").
		Preload("WorkOrders").
		Where("is_locked = ?", 0).
		Order("order_no").
		Find(&orders).Error

	return orders, err
}

func (r *Repository) GetIncomingOrders(id uint) ([]models.Order, error) {
	var orders []models.Order

	err := r.db.
		Preload("Workshop").
		Preload("Insurance").
		Preload("Invoice").
		Preload("Client").
		Preload("WorkOrders").
		Where("is_locked = ? AND workshop_no = ? AND status = ?", 0, id, "incoming").
		Order("order_no").
		Find(&orders).Error

	return orders, err
}

func (r *Repository) FindClientById(id uint) (*models.Client, error) {
	var client models.Client

	err := r.db.
		Preload("City").
		Preload("CreatedByUser").
		Preload("LastModifiedByUser").
		Where("client_no = ? AND is_locked = 0", id).
		First(&client).Error

	return &client, err
}

func (r *Repository) FindOrderById(id uint) (*models.Order, error) {
	var order models.Order

	err := r.db.
		Preload("Workshop").
		Preload("Insurance").
		Preload("Invoice").
		Preload("Client").
		Preload("CreatedByUser").
		Preload("LastModifiedByUser").
		Where("order_no = ? AND is_locked = 0", id).
		First(&order).Error

	return &order, err
}

func (r *Repository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *Repository) AddClient(client *models.Client) error {
	return r.db.Create(client).Error
}
