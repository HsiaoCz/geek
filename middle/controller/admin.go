package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HsiaoCz/geek/middle/dao/mysql"
	"github.com/HsiaoCz/geek/middle/queue"
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

func HandleAdminGetArticles(w http.ResponseWriter, r *http.Request) {
	article, err := queue.MQueue.OutQueue()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(article)
	ResponseJSON(w, http.StatusOK, "处理完毕", nil)
}
