package main

import (
	"database/sql"
	// "fmt"
	"log"

	"github.com/Aerilate/htn-backend/db"
	_ "github.com/mattn/go-sqlite3"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
)

const DATA_FILE = "HTN_2023_BE_Challenge_Data.json"
const SQLITE_FILE = "sqlite.db"
const DB_NAME = "sqlite"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	conn, err := sql.Open("sqlite3", SQLITE_FILE)
	if err != nil {
		return err
	}
	if err := db.RunMigration(conn, SQLITE_FILE); err != nil {
		return err
	}

	// conn, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	// if err != nil {
	// 	return err
	// }

	// serve(db.NewDB(conn))
	return nil
}
