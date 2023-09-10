package mysql

import (
	"github.com/HsiaoCz/geek/lottery/settings"
	"github.com/HsiaoCz/geek/middle/model"
)

func GetAdminByUsernameAndPassword(username string, password string) (err error) {
	admin := new(model.Admin)
	tx := db.Where("username = ? AND password = ?", username, settings.SetMd5Password(password)).Take(admin)
	return tx.Error
}
