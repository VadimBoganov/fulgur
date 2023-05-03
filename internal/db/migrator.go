package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func NewDB(path string) (*sql.DB, error){
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func RunMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}
	
	fileSource, err := (&file.File{}).Open("file://internal/db/migrations")
	if err != nil{
		return err
	}
	
	m, err := migrate.NewWithInstance("file", fileSource, "fulgur.db", driver)
	
	if err = m.Up(); err != nil {
		return err
	}
	
	return nil
}