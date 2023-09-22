package mysql

import (
	"github.com/HsiaoCz/geek/img/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBInstance struct {
	db *gorm.DB
}

var DB DBInstance

func ConnMysql() (err error) {
	dsn := "root:shaw123@tcp(127.0.0.1:3306)/img?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB = DBInstance{db: db}
	return DB.db.AutoMigrate(&model.User{})
}
