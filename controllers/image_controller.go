package controllers

import (
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"go-restapi/models"
	"go-restapi/services"
)

func CreateImage(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(file.Filename)
	path := "uploads/" + filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	image := models.Image{
		UserID: uint(userID),
		Path:   path,
	}

	if err := services.CreateImage(&image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	c.JSON(http.StatusOK, image)
}

func GetImages(c *gin.Context) {
	images, err := services.GetImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch images"})
		return
	}
	c.JSON(http.StatusOK, images)
}
