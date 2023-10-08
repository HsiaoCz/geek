package api

import "github.com/gofiber/fiber/v2"

type server struct {
	addr string
	u    *user
}

func NewServer(addr string) *server {
	return &server{
		addr: addr,
		u:    newUser(),
	}
}

func (s *server) Start() error {
	r := fiber.New()
	r.Post("/api/v1/user/register", s.u.handleUserRegister)
	return r.Listen(s.addr)
}
