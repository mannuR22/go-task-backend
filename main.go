package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mannuR22/go-task-backend.git/services"
)

//CRUD
func main() {
	router := gin.Default()

	router.POST("/create/list")
	router.GET("/read/list")
	router.POST("/update/list")
	router.POST("/delete/list")
	router.GET("/hello", services.PrintHelloWorld)

	router.Run(":8080")
}
