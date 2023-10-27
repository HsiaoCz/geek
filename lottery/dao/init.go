package dao

import (
	"github.com/HsiaoCz/geek/middle/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMysql() (err error) {
	dsn := "root:shaw123@tcp(127.0.0.1:3306)/middle?charset=utf8mb4&parseTime=True&loc=Local"
	mysqldb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	db = mysqldb
	return db.AutoMigrate(&model.User{}, &model.Article{}, &model.Admin{})
}
