package cmd

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run data migrations on the database",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Running migrations...")
		if err := runMigration(SQLitePath, DBName); err != nil {
			log.Fatal(err)
		}
		log.Println("...done migrations")
	},
}

func runMigration(sqlFile string, dbName string) error {
	db, err := sql.Open("sqlite3", sqlFile)
	if err != nil {
		return err
	}
	defer db.Close()

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
	return path + "/" + MigrationPath, nil
}
