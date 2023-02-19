package db

import (
	"errors"
	"github.com/Aerilate/htn-backend/model"
	mapset "github.com/deckarep/golang-set/v2"
	// "github.com/mattn/go-sqlite3"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
)

type DB struct {
	Users []model.User
	// conn *DB
}

func init() {
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}

func (db DB) GetUsers() ([]model.User, error) {
	return db.Users, nil
}

func (DB) GetOneUser(id int) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (DB) UpdateUser(id int, updatedInfo model.User, keysToUpdate mapset.Set[string]) error {
	return errors.New("not implemented")
}

func (DB) GetSkills(minFreq *int, maxFreq *int) ([]model.SkillRating, error) {
	return nil, errors.New("not implemented")
}
