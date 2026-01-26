package admin

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {

	userProfile := router.Group("/user-profile")
	{
		userProfile.POST("", authMiddleware, handler.CreateUserProfile)
	}
}
