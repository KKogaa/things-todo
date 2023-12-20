package main

import (
	"fmt"
	"log"

	"github.com/KKogaa/things-todo-backend/db"
	"github.com/KKogaa/things-todo-backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Init server...")
	r := gin.Default()

	fmt.Println("Connecting to db...")
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/priorities", handlers.HandleListPriorities)

	tasks := r.Group("/tasks")
	tasks.GET("/", handlers.HandleListTasks)
	tasks.POST("/", handlers.HandleCreateTask)
	tasks.PUT("/:taskId", handlers.HandleUpdateTask)
	tasks.DELETE("/:taskId", handlers.HandleDeleteTask)

	r.Run()
}
