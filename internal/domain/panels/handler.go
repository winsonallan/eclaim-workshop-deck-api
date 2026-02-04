package panels

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

// Read
func (h *Handler) GetAllPanels(c *gin.Context) {
	panels, err := h.service.GetAllPanels()

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Panels Retrieved Successfully", gin.H{"panels": panels})
}

func (h *Handler) GetAllWorkshopPanels(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid workshop panel no")
		return
	}

	panels, err := h.service.GetAllWorkshopPanels(uint(id))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Workshop Panels Retrieved Successfully", gin.H{"workshop_panels": panels})
}

func (h *Handler) GetMOUs(c *gin.Context) {
	// 1. Get query params
	insIDStr := c.Query("insurer_no")
	woIDStr := c.Query("workshop_no")
	mouIDStr := c.Query("mou_no")
	active := c.Query("active") == "true"

	// 2. Validate: Must have at least one ID
	if mouIDStr == "" && insIDStr == "" && woIDStr == "" {
		response.Error(c, http.StatusBadRequest, "Either insurer_no or workshop_no must be provided")
		return
	}

	// 3. Parse strings to uint (handling potential parsing errors)
	var insID, woID, mouID uint64
	var err error

	if insIDStr != "" {
		insID, err = strconv.ParseUint(insIDStr, 10, 32)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid insurer_no format")
			return
		}
	}

	if woIDStr != "" {
		woID, err = strconv.ParseUint(woIDStr, 10, 32)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid workshop_no format")
			return
		}
	}

	if mouIDStr != "" {
		mouID, err = strconv.ParseUint(mouIDStr, 10, 32)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid mou_no format")
			return
		}
	}
	// 4. Call Service
	mous, err := h.service.GetMOUs(uint(insID), uint(woID), uint(mouID), active)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "MOUs Retrieved Successfully", gin.H{"mous": mous})
}

func (h *Handler) GetPanelPricings(c *gin.Context) {
	// 1. Get query params
	insIDStr := c.Query("insurer_no")
	woIDStr := c.Query("workshop_no")
	mouIDStr := c.Query("mou_no")

	// 2. Validate: Must have at least one ID
	if mouIDStr == "" && insIDStr == "" && woIDStr == "" {
		response.Error(c, http.StatusBadRequest, "Either mou_no, insurer_no or workshop_no must be provided")
		return
	}

	// 3. Parse strings to uint (handling potential parsing errors)
	var insID, woID, mouID uint64
	var err error

	if insIDStr != "" {
		insID, err = strconv.ParseUint(insIDStr, 10, 32)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid insurer_no format")
			return
		}
	}

	if woIDStr != "" {
		woID, err = strconv.ParseUint(woIDStr, 10, 32)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid workshop_no format")
			return
		}
	}

	if mouIDStr != "" {
		mouID, err = strconv.ParseUint(mouIDStr, 10, 32)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid mou_no format")
			return
		}
	}
	// 4. Call Service
	panelPricings, err := h.service.GetPanelPricings(uint(insID), uint(woID), uint(mouID))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Panel Pricingss Retrieved Successfully", gin.H{"panel_pricings": panelPricings})
}

func (h *Handler) GetAllWorkshopPanelPricings(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid workshop no")
		return
	}

	panelPricings, err := h.service.GetWorkshopPanelPricings(uint(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Workshop Panel Pricingss Retrieved Successfully", gin.H{"panel_pricings": panelPricings})
}

// Create
func (h *Handler) CreateMOU(c *gin.Context) {
	var req CreateMOURequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	mou, err := h.service.CreateMOU(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "MOU created successfully", gin.H{"mou": mou})
}

func (h *Handler) CreatePanelPricing(c *gin.Context) {
	var req CreatePanelPricingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	panelPricing, err := h.service.CreatePanelPricing(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Panel pricing created successfully", gin.H{"panel_pricing": panelPricing})
}

// Update
func (h *Handler) UpdatePanelPricing(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid panel pricing no")
		return
	}

	var req UpdatePanelPricingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	panelPricing, err := h.service.UpdatePanelPricing(uint(id), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Panel pricing updated successfully", gin.H{"panel_pricing": panelPricing})
}

// Delete
func (h *Handler) DeletePanelPricing(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid panel pricing no")
		return
	}

	var req DeletePanelPricingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	pic, err := h.service.DeletePanelPricing(uint(id), req)

	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
	} else {
		response.Success(c, http.StatusOK, "Panel Pricing Deleted Successfully", pic)
	}
}
