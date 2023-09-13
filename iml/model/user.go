package model

import "gorm.io/gorm"

// user Model
type User struct {
	gorm.Model
	Identity int64  `gorm:"column:identity;type:int(11);" json:"identity"`
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Password string `gorm:"column:password;type:varchar(50);" json:"password"`
	Content  string `gorm:"column:content;type:varchar(50)" json:"content"`
}
