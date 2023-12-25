package db

import (
	"fmt"

	"github.com/KKogaa/things-todo-backend/infra/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&repositories.Task{}); err != nil {
		return err
	}

	return nil
}

func connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}
