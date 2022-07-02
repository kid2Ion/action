package repository

import (
	"action/domain/model"
)

type UserRepository interface {
	GetAllUsersByRoomId(string) []*model.Users
	Insert(*model.Users)
}
