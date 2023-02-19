package model

type User struct {
	Name string
	Company string
	Email string
	Phone string
	Skills []SkillRating
}

type SkillRating struct {
	Skill string
	Rating int
}
