package bootstrap

import (
	"eclaim-workshop-deck-api/internal/config"
	"eclaim-workshop-deck-api/internal/domain/admin"
	"eclaim-workshop-deck-api/internal/domain/auth"
	"eclaim-workshop-deck-api/internal/domain/authdemo"
	"eclaim-workshop-deck-api/internal/domain/location"
	"eclaim-workshop-deck-api/internal/domain/panels"
	"eclaim-workshop-deck-api/internal/domain/posts"
	"eclaim-workshop-deck-api/internal/domain/settings"

	"gorm.io/gorm"
)

type Domains struct {
	AuthDemoHandler *authdemo.Handler
	PostsHandler    *posts.Handler
	AuthHandler     *auth.Handler
	AdminsHandler   *admin.Handler
	PanelsHandler   *panels.Handler
	SettingsHandler *settings.Handler
	LocationHandler *location.Handler
}

func InitDomains(db *gorm.DB, cfg *config.Config) *Domains {
	// Auth Demo
	authDemoRepo := authdemo.NewRepository(db)
	authDemoService := authdemo.NewService(authDemoRepo, cfg.JWTSecret)
	authDemoHandler := authdemo.NewHandler(authDemoService)

	// Posts
	postsRepo := posts.NewRepository(db)
	postsService := posts.NewService(postsRepo)
	postsHandler := posts.NewHandler(postsService)

	// Auth
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)

	adminsRepo := admin.NewRepository(db)
	adminsService := admin.NewService(adminsRepo)
	adminsHandler := admin.NewHandler(adminsService)

	// Panels
	panelsRepo := panels.NewRepository(db)
	panelsService := panels.NewService(panelsRepo)
	panelsHandler := panels.NewHandler(panelsService)

	// Settings
	settingsRepo := settings.NewRepository(db)
	settingsService := settings.NewService(settingsRepo)
	settingsHandler := settings.NewHandler(settingsService)

	// Location
	locationRepo := location.NewRepository(db)
	locationService := location.NewService(locationRepo)
	locationHandler := location.NewHandler(locationService)

	return &Domains{
		AuthDemoHandler: authDemoHandler,
		PostsHandler:    postsHandler,
		AuthHandler:     authHandler,
		AdminsHandler:   adminsHandler,
		PanelsHandler:   panelsHandler,
		SettingsHandler: settingsHandler,
		LocationHandler: locationHandler,
	}
}
