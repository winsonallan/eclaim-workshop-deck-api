package posts

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	posts := router.Group("/posts")

	posts.Use(authMiddleware)
	{
		posts.POST("", handler.CreatePost)
		posts.GET("", handler.GetAllPosts)
		posts.GET("/my", handler.GetMyPosts)
		posts.GET("/:id", handler.GetPost)
		posts.PUT("/:id", handler.UpdatePost)
		posts.DELETE("/:id", handler.DeletePost)
	}
}
