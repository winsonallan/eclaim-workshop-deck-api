package panels

type CreateMOURequest struct {
	MouDocumentNumber string `json:"mou_document_number" binding:"required"`
	MouExpiryDate     string `json:"mou_expiry_date"`
	InsurerNo         uint   `json:"insurer_no" binding:"required"`
	WorkshopNo        uint   `json:"workshop_no" binding:"required"`
	CreatedBy         uint   `json:"created_by"`
}

type MeasurementInput struct {
	ConditionText string `json:"condition_text" binding:"required"`
	LaborFee      uint   `json:"labor_fee" binding:"required"`
	Note          string `json:"notes"`
}

type BasePricingRequest struct {
	WorkshopNo uint `json:"workshop_no" binding:"required"`
	InsurerNo  uint `json:"insurer_no"`
	MouNo      uint `json:"mou_no"`

	WorkshopPanelNo uint   `json:"workshop_panel_no"`
	PanelNo         uint   `json:"panel_no"`
	PanelName       string `json:"panel_name"`
	ServiceType     string `json:"service_type" binding:"required"`

	IsCustom        *bool  `json:"is_custom" binding:"required"`
	CustomPanelName string `json:"custom_panel_name"`

	VehicleRangeLow  uint `json:"vehicle_range_low"`
	VehicleRangeHigh uint `json:"vehicle_range_high"`

	IsFixedPrice   *bool  `json:"is_fixed_price" binding:"required"`
	SparePartCost  uint   `json:"spare_part_cost"`
	LaborFee       uint   `json:"labor_fee"`
	AdditionalNote string `json:"additional_notes"`

	Measurements []MeasurementInput `json:"measurements"`
}
type CreatePanelPricingRequest struct {
	BasePricingRequest
	CreatedBy uint `json:"created_by" binding:"required"`
}

type UpdatePanelPricingRequest struct {
	BasePricingRequest
	LastModifiedBy uint `json:"last_modified_by" binding:"required"`
}

type CreateWorkshopPanelRequest struct {
	WorkshopNo uint   `json:"workshop_no" binding:"required"`
	PanelNo    uint   `json:"panel_no"`
	PanelName  string `json:"panel_name" binding:"required"`
	CreatedBy  uint   `json:"created_by" binding:"required"`
}

type PricingRequest interface {
	GetServiceType() string
	GetLaborFee() uint
	GetSparePartCost() uint
}

type DeletePanelPricingRequest struct {
	LastModifiedBy uint `json:"last_modified_by"`
}
