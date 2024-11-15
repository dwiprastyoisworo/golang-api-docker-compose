package migration

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"golang-api/config"
	"log"
)

type Migration struct {
	Postgres config.Postgres
}

func NewMigration(postgres config.Postgres) MigrationInterface {
	return &Migration{Postgres: postgres}
}

type MigrationInterface interface {
	Run()
}

func (m Migration) Run() {
	postgresURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s",
		m.Postgres.Username,
		m.Postgres.Password,
		m.Postgres.Host,
		m.Postgres.Port,
		m.Postgres.Database,
		m.Postgres.SslMode,
		m.Postgres.Schema,
	)

	migratePostgres, err := migrate.New(
		"file://migration/postgres",
		postgresURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = migratePostgres.Up()
	
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migration needed")
		} else {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
	} else {
		log.Println("Migrations applied successfully")
	}
}
