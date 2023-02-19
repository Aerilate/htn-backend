package model

type User struct {
	ID		int
	Name    string
	Company string
	Email   string
	Phone   string
	Skills  []SkillRating
}

type SkillRating struct {
	ID		int
	Skill  string
	Rating int
}
