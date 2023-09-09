package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Identity int    `gorm:"column:identity;type:int(11);" json:"identity"`
	Password string `gorm:"column:password;type:varchar(20);" json:"passwrod"`
	Job      string `gorm:"column:job;type:varchar(20);" json:"job"`
	Content  string `gorm:"column:content;type:varchar(200);" json:"content"`
	Email    string `gorm:"column:email;type:varchar(30);" json:"email"`
}

func (u User) TableName() string {
	return "user"
}
