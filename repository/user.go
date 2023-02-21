package repository

import (
	"reflect"

	"github.com/Aerilate/htn-backend/model"
	"github.com/imdario/mergo"
	"gorm.io/gorm"
)

type UserRepo struct {
	*gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{DB: db}
}

func (u UserRepo) InsertUsers(users []model.User) error {
	if err := u.Create(&users).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepo) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := u.Joins("LEFT JOIN skill_ratings on skill_ratings.user_id = users.id").
		Group("users.id").
		Preload("SkillRating").
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserRepo) GetUser(id int) (model.User, error) {
	var user model.User
	if err := u.Joins("LEFT JOIN skill_ratings on skill_ratings.user_id = users.id").
		Group("users.id").
		Preload("SkillRating").
		First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepo) UpdateUser(id int, updatedInfo model.User) error {
	var userToUpdate model.User
	if err := u.First(&userToUpdate, id).Error; err != nil {
		return err
	}
	if err := mergo.Merge(
		&userToUpdate,
		updatedInfo,
		mergo.WithTransformers(userTransformer{})); err != nil {
		return err
	}
	if err := u.Save(&userToUpdate).Error; err != nil {
		return err
	}

	if updatedInfo.SkillRating == nil {
		return nil
	}
	if err := u.Where("user_id = ?", id).Delete(&model.SkillRating{}).Error; err != nil {
		return err
	}
	if err := u.Model(&userToUpdate).
		Association("SkillRating").
		Replace(updatedInfo.SkillRating); err != nil {
		return err
	}
	return nil
}

type userTransformer struct{}

func (u userTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	var s string
	if typ != reflect.TypeOf(&s) {
		return nil
	}
	return func(dst, src reflect.Value) error {
		if dst.CanSet() && src.UnsafePointer() != nil {
			dst.Set(src)
		}
		return nil
	}
}
