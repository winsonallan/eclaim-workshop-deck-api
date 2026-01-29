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
			userProfile.GET("/workshop/:id", handler.GetWorkshopDetailsFromUserProfileNo)
		}
	}

	workshopDetails := router.Group("/workshop")
	{
		workshopDetails.Use(authMiddleware)
		{
			workshopDetails.GET("/:id", handler.GetWorkshopDetails)
			workshopDetails.POST("", handler.CreateWorkshopDetails)
			workshopDetails.PUT("/:id", handler.UpdateWorkshopDetails)

			workshopDetails.GET("/pic/:id", handler.GetWorkshopPICs)
			workshopDetails.POST("/pic", handler.CreateWorkshopPIC)
			workshopDetails.PUT("/pic/:id", handler.UpdateWorkshopPIC)
			workshopDetails.DELETE("/pic/:id", handler.DeleteWorkshopPIC)
		}
	}

}
