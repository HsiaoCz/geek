package mysql

import (
	"github.com/HsiaoCz/geek/middle/model"
	"github.com/HsiaoCz/geek/middle/settings"
)

// 这里向数据库插入五个管理员

func InsertAdmin(username string, password string) {
	admin := model.Admin{
		Identity: settings.GenIdentity(),
		Username: username,
		Password: settings.SetMd5Password(password),
		Count:    0,
	}
	db.Create(&admin)
}
