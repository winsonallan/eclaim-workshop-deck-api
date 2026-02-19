package main

import (
	"eclaim-workshop-deck-api/internal/bootstrap"
	"eclaim-workshop-deck-api/internal/config"
	"eclaim-workshop-deck-api/internal/domain/admin"
	"eclaim-workshop-deck-api/internal/domain/auth"
	"eclaim-workshop-deck-api/internal/domain/authdemo"
	"eclaim-workshop-deck-api/internal/domain/location"
	"eclaim-workshop-deck-api/internal/domain/orders"
	"eclaim-workshop-deck-api/internal/domain/panels"
	"eclaim-workshop-deck-api/internal/domain/posts"
	"eclaim-workshop-deck-api/internal/domain/settings"
	"eclaim-workshop-deck-api/internal/domain/suppliers"
	"eclaim-workshop-deck-api/internal/domain/usermanagement"
	"eclaim-workshop-deck-api/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig()

	logger, err := config.NewLogger(cfg.Env)
	if err != nil {
		log.Fatal("failed to initialize logger: ", err)
	}
	defer logger.Sync()

	gin.SetMode(cfg.GinMode)

	db := config.ConnectDB(cfg)

	domains := bootstrap.InitDomains(db, cfg, logger)

	r := gin.New()

	// Global middleware — order matters
	r.Use(middleware.RequestID())    // ← new
	r.Use(middleware.Logger(logger)) // ← new (your structured logger)
	r.Use(gin.Recovery())            // ← replaces the one from gin.Default()
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
	orders.RegisterRoutes(api, domains.OrdersHandler, authMiddleware)

	// ✅ Use logger instead of log.Printf for consistency
	logger.Info("server starting", zap.String("port", cfg.Port))
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Fatal("server failed to start", zap.Error(err))
	}
}
