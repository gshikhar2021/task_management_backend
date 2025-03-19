package main

import (
	"backend/config"
	"backend/routes"
	"backend/websockets"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"}, // Add your frontend URLs here
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.SetupRoutes(router)
	routes.TaskRoutes(router)

	// WebSocket route
	router.GET("/ws", websockets.HandleConnections)

	router.Run(":8080")
}
