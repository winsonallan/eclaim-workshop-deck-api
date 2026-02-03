package panels

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	panels := router.Group("/panels")
	{
		panels.Use(authMiddleware)
		{
			panels.GET("", handler.GetAllPanels)
			panels.GET("/workshop/:id", handler.GetAllWorkshopPanels)

			panelsPricing := panels.Group("/pricing")
			{
				panelsPricing.GET("", handler.GetPanelPricings)
				panelsPricing.GET("/workshop/:id", handler.GetAllWorkshopPanelPricings)
				panelsPricing.POST("", handler.CreatePanelPricing)
			}
		}
	}

	mou := router.Group("/mou")
	{
		mou.Use(authMiddleware)
		{
			mou.GET("", handler.GetMOUs)
			mou.POST("", handler.CreateMOU)
		}
	}
}
