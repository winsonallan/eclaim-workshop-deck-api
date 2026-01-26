package models

import (
	"time"
)

type WorkshopPics struct {
	WorkshopPicNoNo   uint      `gorm:"primaryKey;column:workshop_pic_no;type:int(11);autoIncrement" json:"workshop_pic_no"`
	WorkshopDetailsNo uint      `gorm:"not null;column:workshop_details_no;type:int(11);" json:"workshop_details_no"`
	WorkshopPicName   string    `gorm:"not null;type:varchar(255)" json:"workshop_pic_name"`
	WorkshopTitle     string    `gorm:"not null;type:varchar(50)" json:"workshop_title"`
	Phone             string    `gorm:"not null;type:varchar(30)" json:"phone"`
	Email             string    `gorm:"not null;type:varchar(255)" json:"email"`
	IsLocked          bool      `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
	CreatedAt         time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy         *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	UpdatedAt         time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy    *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	WorkshopDetails    *WorkshopDetails `gorm:"foreignKey:WorkshopDetailsNo;references:WorkshopDetailsNo" json:"workshop_details,omitempty"`
	CreatedByUser      *User            `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User            `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (WorkshopPics) TableName() string {
	return "r_workshop_pics"
}
