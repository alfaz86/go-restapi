package main

import (
	"fmt"
	"go-restapi/database"
	"go-restapi/models"
)

func main() {
	database.Connect()

	err := database.DB.AutoMigrate(&models.User{}, &models.Address{}, &models.Image{})
	if err != nil {
		panic("Failed to migrate database")
	}

	fmt.Println("✅ Migration success!")
}
