package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Password string `gorm:"column:password;type:varchar(20);" json:"password"`
<<<<<<< HEAD
=======
	Identity int64  `gorm:"column:identity;type:int(11);" json:"identity"`
	Content  string `gorm:"column:content;type:varchar(300);" json:"content"`
	Level    uint8  `gorm:"column:level;type:int(3);" json:"level"`
	Birthday string `gorm:"column:birthday;type:varchar(20);" json:"birthday"`
	Varval   string `gorm:"column:varval;type:varchar(100);" json:"varval"`
	Job      string `gorm:"column:job;type:varchar(60);" json:"job"`
}

func (u User) TableName() string {
	return "user"
>>>>>>> 0d7b9b4fe7047e20192bd97eaf83da34fc3f19e2
}
