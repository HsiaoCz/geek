package store

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"cloumn:username;type:varchar(20);" json:"username"`
	Password string `gorm:"cloumn:password;type:varchar(20);" json:"password"`
	Identity int64  `gorm:"cloumn:identity;type:int(11);" json:"identity"`
}
