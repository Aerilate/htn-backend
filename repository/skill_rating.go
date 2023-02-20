package repository

import (
	"github.com/Aerilate/htn-backend/model"
	"gorm.io/gorm"
)

const MaxInt = int(^uint(0) >> 1)

type SkillRatingRepo struct {
	*gorm.DB
}

func NewSkillRatingRepo(db *gorm.DB) SkillRatingRepo {
	return SkillRatingRepo{DB: db}
}

func (s SkillRatingRepo) AggregateSkills(minFreq *int, maxFreq *int) ([]model.SkillAggregate, error) {
	if minFreq == nil {
		minFreq = intPtr(0)
	}
	if maxFreq == nil {
		maxFreq = intPtr(MaxInt)
	}

	var result []model.SkillAggregate
	if err := s.Model(&model.SkillRating{}).
		Select("skill, count(*) as count").
		Group("skill").
		Having("count BETWEEN ? AND ?", *minFreq, *maxFreq).
		Order("count DESC").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func intPtr(i int) *int {
	return &i
}
