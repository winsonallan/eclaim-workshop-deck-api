package models

import "time"

type MOU struct {
	MouNo             uint      `gorm:"primaryKey;type:int(11);not null;autoIncrement" json:"mou_no"`
	MouDocumentNumber string    `gorm:"type:varchar(50);not null;" json:"mou_document_number"`
	MouExpiryDate     time.Time `gorm:"type:date" json:"mou_expiry_date"`
	InsurerNo         uint      `gorm:"type:int(11);not null" json:"insurer_no"`
	WorkshopNo        uint      `gorm:"type:int(11);not null" json:"workshop_no"`
	CreatedBy         *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	CreatedAt         time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedAt         time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy    *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	CreatedByUser       *User        `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser  *User        `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
	InsurerUserProfile  *UserProfile `gorm:"foreignKey:InsurerNo;references:UserProfileNo" json:"insurer_profile,omitempty"`
	WorkshopUserProfile *UserProfile `gorm:"foreignKey:WorkshopNo;references:UserProfileNo" json:"workshop_profile,omitempty"`
}

func (MOU) TableName() string {
	return "r_mou"
}
