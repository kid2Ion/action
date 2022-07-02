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
	GenerateAction(roomId string) (string, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

type userGroup struct {
	*model.Users
	UserGroup int
}

func (uu userUseCase) Insert(name string, gender int, roomId string) error {

	u := model.NewUser(name, gender, roomId)

	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		msg := fmt.Sprintf("user struct validation err :%s", err)
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
	}
	uu.userRepository.Insert(u)
	return nil
}

func (uu userUseCase) GenerateAction(roomId string) (string, error) {
	users := uu.userRepository.GetAllUsersByRoomId(roomId)
	action := getActionContent(users)

	return action, nil
}

func getActionContent(users []*model.Users) string {
	var grouA []*userGroup
	var grouB []*userGroup

	// 埋め込み構造体の練習
	for _, v := range users {
		u := &userGroup{UserGroup: v.Gender, Users: v}
		if u.UserGroup == 1 {
			grouA = append(grouA, u)
		} else {
			grouB = append(grouB, u)
		}
	}

	actionContent := fmt.Sprintf("%s と %s が ~~~", grouA[util.RandomInt(len(grouA))].Name, grouB[util.RandomInt(len(grouB))].Name)
	// todo contentの内容の出しわけ
	return actionContent
}
