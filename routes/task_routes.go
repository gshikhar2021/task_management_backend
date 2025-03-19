package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	r.POST("/tasks/create", middleware.AuthMiddleware(), controllers.CreateTask)
	r.GET("/tasks", middleware.AuthMiddleware(), controllers.GetTasks)
	r.PUT("/tasks/:id/done", middleware.AuthMiddleware(), controllers.MarkTaskAsDone)
}
