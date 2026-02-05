package suppliers

type BaseSupplierRequest struct {
	WorkshopNo      uint   `json:"workshop_no"`
	SupplierName    string `json:"supplier_name"`
	SupplierAddress string `json:"supplier_address"`
	SupplierPhone   string `json:"supplier_phone"`
	SupplierEmail   string `json:"supplier_email"`
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
