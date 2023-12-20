package data

import (
	"time"

	"github.com/KKogaa/things-todo-backend/db"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"index;not null" json:"-"`
	UpdatedAt time.Time      `gorm:"index;not null" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Task struct {
	BaseModel
	Title       string
	Description string
	Duration    int8
	Value       int8
}

type CreateTaskParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int8   `json:"duration"`
	Value       int8   `json:"value"`
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	if err := db.Gorm.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func CreateTask(params CreateTaskParams) (*Task, error) {
	task := Task{
		Title:       params.Title,
		Description: params.Description,
		Duration:    params.Duration,
		Value:       params.Value,
	}
	if err := db.Gorm.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(taskId string, params CreateTaskParams) (*Task, error) {
	var task Task
	if err := db.Gorm.Find(&task, taskId).Error; err != nil {
		return nil, err
	}
	if err := db.Gorm.Model(&task).Updates(params).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func DeleteTask(taskId string) error {
	var task Task
	if err := db.Gorm.Find(&task, taskId).Error; err != nil {
		return err
	}
	if err := db.Gorm.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}
