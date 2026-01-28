package settings

type CreateWorkshopDetailsRequest struct {
	ProfileNo    uint   `json:"user_profile_no" binding:"required"`
	Capacity     uint   `json:"capacity"`
	Description  string `json:"description"`
	IsAuthorized bool   `json:"is_authorized" binding:"required"`
	Specialist   string `json:"specialist"`
	CreatedBy    uint   `json:"created_by"`
}

type UpdateWorkshopDetailsRequest struct {
	WorkshopDetailsNo uint   `json:"workshop_details_no" binding:"required"`
	WorkshopName      string `json:"workshop_name" binding:"required"`
	Capacity          uint   `json:"capacity"`
	Address           string `json:"address" binding:"required"`
	CityNo            uint   `json:"city_no"`
	CityType          string `json:"city_type"`
	CityName          string `json:"city_name" binding:"required"`
	Phone             string `json:"phone" binding:"required"`
	Email             string `json:"email" binding:"required"`
	Description       string `json:"description"`
	LastModifiedBy    uint   `json:"last_modified_by"`
}

type UpdateWorkshopPICRequest struct {
	WorkshopPicNo    uint   `json:"workshop_pic_no" binding:"required"`
	WorkshopPicName  string `json:"workshop_pic_name" binding:"required"`
	WorkshopPicTitle string `json:"workshop_pic_title" binding:"required"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	LastModifiedBy   uint   `json:"last_modified_by"`
}

type DeleteWorkshopPICRequest struct {
	LastModifiedBy uint `json:"last_modified_by"`
}

type CreateWorkshopPICRequest struct {
	WorkshopDetailsNo uint   `json:"workshop_details_no" binding:"required"`
	PicName           string `json:"workshop_pic_name" binding:"required"`
	PicTitle          string `json:"workshop_pic_title" binding:"required"`
	Phone             string `json:"phone" binding:"required"`
	Email             string `json:"email" binding:"required"`
	CreatedBy         uint   `json:"created_by"`
}
