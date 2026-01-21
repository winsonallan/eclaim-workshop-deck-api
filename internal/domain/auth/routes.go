package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	auth := router.Group("/auth")
	{
		// Public routes
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		
		auth.GET("/get-user-by-email", handler.GetUserByEmail)
		auth.POST("/change-password", handler.ChangePassword)

		// Admin only - Generate API keys (you can add admin check later)
		auth.POST("/generate-api-key", handler.GenerateAPIKey)
	}
}