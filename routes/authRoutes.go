package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/home", middleware.AuthMiddleware(), controllers.Home)
}
