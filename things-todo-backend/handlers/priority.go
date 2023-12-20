package handlers

import (
	"net/http"

	"github.com/KKogaa/things-todo-backend/data"
	"github.com/gin-gonic/gin"
)

func HandleListPriorities(c *gin.Context) {
	priorities, err := data.GetAllPriorities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, priorities)
}
