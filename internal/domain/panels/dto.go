package panels

type CreateMOURequest struct {
	MouDocumentNumber string `json:"mou_document_number" binding:"required"`
	MouExpiryDate     string `json:"mou_expiry_date"`
	InsurerNo         uint   `json:"insurer_no" binding:"required"`
	WorkshopNo        uint   `json:"workshop_no" binding:"required"`
	CreatedBy         uint   `json:"created_by"`
}
