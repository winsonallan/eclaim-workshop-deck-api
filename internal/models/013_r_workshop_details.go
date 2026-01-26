package models

import (
	"time"
)

type WorkshopDetails struct {
	WorkshopDetailsNo uint      `gorm:"primaryKey;column:workshop_details_no;type:int(11);autoIncrement" json:"workshop_details_no"`
	UserProfileNo     uint      `gorm:"null;column:user_profile_no;type:int(11);" json:"user_profile_no"`
	Capacity          uint      `gorm:"type:tinyint(3);null" json:"capacity"`
	Description       string    `gorm:"type:mediumtext;null" json:"description"`
	IsAuthorized      bool      `gorm:"type:tinyint(1);null;default:0" json:"is_authorized"`
	Specialist        string    `gorm:"type:varchar(50);null" json:"specialist"`
	IsLocked          bool      `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
	CreatedAt         time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy         *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	UpdatedAt         time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy    *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	UserProfile        *UserProfile `gorm:"foreignKey:UserProfileNo;references:UserProfileNo" json:"user_profile,omitempty"`
	CreatedByUser      *User        `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User        `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (WorkshopDetails) TableName() string {
	return "r_workshop_details"
}
