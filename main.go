package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	db := dbinit()
	db.AutoMigrate(&User{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		result, err := getAll(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	})
	r.Run(":8080")

	// insert(db)
	// getOne(db)
	// getAll(db)
	// updates(db)
	// delete(db)
}

//DB接続
func dbinit() *gorm.DB {
	dsn := "root@tcp(db:3306)/gorm-test?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
    panic("failed to connect database")
  }
	return db
}

//データ作成
func insert(db *gorm.DB) {
	user := User{
		Name:     "TeTe",
		Age:      40,
	}
	result := db.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Made:", result.RowsAffected)
	fmt.Println("user:", user)
}

//データ単体取得
func getOne(db *gorm.DB) (User, error){
	user := User{}
	result := db.First(&user, "id = ?", 1)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Get:", result.RowsAffected)
	fmt.Println("user:", user)

	return user, nil
}

//データ全件取得
func getAll(db *gorm.DB) ([]User, error) {
	users := []User{}
	result := db.Find(&users)
	
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Get:", result.RowsAffected)
	fmt.Println("users:", users)

	return users, nil
}

//データ更新
func updates(db *gorm.DB) {
	result := db.Model(&User{}).Where("id = 1").Updates(User{Name: "Tama", Age: 10})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Update:", result.RowsAffected)

	user := User{}
	db.Where("id = 1").Take(&user)
	fmt.Println("user:", user)
}

//データ削除
func delete(db *gorm.DB) {
	db.Where("id = 7").Delete(&User{})
	fmt.Println("Delete: 1")
}