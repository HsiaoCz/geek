package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	name := c.Param("username")
	return c.JSON(http.StatusOK, echo.Map{
		"Message": "Everything is ok !",
		"Data":    name,
	})
}
