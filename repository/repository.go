package repository

import (
	"gorm.io/gorm"
)

type Repo struct {
	userRepo
	skillRatingRepo
}

func NewRepo(db *gorm.DB) Repo {
	return Repo{
		newUserRepo(db),
		newSkillRatingRepo(db),
	}
}
