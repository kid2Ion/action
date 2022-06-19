package main

import (
	handler "action/handler"
	"action/infra/persistence"
	"action/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// 依存関係の注入
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	UserHandler := handler.NewUserHandler(userUseCase)

	e.GET("/action", UserHandler.Action)
	e.POST("/users", UserHandler.UserCreate)
	e.Logger.Fatal(e.Start(":8080"))
}
