package panels

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	panels := router.Group("/panels")
	{
		panels.GET("", handler.GetAllPanels)
	}
}
