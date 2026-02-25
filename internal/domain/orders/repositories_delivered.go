package orders

import "eclaim-workshop-deck-api/internal/models"

func (r *Repository) GetDeliveredOrders(id uint) ([]models.Order, error) {
	var orders []models.Order

	err := r.db.
		Preload("Workshop").
		Preload("Insurance").
		Preload("Client").
		Where("tr_orders.is_locked = ? AND tr_orders.workshop_no = ? AND tr_orders.status = ?", 0, id, "delivered").
		Order("tr_orders.order_no").
		Find(&orders).Error

	return orders, err
}
