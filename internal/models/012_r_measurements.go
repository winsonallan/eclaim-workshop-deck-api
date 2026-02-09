package models

import "time"

type Measurement struct {
	MeasurementNo  uint      `gorm:"primaryKey;type:int(11);not null;autoIncrement;column:measurement_no" json:"measurement_no"`
	PanelPricingNo uint      `gorm:"type:int(11);not null;column:panel_pricing_no;" json:"panel_pricing_no"`
	ConditionText  string    `gorm:"type:varchar(255);not null;column:condition_text" json:"condition_text"`
	Notes          string    `gorm:"type:varchar(255);null" json:"notes"`
	LaborFee       uint      `gorm:"type:int(11);not null;column:labor_fee" json:"labor_fee"`
	IsLocked       bool      `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
	CreatedBy      *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	CreatedAt      time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedAt      time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	PanelPricing       *PanelPricing `gorm:"foreignKey:PanelPricingNo;references:PanelPricingNo" json:"panel_pricing,omitempty"`
	CreatedByUser      *User         `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User         `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (Measurement) TableName() string {
	return "r_measurements"
}
