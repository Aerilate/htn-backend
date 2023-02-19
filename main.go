package main

import (
	"fmt"
	"github.com/Aerilate/htn-backend/db"
)

const DATA_FILE = "HTN_2023_BE_Challenge_Data.json"

func main() {
	users, _ := processfile(DATA_FILE)
	fmt.Printf("%+v\n", users[0])
	serve(db.DB{users})
}
