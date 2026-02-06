package models

import "time"

type Role struct {
	RoleNo         uint      `gorm:"primaryKey;column:role_no;type:int(11);autoIncrement" json:"role_no"`
	RoleName       string    `gorm:"type:varchar(50);not null;column:role_name" json:"role_name"`
	RoleType       string    `gorm:"type:enum('abb','insurer','workshop');not_null;column:role_type" json:"role_type"`
	CreatedBy      *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	CreatedAt      time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedAt      time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	CreatedByUser      *User `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (Role) TableName() string {
	return "m_roles"
}
