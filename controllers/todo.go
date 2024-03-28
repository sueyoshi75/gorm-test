package controllers

import (
	"gorm-test/database"
	"gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {}

func (TodoController) CreateTodos(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	db := database.Get()
	if err := db.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}