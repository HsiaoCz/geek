package main

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/geek/anymore"
)

type UserR struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

func main() {
	r := anymore.App()
	r.GET("/user/register", UserRegister)
	r.Listen("127.0.0.1:9091")
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	userR := new(UserR)
	if err := json.NewDecoder(r.Body).Decode(userR); err != nil {
		return 
	}
}
