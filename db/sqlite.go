package db

import (
	// "database/sql"
	"errors"

	"github.com/Aerilate/htn-backend/model"
	mapset "github.com/deckarep/golang-set/v2"
	// "github.com/mattn/go-sqlite3"
	// "gorm.io/driver/sqlite"
  	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) DB {
	return DB {
		db: db,
	}
}

func (db DB) GetUsers() ([]model.User, error) {
	return nil, nil
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
