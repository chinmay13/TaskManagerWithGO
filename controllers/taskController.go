package controllers

import (
	"strconv"
	"task-manager-api/models"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{
    {ID: 1, Title: "Learn Go", Status: "pending"},
    {ID: 2, Title: "Build API", Status: "pending"},
}

func GetTasks(c *gin.Context) {
    c.JSON(200, tasks)
}

func CreateTask(c *gin.Context) {
    var newTask models.Task

    if err := c.BindJSON(&newTask); err != nil {
        return
    }

    newTask.ID = len(tasks) + 1
    tasks = append(tasks, newTask)
    c.JSON(201, newTask)
}

func UpdateTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var updatedTask models.Task

    for i, task := range tasks {
        if task.ID == id {
            if err := c.BindJSON(&updatedTask); err != nil {
                return
            }
            tasks[i] = updatedTask
            c.JSON(200, updatedTask)
            return
        }
    }

    c.JSON(404, gin.H{"message": "Task not found"})
}

func DeleteTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            c.JSON(200, gin.H{"message": "Task deleted"})
            return
        }
    }

    c.JSON(404, gin.H{"message": "Task not found"})
}