package main

import (
	db "action/config"
	"action/domain/model"
	handler "action/handler"
	"action/infra/persistence"
	"action/usecase"

	"github.com/labstack/echo/v4"
)

func main() {

	conn, err := db.ConnectDB()
	if err != nil {
		panic("failes to connect db")
	}
	defer conn.Close()

	conn.AutoMigrate(&model.User{})

	e := echo.New()
	// 依存関係の注入(全層のインスタンス化を行う)
	userPersistence := persistence.NewUserPersistence(conn)
	userUseCase := usecase.NewUserUseCase(userPersistence)
	UserHandler := handler.NewUserHandler(userUseCase)

	e.GET("/action", UserHandler.Action)
	e.POST("/users", UserHandler.UserCreate)
	e.Logger.Fatal(e.Start(":8080"))
}
