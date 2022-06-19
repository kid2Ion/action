package repository

import (
	"action/domain/model"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetAllUsersByRoomId(conn *gorm.DB, roomId string) ([]*model.User, error)
	Insert(conn *gorm.DB, name string, gender int) error
}
