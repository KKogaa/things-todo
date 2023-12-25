package services

import (
	"time"

	"github.com/KKogaa/things-todo-backend/core/entity"
	"github.com/KKogaa/things-todo-backend/services/adapters"
)

type TaskService struct {
	taskRepo adapters.TaskRepository
}

func NewTaskService(taskRepo adapters.TaskRepository) TaskService {
	return TaskService{taskRepo: taskRepo}
}

func (t TaskService) GetAllTasks() ([]entity.Task, error) {
	return t.taskRepo.GetAllTasks()
}

func (t TaskService) GetTask(taskId uint) (entity.Task, error) {
	return t.taskRepo.GetTask(taskId)
}

func (t TaskService) CreateTask(title string, description string,
	duration int8, priority int8, difficulty int8) (entity.Task, error) {

	task := entity.Task{
		Title:       title,
		Description: description,
		Duration:    time.Duration(duration),
		Priority:    entity.Priority(priority),
		Difficulty:  entity.Difficulty(difficulty),
	}
	return t.taskRepo.CreateTask(task)

}

func (t TaskService) UpdateTask(taskId uint, title string, description string,
	duration int8, priority int8, difficulty int8) (entity.Task, error) {

	task := entity.Task{}
	return t.taskRepo.UpdateTask(task)
}

func (t TaskService) DeleteTask(taskId uint) error {
	return t.taskRepo.DeleteTask(taskId)
}
