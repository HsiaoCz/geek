package model

type Admin struct {
	Identity string `gorm:"column:identity;type:varchar(20);" json:"identity"`
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Passwrod string `gorm:"column:password;type:varchar(20);" json:"passwrod"`
}

func (a Admin) TableName() string {
	return "admin"
}
