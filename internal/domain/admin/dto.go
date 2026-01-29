package admin

type CreateUserProfileRequest struct {
	Type      string `json:"user_profile_type" binding:"required"`
	Name      string `json:"user_profile_name" binding:"required"`
	CityNo    uint   `json:"user_profile_city_no" binding:"required"`
	CityType  string `json:"user_profile_city_type" binding:"required"`
	CityName  string `json:"user_profile_city_name" binding:"required"`
	Address   string `json:"user_profile_address" binding:"required"`
	Email     string `json:"user_profile_email" binding:"required"`
	Phone     string `json:"user_profile_phone" binding:"required"`
	CreatedBy uint   `json:"created_by" binding:"required"`

	Capacity     uint   `json:"capacity"`
	Description  string `json:"description"`
	IsAuthorized bool   `json:"is_authorized"`
	Specialist   string `json:"specialist"`
}
