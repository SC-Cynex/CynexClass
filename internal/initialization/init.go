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
	CourseHandler  *handlers.CourseHandler
	SubjectHandler  *handlers.SubjectHandler
	ClassHandler  *handlers.ClassHandler
	InscriptionHandler  *handlers.InscriptionHandler
}

func InitAppDependencies(db *sql.DB) (*AppDependencies, error) {
	// Repositories
	addressRepo := repository.NewAddressRepository(db)
	authRepo := repository.NewAuthRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)
	classRepo := repository.NewClassRepository(db)
	inscriptionRepo := repository.NewInscriptionRepository(db)

	// Services
	addressService := services.NewAddressService(addressRepo)
	authService := services.NewAuthService(authRepo, addressService)
	courseService := services.NewCourseService(courseRepo)
	subjectService := services.NewSubjectService(subjectRepo)
	classService := services.NewClassService(classRepo)
	inscriptionService := services.NewInscriptionService(inscriptionRepo)

	// Handlers
	addressHandler := handlers.NewHandler(addressService)
	authHandler := handlers.NewAuthHandler(authService)
	courseHandler := handlers.NewCourseHandler(courseService)
	subjectHandler := handlers.NewSubjectHandler(subjectService)
	classHandler := handlers.NewClassHandler(classService)
	inscriptionHandler := handlers.NewInscriptionHandler(inscriptionService)

	return &AppDependencies{
		AuthHandler:    authHandler,
		AddressHandler: addressHandler,
		CourseHandler: courseHandler,
		SubjectHandler: subjectHandler,
		ClassHandler: classHandler,
		InscriptionHandler: inscriptionHandler,
	}, nil
}
