package controllers

import (
	"gorm-test/database"
	"gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct{}

func (EmployeeController) GetAllEmployees(c *gin.Context) {
	db := database.Get()
	var employees []models.Employee
	if err := db.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

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

func (EmployeeController) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		BadRequestError(c, err)
		return
	}
	db := database.Get()

	// employee.ID = 0
	// if err := employee.Create(db, &employee); err != nil {
	// 	InternalServerError(c, err)
	// 	return
	// }
	if err := db.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, employee)
}
