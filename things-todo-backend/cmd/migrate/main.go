package main

import (
	"log"

	"github.com/KKogaa/things-todo-backend/data"
	"github.com/KKogaa/things-todo-backend/db"
)

func migrate() error {
	return db.Gorm.AutoMigrate(
		&data.Priority{},
		&data.Task{},
	)
}

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	if err := migrate(); err != nil {
		log.Fatal(err)
	}
}
