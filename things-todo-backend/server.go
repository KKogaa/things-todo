package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KKogaa/things-todo-backend/infra/db"
	"github.com/KKogaa/things-todo-backend/infra/handlers"
	"github.com/KKogaa/things-todo-backend/infra/repositories"
	"github.com/KKogaa/things-todo-backend/services"
	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func Wire(server *gin.Engine) {
	db, err := db.Init()
	if err != nil {
		log.Fatalf("error establishing db connection: %s", err)

	}

	// TODO: change to uber fx dependency injection
	taskRepo := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandlers(taskService)

	taskGroup := server.Group("/tasks")
	taskGroup.GET("", taskHandler.HandleListTasks)
	taskGroup.GET("/:taskId", taskHandler.HandleGetTask)
	taskGroup.POST("", taskHandler.HandleCreateTask)
	taskGroup.PUT("/:taskId", taskHandler.HandleUpdateTask)
	taskGroup.DELETE("/:taskId", taskHandler.HandleDeleteTask)

}

func (s *Server) Start() {
	instance := gin.New()
	instance.Use(gin.Recovery())
	Wire(instance)

	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: instance,
	}

	log.Printf("server started on port: %s\n ", "8080")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
