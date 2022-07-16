package usecase

import (
	"action/domain/model"
	"action/domain/repository"
	util "action/utility"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type Talk struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Alcohol struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Action struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type ActionContent struct {
	Talk    []Talk    `json:"talk"`
	Alcohol []Alcohol `json:"alcohol"`
	Action  []Action  `json:"action"`
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
	actionContent, err := getActionContentFromJsonFile()
	if err != nil {
		return "", err
	}

	var ac string
	var action string
	// actionのランダムなパターン
	n := util.RandomInt(3)
	switch n {
	case 0:
		ac = actionContent.Talk[util.RandomInt(len(actionContent.Talk))].Content
		action = getUserFromUsers(users) + ac
	case 1:
		ac = actionContent.Alcohol[util.RandomInt(len(actionContent.Alcohol))].Content
		action = getUserFromUsers(users) + ac
	case 2:
		ac = actionContent.Action[util.RandomInt(len(actionContent.Action))].Content
		action = getTwoUsersFromUsers(users) + ac
	}
	return action, nil
}

func getActionContentFromJsonFile() (ActionContent, error) {
	jsonFromFile, err := ioutil.ReadFile("./actionContent.json")
	if err != nil {
		log.Fatal(err)
	}
	var actionContent ActionContent
	err = json.Unmarshal(jsonFromFile, &actionContent)
	if err != nil {
		log.Fatal(err)
	}
	return actionContent, nil
}

func getTwoUsersFromUsers(users []*model.Users) string {
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

	twoUsers := fmt.Sprintf("%s と %s が", grouA[util.RandomInt(len(grouA))].Name, grouB[util.RandomInt(len(grouB))].Name)
	return twoUsers
}

func getUserFromUsers(users []*model.Users) string {
	user := fmt.Sprintf("%s が", *&users[util.RandomInt(len(users))].Name)
	return user
}
