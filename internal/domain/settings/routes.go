package settings

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	user := router.Group("/user")
	{
		user.Use(authMiddleware)
		{
			user.GET("/:id", handler.GetAccount)
		}
	}

	userProfile := router.Group("/user-profile")
	{
		userProfile.Use(authMiddleware)
		{
			userProfile.GET("/:id", handler.GetProfileDetails)
		}
	}

	workshopDetails := router.Group("/workshop")
	{
		workshopDetails.Use(authMiddleware)
		{
			workshopDetails.GET("/:id", handler.GetWorkshopDetails)
			workshopDetails.POST("", handler.CreateWorkshopDetails)

			workshopDetails.GET("/pic/:id", handler.GetWorkshopPICs)
			workshopDetails.POST("/pic", handler.CreateWorkshopPIC)

			workshopDetails.PUT("/:id", handler.UpdateWorkshopDetails)
		}
	}

}
