package main

import (
	"eclaim-workshop-deck-api/internal/config"
	"eclaim-workshop-deck-api/internal/handlers"
	"eclaim-workshop-deck-api/internal/middleware"
	"eclaim-workshop-deck-api/internal/models"
	"eclaim-workshop-deck-api/internal/repository"
	"eclaim-workshop-deck-api/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to database
	db := config.ConnectDB(cfg)

	// Auto migrate
	if err := db.AutoMigrate(&models.User{}, &models.SamplePost{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	samplePostRepo := repository.NewPostRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	postService := services.NewPostService(samplePostRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userRepo)
	postHandler := handlers.NewPostHandler(postService)

	// Setup Gin
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORSMiddleware(cfg.FrontendURL))

	// Public routes
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			protected.GET("/auth/me", userHandler.GetMe)
			// Add more protected routes here
			posts := protected.Group("/posts")
			{
				posts.POST("", postHandler.CreatePost)           // CREATE
				posts.GET("", postHandler.GetAllPosts)            // READ ALL
				posts.GET("/my", postHandler.GetMyPosts)          // READ MY POSTS
				posts.GET("/:id", postHandler.GetPost)            // READ ONE
				posts.PUT("/:id", postHandler.UpdatePost)         // UPDATE
				posts.DELETE("/:id", postHandler.DeletePost)      // DELETE
			}
		}
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}