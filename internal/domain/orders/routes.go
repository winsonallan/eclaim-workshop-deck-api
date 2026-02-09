package orders

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	orders := router.Group("/orders")
	{
		orders.Use(authMiddleware)
		{
			orders.GET("", handler.GetOrders)
			orders.POST("", handler.CreateOrder)
		}
	}

	clients := router.Group("/clients")
	{
		clients.Use(authMiddleware)
		{
			clients.POST("", handler.AddClient)
		}
	}
}
