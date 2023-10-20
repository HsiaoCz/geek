package main

import (
	"encoding/json"

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

func UserRegister(c *anymore.Context) {
	userR := new(UserR)
	if err := json.NewDecoder(c.R.Body).Decode(userR); err != nil {
		return
	}
}
