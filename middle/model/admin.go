package model

import "gorm.io/gorm"

// 管理员的结构体
// 管理员可以审核文章
// 但这里的问题是，这个管理员的类应该有哪些字段呢?
// 这里暂时先写这些
type Admin struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Identity int    `gorm:"column:identity;type:int(11);" json:"identity"`
	Password string `gorm:"column:password;type:varchar(20);" json:"password"`
	Count    int    `gorm:"column:count;type:int(11);" json:"count"`
}

func (a Admin) TableName() string {
	return "admin"
}
