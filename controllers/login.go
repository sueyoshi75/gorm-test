package controllers

import (
	"gorm-test/database"
	"gorm-test/models"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {}

func (LoginController) Login(c *gin.Context) {
	var user models.User
	db := database.Get()
	if err := c.ShouldBindJSON(&user); err != nil {
		BadRequestError(c, err)
		return
	}
	if err := user.Verify(db); err != nil {
		UnauthorizedError(c, err)
		return
	}

	session := sessions.Default(c)
	session.Set("email", user.Email)
	max_age := 60 * 60 * 2
	session.Options(sessions.Options{MaxAge: max_age})
	session.Save()

	url := os.Getenv("UI_URL")
	c.Redirect(http.StatusFound, url+"todos")
}

func (LoginController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: 0})
	session.Delete("email")
	if err := session.Save(); err != nil {
		InternalServerError(c, err)
		return
	}
	url := os.Getenv("UI_URL")
	c.Redirect(http.StatusFound, url+"employees/new")
}