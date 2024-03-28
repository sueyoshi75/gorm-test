package database

import (
	"fmt"
	"gorm-test/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() {
	godotenv.Load()
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	dbname := os.Getenv("DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Employee{}, &models.Todo{})
	DB = db
}

// func OpenTestDB() {
// 	godotenv.Load("/var/www/todo/.env")
// 	user := os.Getenv("MYSQL_ROOT")
// 	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
// 	host := os.Getenv("DB_HOST")
// 	test_db := os.Getenv("MYSQL_TEST_DATABASE")
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?checkConnLiveness=false&maxAllowedPacket=0&parseTime=true&loc=Asia%%2FTokyo", user, pass, host, test_db)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	DB = db
// }

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