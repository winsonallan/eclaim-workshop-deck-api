package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	auth := router.Group("/auth")
	{
		// Public routes
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)

		// Protected routes
		auth.GET("/me", authMiddleware, handler.GetMe)
		
		// Admin only - Generate API keys (you can add admin check later)
		auth.POST("/generate-api-key", authMiddleware, handler.GenerateAPIKey)
	}
}