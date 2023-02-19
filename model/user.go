package model

type User struct {
	ID          int           `gorm:"primaryKey" json:"-"`
	Name        *string       `json:"name"`
	Company     *string       `json:"company"`
	Email       *string       `json:"email"`
	Phone       *string       `json:"phone"`
	SkillRating []SkillRating `gorm:"foreignKey:UserID" json:"skills"`
}

type SkillRating struct {
	UserID int    `gorm:"primaryKey" json:"-"`
	Skill  string `gorm:"primaryKey" json:"skill"`
	Rating int    `json:"rating"`
}
