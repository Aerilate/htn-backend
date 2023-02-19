package main

import "fmt"

const DATA_FILE = "HTN_2023_BE_Challenge_Data.json"

func main() {
    users, _ := processfile(DATA_FILE)
	fmt.Printf("%+v\n", users[0])
}
