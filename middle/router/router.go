package router

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/geek/middle/controller"
)

func RegisterRouter(addr string) {
	http.HandleFunc("/user/register", controller.HandleUserRegister)
	http.HandleFunc("/user/login", controller.HandleUserLogin)
	http.HandleFunc("/user/post/article", controller.HandleUserPostArticle)

	http.HandleFunc("/admin/login", controller.HandleAdminLogin)
	http.HandleFunc("/admin/get/articles", controller.HandleAdminGetArticles)
	srv := http.Server{
		Handler:      nil,
		Addr:         addr,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}
