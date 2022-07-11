// nolint: revive
package migration

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/namnguyen/backend/config"
)

func Up(db *gorm.DB) {
	getDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(getDB, &postgres.Config{MigrationsTable: "migration"})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migration", config.GetConfig().Postgresql.DBName, driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
}
