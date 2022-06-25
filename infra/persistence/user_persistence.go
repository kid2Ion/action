package persistence

import (
	"action/domain/model"
	"action/domain/repository"
	"database/sql"
	"log"
)

type userPersistence struct {
	conn *sql.DB
}

func NewUserPersistence(conn *sql.DB) repository.UserRepository {
	return &userPersistence{
		conn: conn,
	}
}

func (up userPersistence) Insert(u *model.Users) {
	_, err := up.conn.Exec(
		"INSERT INTO users (name, gender, room_id) VALUES (?, ?, ?)",
		u.Name,
		u.Gender,
		u.RoomId,
	)
	if err != nil {
		log.Fatal("faled to insert users")
	}
}

func (up userPersistence) GetAllUsersByRoomId(roomId string) ([]*model.Users, error) {
	// 中身まだ
	return nil, nil
}
