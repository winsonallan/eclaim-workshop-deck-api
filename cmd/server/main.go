package main

import (
	"eclaim-workshop-deck-api/internal/config"
	"eclaim-workshop-deck-api/internal/domain/auth"
	"eclaim-workshop-deck-api/internal/domain/posts"
	"eclaim-workshop-deck-api/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to database
	db := config.ConnectDB(cfg)

	// Auto migrate - Add APIKey model
	if err := db.AutoMigrate(&auth.User{}); err != nil {
		log.Fatal("Failed to migrate User Table:", err)
	}

	if err := db.AutoMigrate(&auth.APIKey{}); err != nil {
		log.Fatal("Failed to migrate APIKey Table:", err)
	}

	if err := db.AutoMigrate(&posts.Post{}); err != nil {
		log.Fatal("Failed to migrate Posts Table:", err)
	}

	// Initialize Auth domain
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)

	// Initialize Posts domain
	postsRepo := posts.NewRepository(db)
	postsService := posts.NewService(postsRepo)
	postsHandler := posts.NewHandler(postsService)

	// Setup Gin
	r := gin.Default()

	// Apply CORS middleware (supports multiple origins)
	r.Use(middleware.CORSMiddleware(cfg.FrontendURLs))

	// API routes
	api := r.Group("/api")
	
	// Apply API Key middleware to ALL API routes
	api.Use(middleware.APIKeyMiddleware(db))
	{
		// Public routes (API Key required, but NO JWT)
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", authHandler.Register)
			authGroup.POST("/login", authHandler.Login)
		}

		// Protected routes (API Key + JWT required)
		authMiddleware := middleware.AuthMiddleware(cfg.JWTSecret)
		protected := api.Group("")
		protected.Use(authMiddleware)
		{
			// Auth protected routes
			protected.GET("/auth/me", authHandler.GetMe)
			
			// Posts routes
			posts.RegisterRoutes(protected, postsHandler)
		}
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Printf("Allowed origins: %v", cfg.FrontendURLs)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}