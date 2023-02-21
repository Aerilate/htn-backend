package repository

import (
	"github.com/Aerilate/htn-backend/model"
	"gorm.io/gorm"
)

const MaxInt = int(^uint(0) >> 1)

type skillRatingRepo struct {
	*gorm.DB
}

func newSkillRatingRepo(db *gorm.DB) skillRatingRepo {
	return skillRatingRepo{DB: db}
}

// AggregateSkills returns a list of skills that have users between minFreq and maxFreq
// To omit a bound, pass in nil
func (s skillRatingRepo) AggregateSkills(minFreq *int, maxFreq *int) ([]model.SkillAggregate, error) {
	// if not specified, give default values to minFreq and maxFreq
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
