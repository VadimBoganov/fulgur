package db

import (
	"database/sql"

	logger2 "github.com/VadimBoganov/fulgur/pkg/logging"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

var logger = logger2.GetLogger()

func NewDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.Fatal(err)
		return nil
	}

	return db
}

func RunMigrations(db *sql.DB) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		logger.Fatal(err)
	}

	fileSource, err := (&file.File{}).Open("file://internal/db/migrations")
	if err != nil {
		logger.Fatal("Cannot open migrations...", err)
	}

	m, err := migrate.NewWithInstance("file", fileSource, "fulgur.db", driver)

	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			logger.Fatal("Cannot execute migrations ", err)
		}

		logger.Info("Have no new migration for current db... ", err)
	}
	logger.Info("Migrations was success")

}
