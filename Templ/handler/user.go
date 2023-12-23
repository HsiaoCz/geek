package handler

import (
	"net/http"

	"github.com/HsiaoCz/geek/Templ/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (u UserHandler) UserLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Message": "Login successed",
		"Code":    10000,
	})
}

func (u UserHandler) UserSignup(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Message": "Signup successed",
		"Code":    10000,
	})
}

func (u UserHandler) Show(c echo.Context) error {
	return render(c, user.Show())
}
