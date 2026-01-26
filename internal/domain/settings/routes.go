package settings

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	user := router.Group("/user")
	{
		user.GET("/:id", authMiddleware, handler.GetAccount)
	}

	userProfile := router.Group("/user-profile")
	{
		userProfile.GET("/:id", authMiddleware, handler.GetProfileDetails)
	}

	workshopDetails := router.Group("/workshop")
	{
		workshopDetails.GET("/:id", authMiddleware, handler.GetWorkshopDetails)
	}
}
