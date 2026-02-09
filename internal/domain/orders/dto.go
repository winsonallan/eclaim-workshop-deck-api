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
