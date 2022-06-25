package model

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Users struct {
	Id        int    `json:"id"`
	Name      string `json:"name" validate:"required,gte=1,lte=10,alphanum"`
	Gender    int    `json:"gender" validate:"required,gte=0,lte=1"`
	RoomId    string `json:"room_id" validate:"required,gte=10,lte=10"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
