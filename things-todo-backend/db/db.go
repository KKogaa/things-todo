package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func Init() error {
	db, err := connect()
	if err != nil {
		return err
	}
	Gorm = db
	return nil
}

func connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}
