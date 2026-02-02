package models

type RolePermission struct {
	RolePermissionNo uint `gorm:"type:int(11);not null;primaryKey;autoIncrement" json:"role_permission_no"`
	RoleNo           uint `gorm:"type:int(11);not null;" json:"role_no"`
	PermissionNo     uint `gorm:"type:int(11);not null" json:"permission_no"`
	IsAllowed        bool `gorm:"type:tinyint(1);default:1" json:"is_allowed"`

	Role       *Role       `gorm:"foreignKey:RoleNo;references:RoleNo;" json:"role,omitempty"`
	Permission *Permission `gorm:"foreignKey:PermissionNo;references:PermissionNo" json:"permission,omitempty"`
}

func (RolePermission) TableName() string {
	return "r_role_permissions"
}
