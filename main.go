package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/Aerilate/htn-backend/db"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	var insertData = flag.Bool("i", false, "insert mock data?")
	flag.Parse()

	conn, err := sql.Open("sqlite3", SQLITE_FILE)
	if err != nil {
		return err
	}
	if err := db.RunMigration(conn, SQLITE_FILE); err != nil {
		return err
	}
	conn.Close()

	orm, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	newDB := db.NewDB(orm)
	if *insertData {
		log.Println("Inserting mock data...")
		if err := insertMockData(DATA_FILE, newDB); err != nil {
			return err
		}
		log.Println("...done inserting mock data")
	}
	serve(newDB)
	return nil
}
