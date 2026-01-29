package admin

import (
	"eclaim-workshop-deck-api/internal/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUserProfile(c *gin.Context) {
	var req CreateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userProfile, err := h.service.CreateUserProfile(req)

	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User Profile created successfully", gin.H{"user_profile": userProfile})
}
