package auth

import (
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *Repository) ChangePassword(user *User) error {
	return r.db.Model(&User{}).
		Where("user_no = ?", user.UserNo).
		Updates(map[string]interface{}{
			"password":           user.Password,
			"last_modified_by":   user.LastModifiedBy,
			"last_modified_date": time.Now(), // if you use GORM timestamps
		}).Error
}
