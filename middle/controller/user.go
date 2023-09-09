package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HsiaoCz/geek/middle/dao/mysql"
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
