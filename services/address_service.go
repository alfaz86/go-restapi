package services

import (
	"go-restapi/database"
	"go-restapi/models"
)

func CreateAddress(addr *models.Address) error {
	return database.DB.Create(addr).Error
}

func GetAddresses() ([]models.Address, error) {
	var addresses []models.Address
	err := database.DB.Find(&addresses).Error
	return addresses, err
}
