package ports

import "github.com/KKogaa/things-todo-backend/core/entity"

type TaskRepository interface {
	GetAllTasks() ([]entity.Task, error)

	CreateTask(task entity.Task) (entity.Task, error)

	UpdateTask(task entity.Task) (entity.Task, error)

	DeleteTask(taskId uint) error

	GetTask(taskId uint) (entity.Task, error)
}
