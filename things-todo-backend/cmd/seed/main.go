package main

import (
	"log"

	"github.com/KKogaa/things-todo-backend/infra/data"
	"github.com/KKogaa/things-todo-backend/infra/db"
)

func migratePriority() error {
	priorities := []data.Priority{
		{Title: "Low"},
		{Title: "Medium"},
		{Title: "High"},
	}
	if err := db.Gorm.Create(&priorities).Error; err != nil {
		return err
	}
	return nil
}

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	if err := migratePriority(); err != nil {
		log.Fatal(err)
	}
}
