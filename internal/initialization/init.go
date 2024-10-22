package initialization

import (
	"database/sql"

	"github.com/SC-Cynex/cynex-class-service/internal/handlers"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
)

type AppDependencies struct {
	AuthHandler    *handlers.AuthHandler
	AddressHandler *handlers.Handler
}

func InitAppDependencies(db *sql.DB) (*AppDependencies, error) {
	// Repositories
	addressRepo := repository.NewAddressRepository(db)
	authRepo := repository.NewAuthRepository(db)

	// Services
	addressService := services.NewAddressService(addressRepo)
	authService := services.NewAuthService(authRepo, addressService)

	// Handlers
	addressHandler := handlers.NewHandler(addressService)
	authHandler := handlers.NewAuthHandler(authService)

	return &AppDependencies{
		AuthHandler:    authHandler,
		AddressHandler: addressHandler,
	}, nil
}
