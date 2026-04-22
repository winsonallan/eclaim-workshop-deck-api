package invoices

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

// FindOrdersByNos fetches the orders matching the given order numbers.
func (r *Repository) FindOrdersByNos(orderNos []uint) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.
		Where("order_no IN ?", orderNos).
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// CreateInvoiceWithInstallments runs everything in a single transaction:
//  1. Insert the Invoice row and capture the auto-incremented InvoiceNo.
//  2. Attach InvoiceNo to every InvoiceInstallment and bulk-insert them
//     (skipped when the slice is empty, i.e. single-payment invoice).
//  3. Update every Order in orderNos to point at the new InvoiceNo.
//  4. Reload the Invoice with its InvoiceInstallments preloaded and return it.
func (r *Repository) CreateInvoiceWithInstallments(
	invoice *models.Invoice,
	installments []models.InvoiceInstallment,
	orderNos []uint,
) (*models.Invoice, error) {

	err := r.db.Transaction(func(tx *gorm.DB) error {

		// 1. Create the invoice record.
		if err := tx.Create(invoice).Error; err != nil {
			return err
		}

		// 2. Insert installments (if any).
		if len(installments) > 0 {
			for i := range installments {
				installments[i].InvoiceNo = invoice.InvoiceNo
			}
			if err := tx.Create(&installments).Error; err != nil {
				return err
			}
		}

		// 3. Link orders to the new invoice.
		if err := tx.Model(&models.Order{}).
			Where("order_no IN ?", orderNos).
			Update("invoice_no", invoice.InvoiceNo).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// 4. Reload with associations.
	var created models.Invoice
	if err := r.db.
		Preload("InvoiceInstallments").
		Preload("Client").
		First(&created, invoice.InvoiceNo).Error; err != nil {
		return nil, err
	}

	return &created, nil
}
