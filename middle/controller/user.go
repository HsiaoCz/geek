package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HsiaoCz/geek/middle/dao/mysql"
	"github.com/HsiaoCz/geek/middle/model"
	"github.com/HsiaoCz/geek/middle/queue"
)

func HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	userR := new(UserR)
	err := json.NewDecoder(r.Body).Decode(userR)
	if err != nil {
		ResponseJSON(w, http.StatusOK, "请求参数出错", nil)
		return
	}
	if userR.Password != userR.RePasswrod {
		ResponseJSON(w, http.StatusOK, "请检查密码和确认密码是否一致", nil)
		return
	}
	if err := mysql.UserRegister(userR.Username, userR.Password, userR.Email); err != nil {
		log.Println(err)
		return
	}
	ResponseJSON(w, http.StatusOK, "注册成功", nil)
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	userL := new(UserL)
	err := json.NewDecoder(r.Body).Decode(userL)
	if err != nil {
		ResponseJSON(w, http.StatusOK, "请求参数出错", nil)
		return
	}
	if err := mysql.GetUserByUsernameAndPassword(userL.Username, userL.Password); err != nil {
		ResponseJSON(w, http.StatusOK, "请检查用户名和密码", nil)
		return
	}
	ResponseJSON(w, http.StatusOK, "登录成功", nil)
}

func HandleUserPostArticle(w http.ResponseWriter, r *http.Request) {
	article := new(model.Article)
	json.NewDecoder(r.Body).Decode(article)
	err := queue.MQueue.InQueue(*article)
	if err != nil {
		log.Fatal(err)
	}
	ResponseJSON(w, http.StatusOK, "提交成功", nil)
}
