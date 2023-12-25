package main

import (
	"log"

	"github.com/KKogaa/things-todo-backend/infra/data"
	"github.com/KKogaa/things-todo-backend/infra/db"
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
