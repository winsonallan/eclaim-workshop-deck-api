package usermanagement

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	usermanagement := router.Group("/user-management")
	{
		usermanagement.Use(authMiddleware)
		{
			usermanagement.GET("/workshop/:id", handler.GetWorkshopUsers)
			usermanagement.POST("/:id", handler.AddUser)
			usermanagement.PUT("/:id", handler.UpdateUserRole)
			usermanagement.DELETE("/:id", handler.DeleteUser)
		}
	}

	roles := router.Group("/roles")
	{
		roles.GET("/:type", handler.GetRoles)
	}

}
