package usecase

import (
	"action/domain/model"
	"action/domain/repository"
	util "action/utility"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserUseCase interface {
	Insert(name string, gender int, roomId string) error
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

func (uu userUseCase) Insert(name string, gender int, roomId string) error {

	var u model.Users
	u.Name = name
	u.Gender = gender
	// 一人目だけランダムroom_id生成(2人目からはフロントから受け取る)
	if roomId == "" {
		randomRoomId, err := util.RandomString(10)
		if err != nil {
			return err
		}
		u.RoomId = randomRoomId
	}

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
