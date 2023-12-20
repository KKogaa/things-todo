package handlers

import (
	"net/http"

	"github.com/KKogaa/things-todo-backend/data"
	"github.com/gin-gonic/gin"
)

func HandleListTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func HandleCreateTask(c *gin.Context) {
	var request data.CreateTaskParams

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := data.CreateTask(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func HandleUpdateTask(c *gin.Context) {
	taskId := c.Param("taskId")
	var request data.CreateTaskParams

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := data.UpdateTask(taskId, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func HandleDeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")

	if err := data.DeleteTask(taskId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
