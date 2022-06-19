package persistence

import (
	"action/domain/model"
	"action/domain/repository"

	"github.com/jinzhu/gorm"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) Insert(conn *gorm.DB, name string, gender int) error {
	// 中身はまだ
	return nil
}

func (up userPersistence) GetAllUsersByRoomId(conn *gorm.DB, roomId string) ([]*model.User, error) {
	// 中身まだ
	return nil, nil
}
