package services

import (
	"go-restapi/config"
	"go-restapi/database"
	"go-restapi/models"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

func CreateImage(image *models.Image) error {
	return database.DB.Create(image).Error
}

func GetImages() ([]models.Image, error) {
	var images []models.Image
	if err := database.DB.Find(&images).Error; err != nil {
		return nil, err
	}

	for i := range images {
		images[i].URL = config.Cfg.URL + images[i].Path
	}

	return images, nil
}
