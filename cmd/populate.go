package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Aerilate/htn-backend/model"
	"github.com/Aerilate/htn-backend/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var inputFile string

func init() {
	rootCmd.AddCommand(populateCmd)
	populateCmd.Flags().StringVarP(
		&inputFile,
		"inputfile",
		"i",
		DefaultInputFile,
		"JSON file of users to populate the database",
	)
}

var populateCmd = &cobra.Command{
	Use:   "populate",
	Short: "Populates the database with data from an input file",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Inserting mock data from input file %s...\n", inputFile)
		if err := populate(inputFile, SQLitePath); err != nil {
			log.Fatal(err)
		}
		log.Println("...done inserting mock data")
	},
}

// populate loads data from an inputFile, processes it, and writes the processed data to a SQLite file
func populate(inputFile string, sqlFile string) error {
	var users []model.User
	users, err := processUsersJSON(inputFile)
	if err != nil {
		return err
	}

	// set up the ORM
	conn, err := gorm.Open(sqlite.Open(sqlFile), &gorm.Config{})
	if err != nil {
		return err
	}
	// initialize our repository struct to interact with the DB
	newDB := repository.NewRepo(conn)
	if err := newDB.InsertUsers(users); err != nil {
		return err
	}
	return nil
}

// processUsersJSON unmarshalls the data from a given inputFile into a User struct
func processUsersJSON(inputFile string) ([]model.User, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	var users []model.User
	json.Unmarshal([]byte(data), &users)
	return users, nil
}
