package repository

import (
	"reflect"

	"github.com/Aerilate/htn-backend/model"
	"github.com/imdario/mergo"
	"gorm.io/gorm"
)

type userRepo struct {
	*gorm.DB
}

func newUserRepo(db *gorm.DB) userRepo {
	return userRepo{DB: db}
}

func (u userRepo) InsertUsers(users []model.User) error {
	if err := u.Create(&users).Error; err != nil {
		return err
	}
	return nil
}

func (u userRepo) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := u.Joins("LEFT JOIN skill_ratings on skill_ratings.user_id = users.id").
		Group("users.id").
		Preload("SkillRating").
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u userRepo) GetUser(id int) (model.User, error) {
	var user model.User
	if err := u.Joins("LEFT JOIN skill_ratings on skill_ratings.user_id = users.id").
		Group("users.id").
		Preload("SkillRating").
		First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

// UpdateUser will update the User of the given id with values in updatedInfo.
// For *string fields in updatedInfo:
//  - nil will not overwrite the existing value
//  - a pointer to "" will overwrite
// For the []SkillRating field in updatedInfo:
//  - nil will not overwrite
//  - the empty slice will overwrite
func (u userRepo) UpdateUser(id int, updatedInfo model.User) error {
	// get the original user information
	var userToUpdate model.User
	if err := u.First(&userToUpdate, id).Error; err != nil {
		return err
	}
	// merge the two User structures
	if err := mergo.Merge(
		&userToUpdate,
		updatedInfo,
		mergo.WithTransformers(userTransformer{})); err != nil {
		return err
	}
	// write updated user information
	if err := u.Save(&userToUpdate).Error; err != nil {
		return err
	}

	// updated SkillRating was not specified, so do not overwrite
	if updatedInfo.SkillRating == nil {
		return nil
	}
	// clear previous skills
	if err := u.Where("user_id = ?", id).Delete(&model.SkillRating{}).Error; err != nil {
		return err
	}
	// write updated skills
	if err := u.Model(&userToUpdate).
		Association("SkillRating").
		Replace(updatedInfo.SkillRating); err != nil {
		return err
	}
	return nil
}

// userTransformer is defined because mergo doesn't merge *string types
type userTransformer struct{}

func (u userTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	var s string
	// check if type is *string
	if typ != reflect.TypeOf(&s) {
		return nil
	}
	return func(dst, src reflect.Value) error {
		if dst.CanSet() && src.UnsafePointer() != nil {
			// overwrite dst value with src
			dst.Set(src)
		}
		return nil
	}
}
