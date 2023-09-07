package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"
)

const (
	addr = "127.0.0.1:9091"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"passwrod"`
}

type Article struct {
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Auther     string    `json:"auther"`
	Topic      string    `json:"topic"`
	CreateTime time.Time `json:"create_time"`
}

type H map[string]any

func main() {
	flag.Parse()
	http.HandleFunc("/v1/helloworld/user/login", HandleUserLogin)
	srv := http.Server{
		Handler:      nil,
		Addr:         addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("登录错误"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"Message": "登录成功",
	})
}
