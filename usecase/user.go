package usecase

import (
	"action/domain/model"
	"action/domain/repository"

	"github.com/jinzhu/gorm"
)

type UserUseCase interface {
	Insert(conn *gorm.DB, name string, gender int) error
	GenarateAction(conn *gorm.DB, roomId string) ([]*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) Insert(conn *gorm.DB, name string, gender int) error {
	// バリデーション
	// domainの呼び出し
	return nil
}

func (uu userUseCase) GenarateAction(conn *gorm.DB, roomId string) ([]*model.User, error) {
	// action生成ロジック
	// domainの呼び出し
	return nil, nil
}
