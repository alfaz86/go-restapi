package services

import (
	"go-restapi/config"
	"go-restapi/database"
	"go-restapi/models"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Preload("Address").Preload("Image").Find(&users).Error

	for i := range users {
		for j := range users[i].Image {
			users[i].Image[j].URL = config.Cfg.URL + users[i].Image[j].Path
		}
	}

	return users, err
}
