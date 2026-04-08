package orders

import (
	"eclaim-workshop-deck-api/internal/models"
	"errors"

	"gorm.io/gorm"
)

// GetRepairingOrders retrieves repairing orders for a given workshop number.
func (r *Repository) GetRepairingOrders(id uint) ([]models.Order, error) {
	var orders []models.Order

	err := r.db.
		Preload("Workshop").
		Preload("Insurance").
		Preload("Client").
		Where("tr_orders.is_locked = ? AND tr_orders.workshop_no = ? AND tr_orders.status = ?", 0, id, "repairing").
		Order("tr_orders.order_no").
		Find(&orders).Error

	return orders, err
}

// CreateRepairHistoryTx creates a new repair history within a transaction.
func (r *Repository) CreateRepairHistoryTx(tx *gorm.DB, history *models.RepairHistory) error {
	return tx.Create(history).Error
}

// CreateRepairPhotosTx creates new repair photos within a transaction.
func (r *Repository) CreateRepairPhotosTx(tx *gorm.DB, photos []models.RepairPhoto) error {
	if len(photos) == 0 {
		return nil
	}
	return tx.Create(&photos).Error
}

// CreateOrderAndRequestTx creates a new order and request within a transaction.
func (r *Repository) CreateOrderAndRequestTx(tx *gorm.DB, orderAndRequest *models.OrderAndRequest) error {
	return tx.Create(orderAndRequest).Error
}

// CreateSparePartQuoteTx creates a new spare part quote within a transaction.
func (r *Repository) CreateSparePartQuoteTx(tx *gorm.DB, sparePartQuote *models.SparePartQuote) error {
	return tx.Create(sparePartQuote).Error
}

// CreateSparePartNegotiationHistoryTx creates a new spare part negotiation history within a transaction.
func (r *Repository) CreateSparePartNegotiationHistoryTx(tx *gorm.DB, sparePartNegotiationHistory *models.SparePartNegotiationHistory) error {
	return tx.Create(sparePartNegotiationHistory).Error
}

// GetSparePartQuoteTx retrieves the spare part quote for a given order request no within a transaction.
func (r *Repository) GetSparePartQuoteTx(tx *gorm.DB, orderRequestNo uint) (*models.SparePartQuote, error) {
	var quote models.SparePartQuote

	err := tx.Where("order_request_no = ?", orderRequestNo).First(&quote).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &quote, nil
}

// GetLatestSparePartNegotiationHistory retrieves the latest spare part negotiation history for a given spare part quote no.
func (r *Repository) GetLatestSparePartNegotiationHistory(db *gorm.DB, sparePartQuoteNo uint) (*models.SparePartNegotiationHistory, error) {
	var history models.SparePartNegotiationHistory

	err := db.Where("spare_part_quotes_no = ? AND is_locked = 0", sparePartQuoteNo).
		Order("round_count DESC").
		First(&history).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No negotiation history yet
		}
		return nil, err
	}

	return &history, nil
}

// FindSupplierFromID retrieves a supplier by its ID.
func (r *Repository) FindSupplierFromID(id uint) (models.Supplier, error) {
	var supplier models.Supplier

	err := r.db.
		Preload("Workshop").
		Preload("Province").
		Preload("City").
		Where("r_suppliers.supplier_no = ? AND is_locked = 0", id).
		Find(&supplier).Error

	return supplier, err
}

// UpdateSparePartQuoteTx updates a spare part quote within a transaction.
func (r *Repository) UpdateSparePartQuoteTx(tx *gorm.DB, sparePartQuote *models.SparePartQuote) error {
	return tx.Save(sparePartQuote).Error
}
