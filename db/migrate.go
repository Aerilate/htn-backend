package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

const MIGRATION_PATH = "/db/migrate"

func RunMigration(db *sql.DB, dbName string) error {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return err
	}
	path, err := getMigrationPath()
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		dbName, driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func getMigrationPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path + MIGRATION_PATH, nil
}
