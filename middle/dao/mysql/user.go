package mysql

import (
	"github.com/HsiaoCz/geek/middle/model"
	"github.com/HsiaoCz/geek/middle/settings"
)

func UserRegister(username string, password string, email string) (err error) {
	user := model.User{
		Username: username,
		Password: settings.SetMd5Password(password),
		Email:    email,
		Identity: settings.GenIdentity(),
	}
	tx := db.Create(&user)
	return tx.Error
}

func GetUserByUsernameAndPassword(username string, password string) error {
	user := new(model.User)
	tx := db.Where("username = ? AND  password=?", username, settings.SetMd5Password(password)).Take(user)
	return tx.Error
}
