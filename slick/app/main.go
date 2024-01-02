package main

import "github.com/HsiaoCz/geek/slick"

func main() {
	app := slick.New()
	app.Get("/user/profile", HandleUserProfile)
}

func HandleUserProfile(c *slick.Context) error {
	return nil
}
