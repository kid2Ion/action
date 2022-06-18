package main

import (
	rest "action/handler/rest"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", rest.Index)
	e.GET("/action", rest.Action)
	e.POST("/users", rest.Users)
	e.Logger.Fatal(e.Start(":8080"))
}
