package api

import "github.com/gofiber/fiber/v2"

type user struct {
}

func newUser() *user {
	return &user{}
}

func (u *user) handleUserRegister(c *fiber.Ctx) error {
	return nil
}
