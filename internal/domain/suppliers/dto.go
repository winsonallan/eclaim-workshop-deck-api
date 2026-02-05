package suppliers

type BaseSupplierRequest struct {
	WorkshopNo           uint   `json:"workshop_no"`
	SupplierName         string `json:"supplier_name"`
	SupplierAddress      string `json:"supplier_address"`
	SupplierCityNo       uint   `json:"supplier_city_no"`
	SupplierCityType     string `json:"supplier_city_type"`
	SupplierCityName     string `json:"supplier_city_name"`
	SupplierProvinceNo   uint   `json:"supplier_province_no"`
	SupplierProvinceName string `json:"supplier_province_name"`
	SupplierPhone        string `json:"supplier_phone"`
	SupplierEmail        string `json:"supplier_email"`
}

type AddSupplierRequest struct {
	BaseSupplierRequest
	CreatedBy uint `json:"created_by" binding:"required"`
}

type UpdateSupplierRequest struct {
	BaseSupplierRequest
	LastModifiedBy uint `json:"last_modified_by" binding:"required"`
}

type DeleteSupplierRequest struct {
	LastModifiedBy uint `json:"last_modified_by" binding:"required"`
}
