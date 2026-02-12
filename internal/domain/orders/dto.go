package orders

import "time"

type AddClientRequest struct {
	ClientName          string `json:"client_name" binding:"required"`
	ClientEmail         string `json:"client_email"`
	ClientPhone         string `json:"client_phone" binding:"required"`
	CityNo              uint   `json:"city_no" binding:"required"`
	CityType            string `json:"city_type" binding:"required"`
	CityName            string `json:"city_name" binding:"required"`
	Address             string `json:"address" binding:"required"`
	VehicleBrandName    string `json:"vehicle_brand_name" binding:"required"`
	VehicleSeriesName   string `json:"vehicle_series_name" binding:"required"`
	VehicleChassisNo    string `json:"vehicle_chassis_no" binding:"required"`
	VehicleLicensePlate string `json:"vehicle_license_plate" binding:"required"`
	VehiclePrice        uint   `json:"vehicle_price" binding:"required"`
}

type CreateOrderRequest struct {
	WorkshopNo    uint              `json:"workshop_no" binding:"required"`
	InsuranceNo   uint              `json:"insurance_no"`
	ClientNo      uint              `json:"client_no"`
	ClientDetails *AddClientRequest `json:"client_details"`
	ClaimDetails  string            `json:"claim_details" binding:"required"`
	ETA           time.Time         `json:"eta"`
	Status        string            `json:"status" binding:"required"`
	CreatedBy     uint              `json:"created_by" binding:"required"`
}

type OrderPanelRequest struct {
	WOGroupNumber           uint   `json:"wo_group_number"`
	InsurancePanelPricingNo uint   `json:"insurance_panel_pricing_no"`
	InsurancePanelName      string `json:"insurance_panel_name"`
	InsurerPrice            uint   `json:"insurer_price"`
	InsurerMeasurementNo    uint   `json:"insurer_measurement_no"`
	InsurerServiceType      string `json:"insurer_service_type"`
	InsurerQty              uint   `json:"insurer_qty"`
	WorkshopPanelPricingNo  uint   `json:"workshop_panel_pricing_no"`
	WorkshopPanelName       string `json:"workshop_panel_name"`
	WorkshopPrice           uint   `json:"workshop_price"`
	WorkshopMeasurementNo   uint   `json:"workshop_measurement_no"`
	WorkshopServiceType     string `json:"workshop_service_type"`
	WorkshopQty             uint   `json:"workshop_qty"`
	IsIncluded              bool   `json:"is_included"`
	IsSpecialRepair         bool   `json:"is_special_repair"`
}

type CreateWorkOrderRequest struct {
	OrderNo                  uint                `json:"order_no" binding:"required"`
	AdditionalWorkOrderCount uint                `json:"add_wo_count"`
	WorkOrderDocumentNumber  string              `json:"wo_doc_number"`
	WorkOrderUrl             string              `json:"wo_url"`
	OrderPanels              []OrderPanelRequest `json:"order_panels"`

	CreatedBy uint `json:"created_by"`
}

type ProposeAdditionalWorkRequest struct {
	WorkOrderNo uint                `json:"work_order_no" binding:"required"`
	OrderPanels []OrderPanelRequest `json:"order_panels"`

	LastModifiedBy uint `json:"last_modified_by"`
}
