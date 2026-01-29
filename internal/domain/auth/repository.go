package auth

import (
	"eclaim-workshop-deck-api/internal/models"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("UserProfile").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *Repository) FindByEmailAndUsername(email, username string) (*models.User, error) {
	var user models.User
	err := r.db.
		Where("email = ? AND user_id = ?", email, username).
		Where("is_locked = ?", 0).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdatePassword(userID uint, hashedPassword string) error {
	return r.db.Model(&models.User{}).Where("user_no = ?", userID).Update("password", hashedPassword).Error
}

func (r *Repository) FindByUserNo(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("UserProfile").Where("user_no = ?", id).First(&user).Error
	return &user, err
}

func (r *Repository) ChangePassword(user *models.User) error {
	return r.db.Model(&models.User{}).
		Where("user_no = ?", user.UserNo).
		Updates(map[string]interface{}{
			"password":           user.Password,
			"last_modified_by":   user.LastModifiedBy,
			"last_modified_date": time.Now(), // if you use GORM timestamps
		}).Error
}

func (r *Repository) UpdateAccount(user *models.User) error {
	return r.db.Save(user).Error
}
