package persistence

import (
	"action/domain/model"
	"action/domain/repository"

	"github.com/jinzhu/gorm"
)

type userPersistence struct {
	Conn *gorm.DB
}

func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

func (up userPersistence) Insert(name string, gender int) error {
	// 中身はまだ
	return nil
}

func (up userPersistence) GetAllUsersByRoomId(roomId string) ([]*model.User, error) {
	// 中身まだ
	return nil, nil
}
