package orders

import (
	"eclaim-workshop-deck-api/internal/common/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNegotiatingOrders(c *gin.Context) {
	woIDStr := c.Query("workshop_no")

	if woIDStr == "" {
		response.Error(c, http.StatusBadRequest, "workshop no is needed")
		return
	}

	woID, err := strconv.ParseUint(woIDStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid workshop no format")
		return
	}

	orders, err := h.service.GetNegotiatingOrders(uint(woID))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Orders Retrieved Successfully", gin.H{"orders": orders})
}

func (h *Handler) CancelNegotiation(c *gin.Context) {
	var req CancelNegotiationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.service.CancelNegotiation(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Negotiation Cancelled Successfully", gin.H{"order": order})
}

func (h *Handler) ForwardAdditionalProposal(c *gin.Context) {
	var req ApproveAdditionalProposalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	workOrder, err := h.service.ForwardAdditionalProposal(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Additional Work Proposal Approved Successfully", gin.H{"work_order": workOrder})
}
