package routes

import (
	"gorm-test/controllers"
	"gorm-test/middleware"
	"github.com/gin-gonic/gin"
)

var (
	employee_controller controllers.EmployeeController
	login_controller controllers.LoginController
)

func Run() error {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.Session())

	router.POST("/login", login_controller.Login)

	router.Use(middleware.Auth)

	router.GET("/logout", login_controller.Logout)
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/employees/employeeid/:employeeid", employee_controller.GetEmployees)
		}
	}
	return router.Run()
}