package suppliers

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	suppliers := router.Group("/suppliers")
	{
		suppliers.Use(authMiddleware)
		{
			suppliers.GET("", handler.GetSuppliers)
			suppliers.PUT("/:id", handler.UpdateSupplier)
			suppliers.DELETE("/:id", handler.DeleteSupplier)

			suppliers.GET("/workshop/:id", handler.GetWorkshopSuppliers)
			suppliers.POST("/workshop/:id", handler.AddSupplier)
		}
	}

}
