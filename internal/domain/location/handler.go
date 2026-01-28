package location

import (
	"eclaim-workshop-deck-api/internal/common/response"
	"net/http"

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

func (h *Handler) GetCities(c *gin.Context) {
	cities, err := h.service.GetCities()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch cities")
		return
	}

	response.Success(c, http.StatusOK, "Cities retrieved successfully", gin.H{
		"cities": cities,
		"count":  len(cities),
	})
}
