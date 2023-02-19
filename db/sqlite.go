package db

import (
	"errors"

	"github.com/Aerilate/htn-backend/model"
	"github.com/imdario/mergo"
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
	if err := db.db.Create(&users).Error; err != nil {
		return err
	}
	return nil
}

func (db DB) GetUsers() ([]model.User, error) {
	var users []model.User
	if err := db.db.Joins("LEFT JOIN skill_ratings on skill_ratings.user_id = users.id").
		Group("users.id").
		Preload("SkillRating").
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (db DB) GetOneUser(id int) (model.User, error) {
	var user model.User
	if err := db.db.Joins("LEFT JOIN skill_ratings on skill_ratings.user_id = users.id").
		Group("users.id").
		Preload("SkillRating").
		First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (db DB) UpdateUser(id int, updatedInfo model.User) error {
	var userToUpdate model.User
	if err := db.db.First(&userToUpdate, id).Error; err != nil {
		return err
	}
	if err := mergo.Merge(&userToUpdate, updatedInfo, mergo.WithOverride); err != nil {
		return err
	}
	if err := db.db.Save(&userToUpdate).Error; err != nil {
		return err
	}

	// update skills
	if updatedInfo.SkillRating == nil {
		return nil
	} else if len(updatedInfo.SkillRating) == 0 {
		db.db.Where("user_id = ?", id).Delete(&model.SkillRating{})
		return nil
	}
	if err := db.db.Model(&userToUpdate).
		Association("SkillRating").
		Replace(updatedInfo.SkillRating); err != nil {
		return err
	}
	return nil
}

func (db DB) GetSkills(minFreq *int, maxFreq *int) ([]model.SkillRating, error) {
	return nil, errors.New("not implemented")
}
