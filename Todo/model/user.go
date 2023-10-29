package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity int64  `gorm:"column:identity" json:"identity"`
	Username string `gorm:"column:username;unique" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email;unique" json:"email"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
}

func (u User) TableName() string {
	return "user"
}
