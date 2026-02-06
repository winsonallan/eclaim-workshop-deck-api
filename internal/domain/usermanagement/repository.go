package usermanagement

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

func (r *Repository) GetWorkshopUsers(id uint) ([]models.User, error) {
	var workshopUsers []models.User

	err := r.db.
		Preload("Role").
		Preload("CreatedByUser").
		Where("is_locked = ? AND user_profile_no = ?", 0, id).
		Find(&workshopUsers).Error

	return workshopUsers, err
}

func (r *Repository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *Repository) GetRoles(roleType string) ([]models.Role, error) {
	var roles []models.Role

	err := r.db.Where("role_type = ?", roleType).Find(&roles).Error

	return roles, err
}

func (r *Repository) GetRole(roleID uint) (models.Role, error) {
	var role models.Role

	err := r.db.Where("role_no", roleID).First(&role).Error

	return role, err
}

func (r *Repository) GetRolePermissions(roleID uint) ([]models.RolePermission, error) {
	var rolePermissions []models.RolePermission

	err := r.db.
		Preload("Role").
		Preload("Permission").
		Where("role_no = ?", roleID).
		Find(&rolePermissions).Error

	return rolePermissions, err
}
