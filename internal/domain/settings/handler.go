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

func (h *Handler) GetWorkshopPICs(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	workshopPICs, err := h.service.GetWorkshopPICs(uint(id))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Workshop PICs Retrieved Successfully", gin.H{"workshop_pics": workshopPICs})
}

func (h *Handler) CreateWorkshopDetails(c *gin.Context) {
	var req CreateWorkshopDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	workshopDetails, err := h.service.CreateWorkshopDetails(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "workshop details created successfully", gin.H{"workshopDetails": workshopDetails})
}

func (h *Handler) CreateWorkshopPIC(c *gin.Context) {
	var req CreateWorkshopPICRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	workshopPIC, err := h.service.CreateWorkshopPIC(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "workshop PIC created successfully", gin.H{"workshop_pics": workshopPIC})
}

func (h *Handler) UpdateWorkshopDetails(c *gin.Context) {
	var req UpdateWorkshopDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.UpdateWorkshopDetails(req)

	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
	} else {
		response.Success(c, http.StatusOK, "Workshop Details Changed Successfully", user)
	}
}
