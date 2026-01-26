package migrations

import (
	"fmt"
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
	migrateOrFail(db, &models.User{})
	migrateOrFail(db, &models.UserProfile{})
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

	// Now manually create the foreign key constraints you actually want
	createForeignKeyConstraints(db)
}

func createForeignKeyConstraints(db *gorm.DB) {
	// First, drop any incorrectly created constraints
	fixUserProfileConstraints(db)

	// Demo tables foreign keys
	createConstraintSafe(db, &posts.Post{}, "User")

	// City Table foreign keys
	createConstraintSafe(db, &models.City{}, "Province")

	// User table foreign keys
	createConstraintSafe(db, &models.User{}, "CreatedByUser")
	createConstraintSafe(db, &models.User{}, "LastModifiedByUser")

	// UserProfile table foreign keys
	createConstraintSafe(db, &models.UserProfile{}, "CreatedByUser")
	createConstraintSafe(db, &models.UserProfile{}, "LastModifiedByUser")
	createConstraintSafe(db, &models.UserProfile{}, "City")

	// Manually create the User -> UserProfile foreign key correctly
	createUserProfileConstraint(db)

	// MOU table foreign keys
	createConstraintSafe(db, &models.MOU{}, "CreatedByUser")
	createConstraintSafe(db, &models.MOU{}, "LastModifiedByUser")
	createConstraintSafe(db, &models.MOU{}, "InsurerUserProfile")
	createConstraintSafe(db, &models.MOU{}, "WorkshopUserProfile")

	// Panel Pricing foreign keys
	createConstraintSafe(db, &models.PanelPricing{}, "Workshop")
	createConstraintSafe(db, &models.PanelPricing{}, "Insurer")
	createConstraintSafe(db, &models.PanelPricing{}, "Panels")
	createConstraintSafe(db, &models.PanelPricing{}, "Mou")
	createConstraintSafe(db, &models.PanelPricing{}, "CreatedByUser")
	createConstraintSafe(db, &models.PanelPricing{}, "LastModifiedByUser")

	// ... rest of your constraints remain the same
}

func fixUserProfileConstraints(db *gorm.DB) {
	// Drop any incorrect constraints that may have been created
	incorrectConstraints := []string{
		"fk_r_users_user_profile",
		"fk_r_user_profiles_users",
	}

	for _, constraintName := range incorrectConstraints {
		// Try dropping from r_user_profiles
		db.Exec(fmt.Sprintf("ALTER TABLE r_user_profiles DROP FOREIGN KEY IF EXISTS %s", constraintName))
		// Try dropping from r_users (in case it's there)
		db.Exec(fmt.Sprintf("ALTER TABLE r_users DROP FOREIGN KEY IF EXISTS %s", constraintName))
	}

	log.Println("Cleaned up any incorrect UserProfile constraints")
}

func createUserProfileConstraint(db *gorm.DB) {
	// Check if the constraint already exists on r_users
	var count int64
	db.Raw(`
		SELECT COUNT(*) 
		FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE 
		WHERE TABLE_SCHEMA = DATABASE() 
		AND TABLE_NAME = 'r_users' 
		AND COLUMN_NAME = 'user_profile_no' 
		AND REFERENCED_TABLE_NAME = 'r_user_profiles'
	`).Scan(&count)

	if count == 0 {
		// Create the correct FK: r_users.user_profile_no -> r_user_profiles.user_profile_no
		err := db.Exec(`
			ALTER TABLE r_users 
			ADD CONSTRAINT fk_r_users_user_profile 
			FOREIGN KEY (user_profile_no) 
			REFERENCES r_user_profiles(user_profile_no) 
			ON DELETE SET NULL 
			ON UPDATE CASCADE
		`).Error

		if err != nil {
			log.Printf("Warning: Could not create User -> UserProfile constraint: %v", err)
		} else {
			log.Println("Manually created User -> UserProfile constraint")
		}
	} else {
		log.Println("User -> UserProfile constraint already exists")
	}
}

func createConstraintSafe(db *gorm.DB, model interface{}, field string) {
	if !db.Migrator().HasConstraint(model, field) {
		if err := db.Migrator().CreateConstraint(model, field); err != nil {
			log.Printf("Warning: Could not create constraint %s for %T: %v", field, model, err)
		} else {
			log.Printf("✓ Created constraint %s for %T", field, model)
		}
	} else {
		log.Printf("ℹ Constraint %s for %T already exists", field, model)
	}
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
