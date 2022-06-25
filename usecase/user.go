package usecase

import (
	"action/domain/model"
	"action/domain/repository"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserUseCase interface {
	Insert(name string, gender int, room_id string) error
	GenarateAction(roomId string) ([]*model.Users, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) Insert(name string, gender int, room_id string) error {

	var u model.Users
	u.Name = name
	u.Gender = gender
	u.RoomId = room_id

	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		msg := fmt.Sprintf("validation err :%s", err)
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
	}
	uu.userRepository.Insert(&u)
	return nil
}

func (uu userUseCase) GenarateAction(roomId string) ([]*model.Users, error) {
	// action生成ロジック
	// vlidation
	// domainの呼び出し
	return nil, nil
}
