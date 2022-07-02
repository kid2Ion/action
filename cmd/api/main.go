package main

import (
	db "action/config"
	room_handler "action/handler/room"
	user_handler "action/handler/user"
	"action/infra/persistence"
	"action/usecase"

	"github.com/labstack/echo/v4"
)

func main() {

	conn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	e := echo.New()
	// 依存関係の注入(全層のインスタンス化を行う)
	userPersistence := persistence.NewUserPersistence(conn)
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := user_handler.NewUserHandler(userUseCase)

	e.GET("/action", userHandler.Action)
	e.GET("/room", room_handler.GetRoomId)
	e.POST("/users", userHandler.CreateUser)
	e.Logger.Fatal(e.Start(":8080"))
}
