package dao

import (
	"github.com/HsiaoCz/geek/Todo/model"
	"github.com/HsiaoCz/geek/Todo/utils"
)

func AuthReg(username string, password string, identity int64, email string) error {
	user := &model.User{
		Username: username,
		Password: utils.SetMd5Password(password),
		Identity: identity,
		Email:    email,
	}
	ctx := db.Create(user)
	return ctx.Error
}

func AuthGetPasswd(username string, password string) error {
	tx := db.Where("username = ? AND password = ?", username, utils.SetMd5Password(password)).First(&model.User{})
	return tx.Error
}
