package repository

import (
	"action/domain/model"
)

type UserRepository interface {
	GetAllUsersByRoomId(roomId string) ([]*model.User, error)
	Insert(name string, gender int) error
}
