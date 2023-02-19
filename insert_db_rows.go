package main

import (
	"encoding/json"
	"os"

	"github.com/Aerilate/htn-backend/model"
)

func processfile(filename string) ([]model.User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []model.User
	json.Unmarshal([]byte(data), &users)
	return users, nil
}
