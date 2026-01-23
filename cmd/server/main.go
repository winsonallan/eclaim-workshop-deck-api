package main

import (
	"eclaim-workshop-deck-api/internal/bootstrap"
	"eclaim-workshop-deck-api/internal/config"
	"eclaim-workshop-deck-api/internal/domain/auth"
	"eclaim-workshop-deck-api/internal/domain/authdemo"
	"eclaim-workshop-deck-api/internal/domain/posts"
	"eclaim-workshop-deck-api/internal/middleware"
	"eclaim-workshop-deck-api/internal/migrations"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect DB
	db := config.ConnectDB(cfg)

	// Run migrations
	migrations.Run(db)

	// Init domains
	domains := bootstrap.InitDomains(db, cfg)

	// Setup Gin
	r := gin.Default()
	r.Use(middleware.CORSMiddleware(cfg.FrontendURLs))

	api := r.Group("/api")
	api.Use(middleware.APIKeyMiddleware(db))

	authMiddleware := middleware.AuthMiddleware(cfg.JWTSecret)

	authdemo.RegisterRoutes(api, domains.AuthDemoHandler, authMiddleware)
	posts.RegisterRoutes(api, domains.PostsHandler, authMiddleware)
	auth.RegisterRoutes(api, domains.AuthHandler, authMiddleware)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
