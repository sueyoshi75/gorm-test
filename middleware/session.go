package middleware

import (
	"gorm-test/database"
	"os"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Session() gin.HandlerFunc {
	db := database.Get()
	godotenv.Load()
	secret := os.Getenv("SECRET")
	store := gormsessions.NewStore(db, true, []byte(secret))
	return sessions.Sessions("sessions", store)
}