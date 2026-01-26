package settings

import (
	"eclaim-workshop-deck-api/internal/common/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAccount(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	account, err := h.service.GetAccount(uint(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Account Retrieved Successfully", gin.H{"account": account})
}

func (h *Handler) GetProfileDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	userProfile, err := h.service.GetProfileDetails(uint(id))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User Profile Retrieved Successfully", gin.H{"user_profile": userProfile})
}

func (h *Handler) GetWorkshopDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	workshopDetails, err := h.service.GetWorkshopDetails(uint(id))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Workshop Details Retrieved Successfully", gin.H{"workshop_details": workshopDetails})
}
