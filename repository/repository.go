package repository

import (
	"gorm.io/gorm"
)

type Repo struct {
	UserRepo
	SkillRatingRepo
}

func NewRepo(db *gorm.DB) Repo {
	return Repo{
		NewUserRepo(db),
		NewSkillRatingRepo(db),
	}
}
