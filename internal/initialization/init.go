package initialization

import (
	"database/sql"
	"github.com/SC-Cynex/cynex-class-service/internal/api/teachers"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
)

type AppDependencies struct {
	TeacherHandler *teachers.Handler
}

func InitAppDependencies(db *sql.DB) (*AppDependencies, error) {
	// Repositories
	teacherRepo := repository.NewTeacherRepository(db)

	// Services
	teacherService := services.NewTeacherService(teacherRepo)

	// Handlers
	teacherHandler := teachers.NewHandler(teacherService)

	return &AppDependencies{
		TeacherHandler: teacherHandler,
	}, nil
}
