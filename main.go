package main

import (
	"fmt"
	"go-restapi/config"
	"go-restapi/database"
	"go-restapi/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config.LoadConfig()

	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Connect DB
	database.Connect()

	// Routes
	routes.SetupRoutes(r)

	// Run server
	fmt.Println("Server running at", config.Cfg.Addr())
	r.Run(config.Cfg.Addr())
}
