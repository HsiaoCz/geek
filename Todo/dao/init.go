package dao

import (
	"fmt"

	"github.com/HsiaoCz/geek/Todo/config"
	"github.com/HsiaoCz/geek/Todo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.MC.User,
		config.Conf.MC.Passwd,
		config.Conf.MC.Host,
		config.Conf.MC.Port,
		config.Conf.MC.DBname)
	mysqldb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	db = mysqldb
	return db.AutoMigrate(&model.User{}, &model.TodoList{})
}
