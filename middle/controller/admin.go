package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HsiaoCz/geek/middle/dao/mysql"
)

func HandleAdminLogin(w http.ResponseWriter, r *http.Request) {
	adminL := new(AdminL)
	err := json.NewDecoder(r.Body).Decode(adminL)
	if err != nil {
		log.Println(err)
		return
	}
	if err := mysql.GetAdminByUsernameAndPassword(adminL.Username, adminL.Password); err != nil {
		log.Println(err)
		ResponseJSON(w, http.StatusOK, "请检查用户名或密码", nil)
		return
	}
	ResponseJSON(w, http.StatusOK, "登录成功", nil)
}

func HandleAdminGetArticles(w http.ResponseWriter, r *http.Request) {}
