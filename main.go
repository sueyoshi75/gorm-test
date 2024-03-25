package main

import (
	"gorm-test/database"
	"gorm-test/routes"
)

func main() {
	database.Open()
	defer database.Close()

	if err := routes.Run(); err != nil {
		panic(err)
	}
}