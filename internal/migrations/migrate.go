package migrations

import (
	"log"

	"eclaim-workshop-deck-api/internal/domain/authdemo"
	"eclaim-workshop-deck-api/internal/domain/posts"
	"eclaim-workshop-deck-api/internal/models"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	// Sample Tables
	migrateOrFail(db, &authdemo.UserDemo{})
	migrateOrFail(db, &authdemo.APIKey{})
	migrateOrFail(db, &posts.Post{})

	// API Tables
	migrateOrFail(db, &models.APIKey{})

	// Master Tables
	migrateOrFail(db, &models.Panel{})
	migrateOrFail(db, &models.Permission{})
	migrateOrFail(db, &models.Province{})
	migrateOrFail(db, &models.Role{})

	// Reference Tables (r_)
	migrateOrFail(db, &models.City{})
	migrateOrFail(db, &models.UserProfile{})
	migrateOrFail(db, &models.User{})
	migrateOrFail(db, &models.MOU{})
	migrateOrFail(db, &models.PanelPricing{})
	migrateOrFail(db, &models.Measurement{})
	migrateOrFail(db, &models.RolePermission{})
	migrateOrFail(db, &models.Supplier{})

	// Transaction Tables (tr_)
	migrateOrFail(db, &models.Delivery{})
	migrateOrFail(db, &models.Invoice{})
	migrateOrFail(db, &models.InvoiceInstallment{})
	migrateOrFail(db, &models.Order{})
	migrateOrFail(db, &models.NegotiationHistory{})
	migrateOrFail(db, &models.PaymentRecord{})
	migrateOrFail(db, &models.PickupReminder{})
	migrateOrFail(db, &models.RepairHistory{})
	migrateOrFail(db, &models.OrderAndRequest{})
	migrateOrFail(db, &models.RepairPhoto{})
	migrateOrFail(db, &models.Review{})
	migrateOrFail(db, &models.WorkOrder{})
	migrateOrFail(db, &models.OrderPanel{})
	migrateOrFail(db, &models.SparePartQuote{})
	migrateOrFail(db, &models.SparePartNegotiationHistory{})
}

func migrateOrFail(db *gorm.DB, models ...any) {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf(
				"Database migration failed for model %T: %v",
				model,
				err,
			)
		}
	}
}
