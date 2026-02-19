package location

import (
	"eclaim-workshop-deck-api/internal/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service *Service
	log     *zap.Logger
}

type GenerateAPIKeyRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ExpiresIn   int    `json:"expires_in_days"`
}

func NewHandler(service *Service, log *zap.Logger) *Handler {
	return &Handler{service: service, log: log}
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

func (h *Handler) GetProvinces(c *gin.Context) {
	provinces, err := h.service.GetProvinces()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch provinces")
		return
	}

	response.Success(c, http.StatusOK, "provinces retrieved successfully", gin.H{
		"provinces": provinces,
		"count":     len(provinces),
	})
}
