package location

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	location := router.Group("/location")
	{
		// Public routes
		location.GET("/cities", handler.GetCities)
		location.GET("/provinces", handler.GetProvinces)
	}
}
