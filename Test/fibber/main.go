package main

import "github.com/gofiber/fiber/v2"

const (
	addr = "127.0.0.1:9911"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Identity int    `json:"identity"`
}

func main() {
	r := fiber.New()
	r.Get("/user", GetUser)
	r.Listen(addr)
}

func GetUser(c *fiber.Ctx) error {
	user := User{
		Username: "bob",
		Password: "12222",
		Identity: 1222111,
	}
	
	return c.JSON(user)
}
