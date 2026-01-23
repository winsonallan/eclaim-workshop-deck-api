package auth

import (
	"crypto/rand"
	"eclaim-workshop-deck-api/internal/common/response"
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
	ExpiresIn   int    `json:"expires_in_days"` // Optional: days until expiration (0 = never)
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, token, err := h.service.Register(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, AuthResponse{
		User:  user,
		Token: token,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, token, err := h.service.Login(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		User:  user,
		Token: token,
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

func (h *Handler) GenerateAPIKey(c *gin.Context) {
	var req GenerateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Generate random API key
	apiKey, err := generateRandomKey(32) // 32 bytes = 64 hex characters
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

// Generate cryptographically secure random key
func generateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
