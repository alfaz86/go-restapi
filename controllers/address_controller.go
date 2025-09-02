package controllers

import (
	"go-restapi/models"
	"go-restapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {
	var address models.Address

	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateAddress(&address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create address"})
		return
	}

	c.JSON(http.StatusOK, address)
}

func GetAddresses(c *gin.Context) {
	addresses, err := services.GetAddresses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get addresses"})
		return
	}

	c.JSON(http.StatusOK, addresses)
}
