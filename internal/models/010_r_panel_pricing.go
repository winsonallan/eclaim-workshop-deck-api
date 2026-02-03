package models

import "time"

type PanelPricing struct {
	PanelPricingNo   uint      `gorm:"primaryKey;type:int(11);not null;autoIncrement" json:"panel_pricing_no"`
	WorkshopNo       uint      `gorm:"type:int(11);not null;" json:"workshop_no"`
	InsurerNo        *uint     `gorm:"type:int(11);null;" json:"insurer_no"`
	MouNo            *uint     `gorm:"type:int(11);null;" json:"mou_no"`
	WorkshopPanelNo  uint      `gorm:"type:int(11);not null;" json:"workshop_panel_no"`
	ServiceType      string    `gorm:"type:enum('repair','replacement');not null" json:"service_type"`
	IsFixedPrice     bool      `gorm:"type:tinyint(1);not null" json:"is_fixed_price"`
	VehicleRangeLow  uint      `gorm:"type:bigint(20);default:0;not null" json:"vehicle_range_low"`
	VehicleRangeHigh uint      `gorm:"type:bigint(20);default:999999999999;not null" json:"vehicle_range_high"`
	SparePartCost    uint      `gorm:"type:int(11);null" json:"spare_part_cost"`
	LaborFee         uint      `gorm:"type:int(11);null" json:"labor_fee"`
	AdditionalNotes  string    `gorm:"type:varchar(255);null" json:"additional_notes"`
	CreatedBy        *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	CreatedAt        time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedAt        time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy   *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	Workshop           *UserProfile    `gorm:"foreignKey:WorkshopNo;references:UserProfileNo" json:"workshop,omitempty"`
	Insurer            *UserProfile    `gorm:"foreignKey:InsurerNo;references:UserProfileNo" json:"insurer,omitempty"`
	WorkshopPanels     *WorkshopPanels `gorm:"foreignKey:WorkshopPanelNo;references:WorkshopPanelNo" json:"workshop_panel,omitempty"`
	Mou                *MOU            `gorm:"foreignKey:MouNo;references:MouNo" json:"mou,omitempty"`
	CreatedByUser      *User           `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User           `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
	Measurements       []Measurement   `gorm:"foreignKey:PanelPricingNo;references:PanelPricingNo;" json:"measurements,omitempty"`
}

func (PanelPricing) TableName() string {
	return "r_panel_pricing"
}
