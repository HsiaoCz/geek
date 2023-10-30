package dao

import "github.com/HsiaoCz/geek/Todo/model"

func AuthReg(username string, password string, identity int64, email string) error {
	user := &model.User{
		Username: username,
		Password: password,
		Identity: identity,
		Email:    email,
	}
	ctx := db.Create(user)
	return ctx.Error
}
