package main

import (
	"encoding/json"
	"os"

	"github.com/Aerilate/htn-backend/model"
)

type UserInserter interface {
	InsertUsers(users []model.User) error 
}

func insertMockData(filename string, userInserter UserInserter ) error {
	var users []model.User
	users, err := processfile(filename)
	if err != nil {
		return err
	}
	if err := userInserter.InsertUsers(users); err != nil {
		return err
	}
	return nil
}

func processfile(filename string) ([]model.User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []model.User
	json.Unmarshal([]byte(data), &users)
	return users, nil
}

