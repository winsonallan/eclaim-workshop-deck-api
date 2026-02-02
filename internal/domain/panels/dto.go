package panels

type CreateMOURequest struct {
	MouDocumentNumber string `json:"mou_document_number" binding:"required"`
	MouExpiryDate     string `json:"mou_expiry_date"`
	InsurerNo         uint   `json:"insurer_no" binding:"required"`
	WorkshopNo        uint   `json:"workshop_no" binding:"required"`
	CreatedBy         uint   `json:"created_by"`
}

type CreatePanelPricingRequest struct {
	WorkshopNo       uint   `json:"workshop_no" binding:"required"`
	InsurerNo        uint   `json:"insurer_no"`
	MouNo            uint   `json:"mou_no"`
	WorkshopPanelNo  uint   `json:"workshop_panel_no"`
	ServiceType      string `json:"service_type" binding:"required"`
	IsFixedPrice     bool   `json:"is_fixed_price" binding:"required"`
	VehicleRangeLow  uint   `json:"vehicle_range_low"`
	VehicleRangeHigh uint   `json:"vehicle_range_high"`
	SparePartCost    uint   `json:"spare_part_cost"`
	LaborFee         uint   `json:"labor_fee"`
	CreatedBy        uint   `json:"created_by"`
}
