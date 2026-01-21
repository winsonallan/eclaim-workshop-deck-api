package main

import (
	"eclaim-workshop-deck-api/internal/config"
	"eclaim-workshop-deck-api/internal/domain/auth"
	"eclaim-workshop-deck-api/internal/domain/authdemo"
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
	if err := db.AutoMigrate(&authdemo.UserDemo{}); err != nil {
		log.Fatal("Failed to migrate UserDemo Table:", err)
	}

	if err := db.AutoMigrate(&authdemo.APIKey{}); err != nil {
		log.Fatal("Failed to migrate APIKey (Demo) Table:", err)
	}

	if err := db.AutoMigrate(&posts.Post{}); err != nil {
		log.Fatal("Failed to migrate Posts Table:", err)
	}

	if err := db.AutoMigrate(&auth.APIKey{}); err != nil {
		log.Fatal("Failed to migrate APIKey Table:", err)
	}

	// Initialize Auth (Demo) domain
	authDemoRepo := authdemo.NewRepository(db)
	authDemoService := authdemo.NewService(authDemoRepo, cfg.JWTSecret)
	authDemoHandler := authdemo.NewHandler(authDemoService)

	// Initialize Posts domain
	postsRepo := posts.NewRepository(db)
	postsService := posts.NewService(postsRepo)
	postsHandler := posts.NewHandler(postsService)

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)
	
	// Setup Gin
	r := gin.Default()

	// Apply CORS middleware (supports multiple origins)
	r.Use(middleware.CORSMiddleware(cfg.FrontendURLs))

	// API routes
	api := r.Group("/api")
	api.Use(middleware.APIKeyMiddleware(db))
	authMiddleware := middleware.AuthMiddleware(cfg.JWTSecret)

	// Apply API Key middleware to ALL API routes
	authdemo.RegisterRoutes(api, authDemoHandler, authMiddleware)
	posts.RegisterRoutes(api, postsHandler, authMiddleware)

	auth.RegisterRoutes(api, authHandler, authMiddleware)
	
	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Printf("Allowed origins: %v", cfg.FrontendURLs)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}