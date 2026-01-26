package models

import (
	"time"
)

type WorkshopPanels struct {
	WorkshopPanelNo uint      `gorm:"primaryKey;column:workshop_panel_no;type:int(11);autoIncrement" json:"workshop_panel_no"`
	WorkshopNo      uint      `gorm:"not null;column:workshop_no;type:int(11);" json:"workshop_no"`
	PanelNo         uint      `gorm:"not null;type:int(11)" json:"panel_no"`
	PanelName       string    `gorm:"type:varchar(255);not null" json:"panel_name"`
	IsLocked        bool      `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
	CreatedAt       time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy       *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	UpdatedAt       time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy  *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	UserProfile        *UserProfile `gorm:"foreignKey:WorkshopNo;references:UserProfileNo" json:"workshop,omitempty"`
	Panel              *Panel       `gorm:"foreignKey:PanelNo;references:PanelNo" json:"panel,omitempty"`
	CreatedByUser      *User        `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User        `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (WorkshopPanels) TableName() string {
	return "r_workshop_panels"
}
