package orders

import (
	"eclaim-workshop-deck-api/internal/common/response"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) GetIncomingOrders(c *gin.Context) {
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

	orders, err := h.service.GetIncomingOrders(uint(woID))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Orders Retrieved Successfully", gin.H{"orders": orders})
}

func (h *Handler) SubmitNegotiation(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to parse multipart form")
		return
	}

	dataStr := c.PostForm("data")
	if dataStr == "" {
		response.Error(c, http.StatusBadRequest, "Missing 'data' field in form")
		return
	}

	var req SubmitNegotiationRequest
	if err := json.Unmarshal([]byte(dataStr), &req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid JSON in 'data' field")
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to get multipart form")
		return
	}

	files := form.File["files"]

	uploadFn := func(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
		return h.storage.Upload(file, header, folder)
	}

	result, err := h.service.SubmitNegotiation(&req, files, uploadFn)
	if err != nil {
		h.log.Error("Failed to submit negotiation", zap.Error(err))

		// You can return a mapped error message here
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Negotiation successfully submitted", gin.H{
		"work_order": result,
	})
}

// Update
func (h *Handler) AcceptOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid order no")
		return
	}

	var req AcceptDeclineOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.service.AcceptOrder(uint(id), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Order successfully accepted", gin.H{"order": order})
}

func (h *Handler) DeclineOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid order no")
		return
	}

	var req AcceptDeclineOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.service.DeclineOrder(uint(id), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Order successfully accepted", gin.H{"order": order})
}
