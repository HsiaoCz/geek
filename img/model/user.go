package model

import "gorm.io/gorm"

// 这里只涉及用户一张表
// 这里的用户表应该怎么设计呢?
type User struct {
	gorm.Model
	// 用户名称，不过这里是用户昵称
	// 如果昵称违规怎么办?
	// 可以维护一个敏感词的切片
	// 循环取出切片里的敏感词与用户的昵称进行匹配
	// 如果出现了敏感词，就告知用户该昵称不可用
	// 还有一个问题?如果昵称重复？
	// 用户输入一个昵称，前端拿到这个昵称发给后端，后端拿到昵称后去数据库里查一下，看有没有这个昵称
	// 有则提示不可用
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Passwrod string `gorm:"column:passwrod;type:varchar(20);" json:"password"`
	// 这里生成一个qq号
	// 随机数可能不行，因为随机数可能会出现重复的
	Identity int `gorm:"column:identity;type:int(11);" json:"identity"`
	// 用户等级，这里默认为0
	// 那么怎么提升等级呢?
	Level    int    `gorm:"column:level;type:int(3);" json:"level"`
	Birthday string `gorm:"column:birthday;type:varchar(10);" json:"birthday"`
	Job      string `gorm:"column:job;type:varchar(20);" json:"job"`
	// 类似于个性签名
	Signature   string `gorm:"column:signature;type:varcahr(200);" json:"signature"`
	Address     string `gorm:"column:address;type:varchar(100);" json:"address"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(20);" json:"phone_number"`
	// 类似于个人简介
	Content string `gorm:"column:content;type:varchar(200);" json:"content"`
}

func (u User) TableName() string {
	return "user"
}
