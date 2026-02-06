package usermanagement

type AddUserRequest struct {
	UserProfileNo uint   `json:"user_profile_no"`
	Name          string `json:"name"`
	UserName      string `json:"username"`
	Email         string `json:"email"`
	RoleNo        uint   `json:"role_no"`
	CreatedBy     uint   `json:"created_by"`
}

type ChangeUserRoleRequest struct {
	RoleNo         uint `json:"role_no" binding:"required"`
	LastModifiedBy uint `json:"last_modified_by" binding:"required"`
}

type DeleteUserRequest struct {
	LastModifiedBy uint `json:"last_modified_by"`
}
