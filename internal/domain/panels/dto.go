package panels

import "eclaim-workshop-deck-api/internal/models"

type RegisterRequest struct {
	UserProfileNo uint   `json:"user_profile_no"`
	RoleNo        uint   `json:"role_no" binding:"required"`
	Name          string `json:"user_name" binding:"required"`
	UserId        string `json:"user_id" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type FindByEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ChangePasswordRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Username        string `json:"user_id" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}
