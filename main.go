package main

import (
	"task-manager-api/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	r:=gin.Default()

	routes.TaskRoutes(r)
	
	r.Run()
}