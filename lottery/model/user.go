package model

type User struct {
	Identity string `gorm:"column:identity;type:varchar(20);" json:"identity"`
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Passwrod string `gorm:"column:password;type:varchar(20);" json:"passwrod"`
	// 这里需要订阅一个消息队列
}

func (u User) TableName() string {
	return "user"
}
