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
		Preload("Client").
		Where("tr_orders.is_locked = ?", 0).
		Order("tr_orders.order_no").
		Find(&orders).Error

	return orders, err
}

func (r *Repository) GetIncomingOrders(id uint) ([]models.Order, error) {
	var orders []models.Order

	err := r.db.
		Preload("Workshop").
		Preload("Insurance").
		Preload("Client").
		Where("tr_orders.is_locked = ? AND tr_orders.workshop_no = ? AND tr_orders.status = ?", 0, id, "incoming").
		Order("tr_orders.order_no").
		Find(&orders).Error

	return orders, err
}

func (r *Repository) ViewOrderDetails(id uint) (models.Order, error) {
	var order models.Order

	err := r.db.
		Preload("Workshop").
		Preload("Insurance").
		Preload("Invoice").
		Preload("Client").
		Preload("WorkOrders").
		Preload("WorkOrders.OrderPanels").
		Preload("WorkOrders.OrderPanels.RepairHistory").
		Preload("WorkOrders.OrderPanels.NegotiationHistory").
		Preload("WorkOrders.OrderPanels.RepairHistory.RepairPhotos").
		Preload("WorkOrders.OrderPanels.RepairHistory.OrdersAndRequests").
		Preload("WorkOrders.OrderPanels.RepairHistory.OrdersAndRequests.SparePartQuotes").
		Preload("WorkOrders.OrderPanels.RepairHistory.OrdersAndRequests.SparePartQuotes.SparePartNegotiationHistory").
		Where("tr_orders.is_locked = ? AND tr_orders.order_no = ?", 0, id).
		Order("tr_orders.order_no").
		Find(&order).Error

	return order, err
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

func (r *Repository) FindWorkOrderById(id uint) (*models.WorkOrder, error) {
	var workOrder models.WorkOrder

	err := r.db.
		Preload("Order").
		Preload("OrderPanels").
		Preload("CreatedByUser").
		Preload("LastModifiedByUser").
		Where("work_order_no = ? AND is_locked = 0", id).
		First(&workOrder).Error

	return &workOrder, err
}

func (r *Repository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *Repository) AddClient(client *models.Client) error {
	return r.db.Create(client).Error
}

func (r *Repository) CreateWorkOrder(workOrder *models.WorkOrder) error {
	return r.db.Create(workOrder).Error
}

func (r *Repository) CreateOrderPanel(orderPanel *models.OrderPanel) error {
	return r.db.Create(orderPanel).Error
}

func (r *Repository) CreateOrderPanelsBatch(orderPanels []*models.OrderPanel) error {
	if len(orderPanels) == 0 {
		return nil
	}
	return r.db.Create(&orderPanels).Error
}

func (r *Repository) UpdateWorkOrder(workOrder *models.WorkOrder) error {
	return r.db.Save(workOrder).Error
}
