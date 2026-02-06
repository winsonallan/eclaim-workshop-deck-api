package main

import (
	"eclaim-workshop-deck-api/internal/bootstrap"
	"eclaim-workshop-deck-api/internal/config"
	"eclaim-workshop-deck-api/internal/domain/admin"
	"eclaim-workshop-deck-api/internal/domain/auth"
	"eclaim-workshop-deck-api/internal/domain/authdemo"
	"eclaim-workshop-deck-api/internal/domain/location"
	"eclaim-workshop-deck-api/internal/domain/panels"
	"eclaim-workshop-deck-api/internal/domain/posts"
	"eclaim-workshop-deck-api/internal/domain/settings"
	"eclaim-workshop-deck-api/internal/domain/suppliers"
	"eclaim-workshop-deck-api/internal/domain/usermanagement"
	"eclaim-workshop-deck-api/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect DB
	db := config.ConnectDB(cfg)

	db.Debug()
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

	admin.RegisterRoutes(api, domains.AdminsHandler, authMiddleware)

	panels.RegisterRoutes(api, domains.PanelsHandler, authMiddleware)
	settings.RegisterRoutes(api, domains.SettingsHandler, authMiddleware)
	location.RegisterRoutes(api, domains.LocationHandler, authMiddleware)
	suppliers.RegisterRoutes(api, domains.SupplierHandler, authMiddleware)
	usermanagement.RegisterRoutes(api, domains.UserManagementHandler, authMiddleware)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
