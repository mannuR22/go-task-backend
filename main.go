package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mannuR22/go-task-backend.git/services"
)

func main() {
	router := gin.Default()

	router.GET("/hello", services.PrintHelloWorld)

	router.Run(":8080")
}
