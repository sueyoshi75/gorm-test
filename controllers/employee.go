package controllers

import (
	"gorm-test/database"
	"gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct{}

func (EmployeeController) GetEmployees(c *gin.Context) {
	employeeid := c.Param("employeeid")
	db := database.Get()
	var (
		employee models.Employee
	)

	employees, err := employee.GetEmployeesByEmployeeid(db, employeeid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[int]any{}
	for _, employee := range employees {
		response[int(employee.ID)] = employee
	}

	c.JSON(http.StatusOK, response)
}
