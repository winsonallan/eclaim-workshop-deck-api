package models

import "time"

type NegotiationPhotos struct {
	NegotiationPhotoNo   uint      `gorm:"primaryKey;not null;autoIncrement;type:int(11)" json:"negotiation_photo_no"`
	NegotiationHistoryNo uint      `gorm:"not null;type:int(11)" json:"repair_history_no"`
	PhotoUrl             string    `gorm:"type:varchar(255);not null" json:"photo_url"`
	CreatedAt            time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy            *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	UpdatedAt            time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy       *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	NegotiationHistory *NegotiationHistory `gorm:"foreignKey:NegotiationHistoryNo;references:NegotiationHistoryNo" json:"negotiation_history"`
	CreatedByUser      *User               `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User               `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (NegotiationPhotos) TableName() string {
	return "tr_negotiation_photos"
}
