package algo

import "github.com/KKogaa/things-todo-backend/core/entity"

type TaskManager struct {
	Tasks []*entity.Task
}

func NewTaskManager(tasks []*entity.Task) TaskManager {
	return TaskManager{
		Tasks: tasks,
	}
}
