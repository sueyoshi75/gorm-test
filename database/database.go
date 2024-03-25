package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() {
	godotenv.Load()
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/todo?checkConnLiveness=false&maxAllowedPacket=0&parseTime=true&loc=Asia%%2FTokyo", user, pass, host)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func OpenTestDB() {
	godotenv.Load("/var/www/todo/.env")
	user := os.Getenv("MYSQL_ROOT")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	host := os.Getenv("DB_HOST")
	test_db := os.Getenv("MYSQL_TEST_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?checkConnLiveness=false&maxAllowedPacket=0&parseTime=true&loc=Asia%%2FTokyo", user, pass, host, test_db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		panic(err)
	}
	db.Close()
}

func Get() *gorm.DB {
	return DB
}