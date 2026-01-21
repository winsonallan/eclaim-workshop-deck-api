package authdemo

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *UserDemo) error {
	return r.db.Create(user).Error
}

func (r *Repository) FindByEmail(email string) (*UserDemo, error) {
	var user UserDemo
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *Repository) FindByID(id uint) (*UserDemo, error) {
	var user UserDemo
	err := r.db.First(&user, id).Error
	return &user, err
}