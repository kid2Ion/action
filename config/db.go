package db

import "github.com/jinzhu/gorm"

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
