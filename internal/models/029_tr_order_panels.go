package models

import "time"

type OrderPanel struct {
	OrderPanelNo            uint      `gorm:"primaryKey;type:int(11);not null;autoIncrement" json:"order_panel_no"`
	WorkOrderNo             uint      `gorm:"type:int(11);not null" json:"work_order_no"`
	WorkOrderGroupNumber    uint      `gorm:"type:tinyint(3);not null;default:0" json:"work_order_group_no"`
	InitialProposer         string    `gorm:"type:enum('insurer','workshop');not null;default:'insurer'" json:"initial_proposer"`
	CurrentRound            uint      `gorm:"type:tinyint(3);null" json:"current_round"`
	InsurancePanelPricingNo *uint     `gorm:"type:int(11);null" json:"insurance_panel_pricing_no"`
	InsurancePanelName      string    `gorm:"type:varchar(255);null" json:"insurance_panel_name"`
	InsurerPrice            uint      `gorm:"type:int(11);null" json:"insurer_price"`
	InsurerMeasurementNo    *uint     `gorm:"type:int(11);null" json:"insurer_measurement_no"`
	InsurerServiceType      string    `gorm:"type:enum('repair','replacement');null" json:"insurer_service_type"`
	InsurerQty              uint      `gorm:"type:int(11);null" json:"insurer_qty"`
	WorkshopPanelPricingNo  *uint     `gorm:"type:int(11);null" json:"workshop_panel_pricing_no"`
	WorkshopPanelName       string    `gorm:"type:varchar(255);null" json:"workshop_panel_name"`
	WorkshopPrice           *uint     `gorm:"type:int(11);null" json:"workshop_price"`
	WorkshopMeasurementNo   *uint     `gorm:"type:int(11);null" json:"workshop_measurement_no"`
	WorkshopServiceType     string    `gorm:"type:enum('repair','replacement');null" json:"workshop_service_type"`
	WorkshopQty             *uint     `gorm:"type:int(11);null" json:"workshop_qty"`
	FinalPanelPricingNo     *uint     `gorm:"type:int(11);null" json:"final_panel_pricing_no"`
	FinalPanelName          string    `gorm:"type:varchar(255);null" json:"final_panel_name"`
	FinalPrice              *uint     `gorm:"type:int(11);null" json:"final_price"`
	FinalMeasurementNo      *uint     `gorm:"type:int(11);null" json:"final_measurement_no"`
	FinalServiceType        string    `gorm:"type:enum('repair','replacement')" json:"final_service_type"`
	FinalQty                *uint     `gorm:"type:int(11);null" json:"final_qty"`
	IsIncluded              bool      `gorm:"type:tinyint(1);default:1;not null" json:"is_included"`
	IsSpecialRepair         bool      `gorm:"type:tinyint(1);default:0;not null" json:"is_special_repair"`
	NegotiationStatus       string    `gorm:"type:enum('pending_workshop','negotiating','proposed_additional','additional_work','accepted','rejected');default:pending_workshop;not null" json:"negotiation_status"`
	CompletionStatus        string    `gorm:"type:enum('incomplete','pending_sparepart','completed');null" json:"completion_status"`
	IsLocked                bool      `gorm:"type:tinyint(1);default:0;not null" json:"is_locked"`
	CreatedAt               time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy               *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	UpdatedAt               time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy          *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	WorkOrder            *WorkOrder    `gorm:"foreignKey:WorkOrderNo;references:WorkOrderNo" json:"work_order,omitempty"`
	InsurerPanelPricing  *PanelPricing `gorm:"foreignKey:InsurancePanelPricingNo;references:PanelPricingNo" json:"insurer_panel_pricing,omitempty"`
	WorkshopPanelPricing *PanelPricing `gorm:"foreignKey:WorkshopPanelPricingNo;references:PanelPricingNo" json:"workshop_panel_pricing,omitempty"`
	FinalPanelPricing    *PanelPricing `gorm:"foreignKey:FinalPanelPricingNo;references:PanelPricingNo" json:"final_panel_pricing,omitempty"`
	InsurerMeasurement   *Measurement  `gorm:"foreignKey:InsurerMeasurementNo;references:MeasurementNo" json:"insurer_measurement,omitempty"`
	WorkshopMeasurement  *Measurement  `gorm:"foreignKey:WorkshopMeasurementNo;references:MeasurementNo" json:"workshop_measurement,omitempty"`
	FinalMeasurement     *Measurement  `gorm:"foreignKey:FinalMeasurementNo;references:MeasurementNo" json:"final_measurement,omitempty"`
	CreatedByUser        *User         `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser   *User         `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`

	RepairHistory      []RepairHistory      `gorm:"foreignKey:OrderPanelNo;references:OrderPanelNo" json:"repair_history,omitempty"`
	NegotiationHistory []NegotiationHistory `gorm:"foreignKey:OrderPanelNo;references:OrderPanelNo" json:"negotiation_history,omitempty"`
}

func (OrderPanel) TableName() string {
	return "tr_order_panels"
}
