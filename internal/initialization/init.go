package initialization

import (
	"database/sql"

	"github.com/SC-Cynex/cynex-class-service/internal/api/auth"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
)

type AppDependencies struct {
	AuthHandler *auth.Handler
}

func InitAppDependencies(db *sql.DB) (*AppDependencies, error) {
	// Repositories
	authRepo := repository.NewAuthRepository(db)

	// Services
	authService := services.NewAuthService(authRepo)

	// Handlers
	authHandler := auth.NewHandler(authService)

	return &AppDependencies{
		AuthHandler: authHandler,
	}, nil
}
