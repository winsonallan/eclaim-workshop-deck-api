package panels

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

func (h *Handler) GetAllPanels(c *gin.Context) {
	panels, err := h.service.GetAllPanels()

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Panels Retrieved Successfully", gin.H{"panels": panels})
}
