package models

import "time"

type Permission struct {
	PermissionNo   uint      `gorm:"primaryKey;column:permission_no;type:int(11);autoIncrement" json:"permission_no"`
	PageKey        string    `gorm:"type:varchar(50);not null;column:page_key" json:"page_key"`
	ActionKey      string    `gorm:"type:varchar(50);not null;column:action_key" json:"action_key"`
	Description    string    `gorm:"type:mediumtext;null" json:"description"`
	CreatedBy      *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	CreatedAt      time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedAt      time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	CreatedByUser      *User `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (Permission) TableName() string {
	return "m_permissions"
}
