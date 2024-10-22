package migrate

import (
	"log"
	"os"
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
		log.Fatalf("Não foi possível inicializar o driver de migração: %v", err)
	}

	// Lista de possíveis caminhos para o diretório de migrações
	possiblePaths := []string{
		"./migrations",
		"../migrations",
		"../../migrations",
	}

	var migrationDir string
	for _, path := range possiblePaths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Fatalf("Não foi possível determinar o caminho absoluto: %v", err)
		}

		if _, err := os.Stat(absPath); !os.IsNotExist(err) {
			migrationDir = absPath
			break
		}
	}

	if migrationDir == "" {
		log.Fatalf("Diretório de migrações não encontrado em nenhum dos caminhos possíveis")
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationDir, "postgres", driver)
	if err != nil {
		log.Fatalf("Não foi possível inicializar as migrações: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Não foi possível executar as migrações: %v", err)
	}

	if err == migrate.ErrNoChange {
		log.Println("Nenhuma migração para aplicar.")
	} else {
		log.Println("Migrações aplicadas com sucesso!")
	}
}
