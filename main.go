package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
}

func main() {
	db := dbinit()
	db.AutoMigrate(&Todo{})

	r := gin.Default()

	//corsの設定
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
				"http://localhost:5173",
		},
		AllowMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
		},
		AllowHeaders: []string{
				"Content-Type",
		},
		AllowCredentials: false,
		MaxAge: 24 * time.Hour,
	}))

	//全件表示
	r.GET("/todos", func(c *gin.Context) {
		todos := []Todo{}
		result := db.Find(&todos)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(http.StatusOK, todos)
	})

	//ユーザー作成
	r.POST("/todos", func(c *gin.Context) {
		var todo Todo
		c.BindJSON(&todo)
		result := db.Create(&todo)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(http.StatusOK, gin.H{"message": "todo created successfully"})
	})

	//データ検索
	r.GET("/todos/:id", func(c *gin.Context) {
		var todo Todo
		id := c.Param("id")
		result := db.First(&todo, id)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(http.StatusOK, todo)
	})

	//データ更新
	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo Todo
		c.BindJSON(&todo)
		result := db.Model(&Todo{}).Where("id = ?", id).Updates(&todo)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(http.StatusOK, gin.H{"message": "todo updated successfully"})
	})

	//データ削除
	r.DELETE("/todos/:id", func(c *gin.Context) {
		var todo Todo
		id := c.Param("id")
		result := db.Delete(&todo, id)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})
	})

	r.Run(":8080")
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
