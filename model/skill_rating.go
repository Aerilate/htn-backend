package model

type SkillRating struct {
	UserID int    `gorm:"primaryKey" json:"-"`
	Skill  string `gorm:"primaryKey" json:"skill"`
	Rating int    `json:"rating"`
}

type SkillAggregate struct {
	Skill string
	Count int
}
