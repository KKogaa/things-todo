package repositories

import (
	"log"
	"time"

	"github.com/KKogaa/things-todo-backend/core/entity"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return TaskRepository{
		db: db,
	}

}

type BaseModel struct {
	ID uint `gorm:"primaryKey" json:"id"`
	//triple index, no need ?
	CreatedAt time.Time      `gorm:"index;not null" json:"-"`
	UpdatedAt time.Time      `gorm:"index;not null" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Task struct {
	BaseModel
	Title       string
	Description string
	Duration    int8
	Priority    int8
	Difficulty  int8
}

func entityToModel(task entity.Task) Task {

	return Task{
		Title:       task.Title,
		Description: task.Description,
		Duration:    int8(task.Duration),
		Priority:    int8(task.Priority),
		Difficulty:  int8(task.Difficulty),
		BaseModel:   BaseModel{ID: task.Id},
	}

}

func modelToEntity(task Task) entity.Task {
	return entity.Task{
		Title:       task.Title,
		Description: task.Description,
		Priority:    entity.Priority(task.Priority),
		Duration:    time.Duration(task.Duration),
		Difficulty:  entity.Difficulty(task.Difficulty),
		Id:          task.BaseModel.ID,
	}
}

func (t TaskRepository) GetAllTasks() ([]entity.Task, error) {
	var tasks []Task
	if err := t.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	var taskEntities []entity.Task
	for _, task := range tasks {
		taskEntities = append(taskEntities, modelToEntity(task))
	}
	return taskEntities, nil
}

func (t TaskRepository) CreateTask(task entity.Task) (entity.Task, error) {

	taskModel := entityToModel(task)

	if err := t.db.Create(&taskModel).Error; err != nil {
		return entity.Task{}, err
	}
	return modelToEntity(taskModel), nil
}

func (t TaskRepository) GetTask(taskId uint) (entity.Task, error) {
	var task Task
	if err := t.db.Find(&task, taskId).Error; err != nil {
		return entity.Task{}, err
	}
	return modelToEntity(task), nil

}

func (t TaskRepository) UpdateTask(task entity.Task) (entity.Task, error) {
	taskModel := entityToModel(task)
	if err := t.db.Save(&task).Error; err != nil {
		return entity.Task{}, err
	}
	return modelToEntity(taskModel), nil
}

func (t TaskRepository) DeleteTask(taskId uint) error {
	if err := t.db.Where("id = ?", taskId).Delete(&entity.Task{}).Error; err != nil {
		return err
	}

	return nil
}
