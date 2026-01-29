package admin

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {

	admin := router.Group("/admin")
	{
		admin.Use(authMiddleware)
		{
			userProfile := admin.Group("/user-profile")
			{
				userProfile.POST("", handler.CreateUserProfile)
			}
		}

	}
}
