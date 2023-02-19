package db

import (
	"errors"

	"github.com/Aerilate/htn-backend/model"
	mapset "github.com/deckarep/golang-set/v2"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) DB {
	return DB{
		db: db,
	}
}

func (db DB) InsertUsers(users []model.User) error {
	tx := db.db.Create(&users)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (db DB) GetUsers() ([]model.User, error) {
	var users []model.User
	tx := db.db.Joins("LEFT JOIN skill_ratings on skill_ratings.user_id = users.id").
		Group("users.id").
		Preload("SkillRating").
		Find(&users)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return users, nil
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
