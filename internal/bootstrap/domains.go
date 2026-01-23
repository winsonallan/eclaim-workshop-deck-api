package bootstrap

import (
	"eclaim-workshop-deck-api/internal/config"
	"eclaim-workshop-deck-api/internal/domain/auth"
	"eclaim-workshop-deck-api/internal/domain/authdemo"
	"eclaim-workshop-deck-api/internal/domain/panels"
	"eclaim-workshop-deck-api/internal/domain/posts"

	"gorm.io/gorm"
)

type Domains struct {
	AuthDemoHandler *authdemo.Handler
	PostsHandler    *posts.Handler
	AuthHandler     *auth.Handler
	PanelsHandler   *panels.Handler
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

	// Panels
	panelsRepo := panels.NewRepository(db)
	panelsService := panels.NewService(panelsRepo)
	panelsHandler := panels.NewHandler(panelsService)

	return &Domains{
		AuthDemoHandler: authDemoHandler,
		PostsHandler:    postsHandler,
		AuthHandler:     authHandler,
		PanelsHandler:   panelsHandler,
	}
}
