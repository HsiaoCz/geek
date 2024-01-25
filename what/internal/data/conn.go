package data

import (
	"fmt"

	config "github.com/HsiaoCz/geek/what/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnDataBase() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Mc.User,
		config.Conf.Mc.Passwd,
		config.Conf.Mc.Host,
		config.Conf.Mc.Port,
		config.Conf.Mc.DBname)
	mysqldb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = mysqldb
	return db.AutoMigrate(&User{})
}
