package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "config/example.db")
	if err != nil {
		return nil, err
	}

	tableName := "users"
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			gender INTEGER NOT NULL,
			room_id INTEGER,
			created_at TIMESTAMP DEFAULT (DATETIME('now','localtime')),
			updated_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))
			)`, tableName)
	_, err = db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
