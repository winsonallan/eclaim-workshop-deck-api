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

	// Auto migrate
	if err := db.AutoMigrate(&auth.User{}, &posts.Post{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
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
	r.Use(middleware.CORSMiddleware(cfg.FrontendURL))

	// API routes
	api := r.Group("/api")

	// Auth routes (public + protected)
	authMiddleware := middleware.AuthMiddleware(cfg.JWTSecret)
	auth.RegisterRoutes(api, authHandler, authMiddleware)

	// Protected routes
	protected := api.Group("")
	protected.Use(authMiddleware)
	{
		posts.RegisterRoutes(protected, postsHandler)
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}