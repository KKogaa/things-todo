package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/KKogaa/things-todo-backend/infra/handlers/dtos"
	"github.com/KKogaa/things-todo-backend/services"
	"github.com/gin-gonic/gin"
)

type TaskHandlers struct {
	taskService services.TaskService
}

func NewTaskHandlers(taskService services.TaskService) TaskHandlers {
	return TaskHandlers{
		taskService: taskService,
	}
}

func (t TaskHandlers) ExtractId(c *gin.Context) (uint, error) {
	taskIdStr := c.Param("taskId")
	taskId, err := strconv.ParseUint(taskIdStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to conver to uint: %s", err)
	}
	return uint(taskId), nil

}

func (t TaskHandlers) HandleListTasks(c *gin.Context) {
	tasks, err := t.taskService.GetAllTasks()
	if err != nil {
		//TODO: also add log handle the errors, erros need to persist
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (t TaskHandlers) HandleGetTask(c *gin.Context) {
	taskId, err := t.ExtractId(c)
	if err != nil {
		log.Fatal(err)
	}

	task, err := t.taskService.GetTask(taskId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (t TaskHandlers) HandleCreateTask(c *gin.Context) {
	var request dtos.CreateTaskDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := t.taskService.CreateTask(request.Title,
		request.Description, request.Duration, request.Priority, request.Difficulty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t TaskHandlers) HandleUpdateTask(c *gin.Context) {
	taskId, err := t.ExtractId(c)
	if err != nil {
		log.Fatal(err)
	}
	var request dtos.CreateTaskDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := t.taskService.UpdateTask(taskId, request.Title,
		request.Description, request.Duration, request.Priority, request.Difficulty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t TaskHandlers) HandleDeleteTask(c *gin.Context) {
	taskId, err := t.ExtractId(c)
	if err != nil {
		log.Fatal(err)
	}

	if err := t.taskService.DeleteTask(taskId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
