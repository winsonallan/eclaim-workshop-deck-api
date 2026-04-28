package orders

import (
	"eclaim-workshop-deck-api/internal/common/response"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetRepairedOrders retrieves repaired orders for a given workshop number.
func (h *Handler) GetRepairedOrders(c *gin.Context) {
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

	orders, err := h.service.GetRepairedOrders(uint(woID))

	for _, o := range orders {
		AttachFullPhotoURLs(&o, os.Getenv("BASE_URL"))
	}

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Orders Retrieved Successfully", gin.H{"orders": orders})
}

// SetRepairedAsUnfinished marks a repaired vehicle as unfinished, allowing for further repairs or adjustments.
func (h *Handler) SetRepairedAsUnfinished(c *gin.Context) {
	var req CancelNegotiationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.service.SetRepairedAsUnfinished(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Set Order as Unfinished Successfully", gin.H{"order": order})
}

// RemindPickup sends a reminder for pickup to the customer for an vehicle that has been repaired and is ready for pickup.
func (h *Handler) RemindPickup(c *gin.Context) {
	var req RemindPickupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	pickupReminders, err := h.service.RemindPickup(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Remind Pickup done successfully", gin.H{"pickupReminders": pickupReminders})
}

// SetAsDelivered marks a repaired order as delivered, indicating that the customer has picked up the vehicle.
func (h *Handler) SetAsDelivered(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		response.Error(c, http.StatusBadRequest, "failed to parse multipart form")
		return
	}

	lastModifiedByStr := c.PostForm("last_modified_by")
	lastModifiedBy, err := strconv.ParseUint(lastModifiedByStr, 10, 32)
	if err != nil || lastModifiedBy == 0 {
		response.Error(c, http.StatusBadRequest, "last_modified_by is required")
		return
	}

	deliveredAtStr := c.PostForm("delivered_at")
	if len(deliveredAtStr) == 0 {
		response.Error(c, http.StatusBadRequest, "delivered_at is required")
	}

	deliveredAt, err := time.Parse("2006-01-02", deliveredAtStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid delivered_at format, expected YYYY-MM-DD")
		return
	}

	invoiceNosStrs := c.PostFormArray("invoice_nos")
	if len(invoiceNosStrs) == 0 {
		response.Error(c, http.StatusBadRequest, "invoice_nos is required")
		return
	}

	uploadFn := func(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
		return h.storage.Upload(file, header, folder)
	}

	var invoiceNos []uint
	for _, s := range invoiceNosStrs {
		n, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid invoice_nos value")
			return
		}
		invoiceNos = append(invoiceNos, uint(n))
	}

	req := SetAsDeliveredRequest{
		InvoiceNos:     invoiceNos,
		DeliveredAt:    deliveredAt,
		LastModifiedBy: uint(lastModifiedBy),
	}

	// Handle optional proof photo
	var proofFile *multipart.FileHeader
	proofFile, _ = c.FormFile("proof_photo") // optional, no error check

	delivery, err := h.service.SetAsDelivered(req, proofFile, uploadFn)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Delivery created successfully", gin.H{"delivery": delivery})
}
