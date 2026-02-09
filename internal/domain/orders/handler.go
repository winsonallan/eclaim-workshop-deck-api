package orders

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

// Read
func (h *Handler) GetOrders(c *gin.Context) {
	orders, err := h.service.GetOrders()

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Orders Retrieved Successfully", gin.H{"orders": orders})
}

// Create
func (h *Handler) AddClient(c *gin.Context) {
	var req AddClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	client, err := h.service.AddClient(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "client created successfully", gin.H{"clients": client})
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.service.CreateOrder(req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "order created successfully", gin.H{"orders": order})
}
