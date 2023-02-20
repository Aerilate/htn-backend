package cmd

type Config struct {
	inputDataFile string
	migrationDir  string
	sqliteConfig  SQLiteConfig
}

type SQLiteConfig struct {
	dbName string
	file   string
}

const DefaultInputFile = "HTN_2023_BE_Challenge_Data.json"
const SQLitePath = "sqlite.db"
const DBName = "/sqlite"
const MigrationPath = "/migration"
