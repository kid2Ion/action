package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func Action(c echo.Context) error {
	return c.String(http.StatusOK, "return action")
}

func Users(c echo.Context) error {
	return c.String(http.StatusOK, "create user")
}
