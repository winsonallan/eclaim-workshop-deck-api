package orders

import (
	"eclaim-workshop-deck-api/internal/models"

	"gorm.io/gorm"
)

func (r *Repository) GetNegotiatingOrders(id uint) ([]models.Order, error) {
	var orders []models.Order

	err := r.db.
		Preload("Workshop").
		Preload("Insurance").
		Preload("Client").
		Where("tr_orders.is_locked = ? AND tr_orders.workshop_no = ? AND tr_orders.status = ? OR tr_orders.status = ?", 0, id, "negotiating", "additional_work").
		Order("tr_orders.order_no").
		Find(&orders).Error

	return orders, err
}

func (r *Repository) CancelNegotiation(tx *gorm.DB, order *models.Order) error {
	return tx.Save(order).Error
}
