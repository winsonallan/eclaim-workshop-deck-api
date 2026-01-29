package auth

import (
	"crypto/rand"
	"eclaim-workshop-deck-api/internal/common/response"
	"eclaim-workshop-deck-api/internal/domain/settings"
	"eclaim-workshop-deck-api/internal/models"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

type GenerateAPIKeyRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ExpiresIn   int    `json:"expires_in_days"`
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Register - updated to return both tokens
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, accessToken, refreshToken, err := h.service.Register(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    900, // 15 minutes in seconds
	})
}

// Login - updated to return both tokens
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, accessToken, refreshToken, err := h.service.Login(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	userProfileNo := user.UserProfileNo

	if userProfileNo != nil {
		settingsRepo := settings.NewRepository(h.service.repo.db)
		settingsService := settings.NewService(settingsRepo)

		workshopDetails, err := settingsService.GetWorkshopDetailsFromUserProfileNo(*userProfileNo)
		if err != nil {
			c.JSON(http.StatusOK, AuthResponse{
				User:         user,
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				TokenType:    "Bearer",
				ExpiresIn:    900,
			})
		} else {
			c.JSON(http.StatusOK, AuthResponse{
				User:            user,
				WorkshopDetails: workshopDetails,
				AccessToken:     accessToken,
				RefreshToken:    refreshToken,
				TokenType:       "Bearer",
				ExpiresIn:       900,
			})
		}
	} else {
		c.JSON(http.StatusOK, AuthResponse{
			User:         user,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			TokenType:    "Bearer",
			ExpiresIn:    900, // 15 minutes in seconds
		})
	}

}

// NEW: RefreshToken handler
func (h *Handler) RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.service.RefreshToken(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
		"token_type":    "Bearer",
		"expires_in":    900, // 15 minutes in seconds
	})
}

func (h *Handler) GetUserByEmail(c *gin.Context) {
	var req FindByEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.GetUserByEmail(req)
	if err != nil {
		response.Error(c, http.StatusNotFound, "User with that email is not found")
		return
	}

	response.Success(c, http.StatusOK, "User Found", user)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.ChangePassword(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
	} else {
		response.Success(c, http.StatusOK, "Password Changed Successfully", user)
	}
}

func (h *Handler) UpdateAccount(c *gin.Context) {
	var req UpdateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.UpdateAccount(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
	} else {
		response.Success(c, http.StatusOK, "Account Changed Successfully", user)
	}
}

func (h *Handler) GenerateAPIKey(c *gin.Context) {
	var req GenerateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	apiKey, err := generateRandomKey(32)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate API key")
		return
	}

	var expiresAt *time.Time
	if req.ExpiresIn > 0 {
		expiry := time.Now().AddDate(0, 0, req.ExpiresIn)
		expiresAt = &expiry
	}

	key := &models.APIKey{
		Key:         apiKey,
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
		ExpiresAt:   expiresAt,
	}

	if err := h.service.repo.db.Create(key).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create API key")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "API key created successfully. Store this key securely - it won't be shown again!",
		"data": gin.H{
			"api_key":     apiKey,
			"name":        req.Name,
			"description": req.Description,
			"expires_at":  expiresAt,
		},
	})
}

func (h *Handler) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.ResetPassword(req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "A new password has been sent to your email.", nil)
}

func generateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
