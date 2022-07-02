package user_handler

import (
	"action/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
	GetAction(c echo.Context) error
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) CreateUser(c echo.Context) error {

	name := c.QueryParam("name")
	gender, err := strconv.Atoi(c.QueryParam("gender"))
	if err != nil {
		return err
	}
	roomId := c.QueryParam("room_id")

	if err := uh.userUseCase.Insert(name, gender, roomId); err != nil {
		return err
	}

	return c.String(http.StatusOK, "success create user")
}

func (uh userHandler) GetAction(c echo.Context) error {
	roomId := c.QueryParam("room_id")
	action, err := uh.userUseCase.GenerateAction(roomId)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, action)
}
