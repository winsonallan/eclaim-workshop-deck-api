package models

import "time"

type NegotiationHistory struct {
	NegotiationHistoryNo   uint      `gorm:"type:int(11);primaryKey;autoIncrement;not null" json:"negotiation_history_no"`
	OrderPanelNo           uint      `gorm:"type:int(11);not null" json:"order_panel_no"`
	RoundCount             uint      `gorm:"type:tinyint(3);default:1;not null" json:"round_count"`
	OldPanelPricingNo      *uint     `gorm:"type:int(11);null" json:"old_panel_pricing_no"`
	OldPrice               uint      `gorm:"type:int(11);null" json:"old_price"`
	OldMeasurementNo       *uint     `gorm:"type:int(11);null" json:"old_measurement_no"`
	OldServiceType         string    `gorm:"type:enum('repair','replacement');null" json:"old_service_type"`
	OldQty                 uint      `gorm:"type:int(11);null" json:"old_qty"`
	ProposedPanelPricingNo uint      `gorm:"type:int(11);not null" json:"proposed_panel_pricing_no"`
	ProposedPrice          uint      `gorm:"type:int(11);not null" json:"proposed_price"`
	ProposedMeasurementNo  uint      `gorm:"type:int(11);not null" json:"proposed_measurement_no"`
	ProposedServiceType    string    `gorm:"enum('repair','replacement');not null" json:"proposed_service_type"`
	ProposedQty            uint      `gorm:"type:int(11);null" json:"proposed_qty"`
	WorkshopNotes          string    `gorm:"type:varchar(255);null" json:"workshop_notes"`
	InsuranceDecision      string    `gorm:"type:enum('pending','accepted','declined');not null" json:"insurance_decision"`
	InsuranceNotes         string    `gorm:"type:varchar(255);null" json:"insurance_notes"`
	DecidedBy              uint      `gorm:"type:int(11);null" json:"decided_by"`
	CompletedDate          time.Time `gorm:"type:datetime;null" json:"completed_date"`
	IsLocked               bool      `gorm:"type:tinyint(1);default:0;not null" json:"is_locked"`
	CreatedAt              time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy              *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	UpdatedAt              time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy         *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	DecidedByUser        *User         `gorm:"foreignKey:DecidedBy;references:UserNo" json:"decided_by_user,omitempty"`
	OrderPanel           *OrderPanel   `gorm:"foreignKey:OrderPanelNo;references:OrderPanelNo" json:"order_panel,omitempty"`
	OldPanelPricing      *PanelPricing `gorm:"foreignKey:OldPanelPricingNo;references:PanelPricingNo" json:"old_panel_pricing,omitempty"`
	ProposedPanelPricing *PanelPricing `gorm:"foreignKey:ProposedPanelPricingNo;references:PanelPricingNo" json:"proposed_panel_pricing,omitempty"`
	OldMeasurement       *Measurement  `gorm:"foreignKey:OldMeasurementNo;references:MeasurementNo" json:"old_measurement,omitempty"`
	ProposedMeasurement  *Measurement  `gorm:"foreignKey:ProposedMeasurementNo;references:MeasurementNo" json:"proposed_measurement,omitempty"`
	CreatedByUser        *User         `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser   *User         `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (NegotiationHistory) TableName() string {
	return "tr_negotiation_history"
}
