package routes

import (
	"go-restapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Uploads
	r.Static("/uploads", "./uploads")

	// Users
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)

	// Addresses
	r.POST("/addresses", controllers.CreateAddress)
	r.GET("/addresses", controllers.GetAddresses)

	// Images
	r.POST("/images", controllers.CreateImage)
	r.GET("/images", controllers.GetImages)
}
