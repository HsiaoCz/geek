package dao

import (
	"fmt"
	"log"
	"testing"

	"github.com/HsiaoCz/geek/Todo/config"
	"github.com/HsiaoCz/geek/Todo/model"
	"github.com/HsiaoCz/geek/Todo/utils"
)

func TestAuthCreate(t *testing.T) {
	err := config.InitConf()
	if err != nil {
		log.Fatal(err)
	}
	err = InitMysql()
	if err != nil {
		t.Fatal(err)
	}
	user := model.User{
		Username: "liii",
		Password: utils.SetMd5Password("1244556"),
		Identity: utils.GenIdentity(),
		Email:    "1234555@qq.com",
		Avatar:   "",
	}
	res := AuthReg(user.Username, user.Password, user.Identity, user.Email)
	fmt.Println(res.Error())
}

func TestGetUserByUsernameAndPassword(t *testing.T) {
	err := config.InitConf()
	if err != nil {
		log.Fatal(err)
	}
	err = InitMysql()
	if err != nil {
		t.Fatal(err)
	}
	user := new(model.User)
	result := db.Where("username = ? AND password = ?", "zhangsan", utils.SetMd5Password("12333")).Find(user)
	fmt.Println(result.RowsAffected)
}

func TestGetUserByUsernameAndEmail(t *testing.T) {
	err := config.InitConf()
	if err != nil {
		log.Fatal(err)
	}
	err = InitMysql()
	if err != nil {
		t.Fatal(err)
	}
	result := AuthGetUserByUsernameAndEmail("zhangsan", "1233765@qq.com")
	fmt.Println(result)
}
