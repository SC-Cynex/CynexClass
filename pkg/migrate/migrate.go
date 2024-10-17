package migrate

import (
	"log"
	"path/filepath"

	"github.com/SC-Cynex/cynex-class-service/configs"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	db := configs.GetDB()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not initialize migration driver: %v", err)
	}

	migrationDir, err := filepath.Abs("./migrations")
	if err != nil {
		log.Fatalf("Could not determine absolute path: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationDir,"postgres", driver)
	if err != nil {
		log.Fatalf("Could not initialize migrations: %v", err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply.")
			return
		}
		log.Fatalf("Could not run migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
