package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mannuR22/go-task-backend.git/services"
)

//CRUD
func main() {
	router := gin.Default()

	//API for creating List and provide reference/refId in response
	router.POST("/create/list", services.CreateList)

	//API for reading specifice list using unique reference/refId
	router.GET("/read/list/:refId", services.ReadList)

	//API for updating/adding more element into list, also it prevents
	//addition of dublicate items in list
	router.POST("/update/list/:refId", services.UpdateListItems)

	//API for deleting items from list
	router.POST("/delete/list/:refId", services.DeleteListItems)

	//API to print hello world
	router.GET("/hello", services.PrintHelloWorld)

	router.Run(":8080")
}
