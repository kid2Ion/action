package handler

import (
	"action/usecase"
	"net/http"

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
	// param取得
	// usecaseの呼び出し
	return c.String(http.StatusOK, "create user")
}

func (uh userHandler) Action(c echo.Context) error {
	// param取得
	// usecaseの呼び出し
	// 値の返却
	return c.String(http.StatusOK, "return action")
}
