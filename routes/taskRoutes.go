package routes

import (
	"task-manager-api/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
    router.GET("/tasks", controllers.GetTasks)
    router.POST("/tasks", controllers.CreateTask)
    router.PUT("/tasks/:id", controllers.UpdateTask)
    router.DELETE("/tasks/:id", controllers.DeleteTask)
}