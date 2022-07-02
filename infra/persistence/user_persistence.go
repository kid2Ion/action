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
	// todo roomIdの返却
	if err != nil {
		log.Fatal("faled to insert users")
	}
}

func (up userPersistence) GetAllUsersByRoomId(r string) []*model.Users {
	rows, err := up.conn.Query("SELECT * FROM users WHERE room_id = ?", r)
	if err != nil {
		log.Fatal("faled to get users")
	}
	users := []*model.Users{}

	for rows.Next() {
		user := &model.Users{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Gender, &user.RoomId, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Fatal("getRows rows.Scan error err:%v", err)
		}
		users = append(users, user)
	}
	return users
}
