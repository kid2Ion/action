package handler

import (
	"action/usecase"
	util "action/utility"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	UserCreate(c echo.Context) error
	Action(c echo.Context) error
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) UserCreate(c echo.Context) error {

	name := c.QueryParam("name")
	gender, err := strconv.Atoi(c.QueryParam("gender"))
	if err != nil {
		return err
	}
	room_id := c.QueryParam("room_id")
	// 一人目だけランダムroom_id生成(2人目からはフロントから受け取る)
	if room_id == "" {
		room_id, err = util.RandomString(10)
		if err != nil {
			return err
		}
	}

	if err := uh.userUseCase.Insert(name, gender, room_id); err != nil {
		return err
	}

	return c.String(http.StatusOK, "create user")
}

func (uh userHandler) Action(c echo.Context) error {
	// param取得
	// usecaseの呼び出し
	// 値の返却
	return c.String(http.StatusOK, "return action")
}
