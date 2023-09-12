package model

import "gorm.io/gorm"

// 这里只涉及用户一张表
// 这里的用户表应该怎么设计呢?
type User struct {
	gorm.Model
	Username    string `gorm:"column:username;type:varchar(20);" json:"username"`
	Passwrod    string `gorm:"column:passwrod;type:varchar(20);" json:"password"`
	Identity    int    `gorm:"column:identity;type:int(11);" json:"identity"`
	Level       int    `gorm:"column:level;type:int(3);" json:"level"`
	Birthday    string `gorm:"column:birthday;type:varchar(10);" json:"birthday"`
	Job         string `gorm:"column:job;type:varchar(20);" json:"job"`
	Signature   string `gorm:"column:signature;type:varcahr(200);" json:"signature"`
	Address     string `gorm:"column:address;type:varchar(100);" json:"address"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(20);" json:"phone_number"`
	Content     string `gorm:"column:content;type:varchar(200);" json:"content"`
}

func (u User) TableName() string {
	return "user"
}
