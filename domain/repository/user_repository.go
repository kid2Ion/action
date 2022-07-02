package repository

import (
	"action/domain/model"
)

type UserRepository interface {
	GetAllUsersByRoomId(roomId string) ([]*model.Users, error)
	Insert(*model.Users)
}
