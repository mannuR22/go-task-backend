package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/v1/users", getUsers)

	router.Run(":8080")
}

func getUsers(c *gin.Context) {
	// Handle GET request for fetching all users
}

func createUser(c *gin.Context) {
	// Handle POST request for creating a new user
}

func updateUser(c *gin.Context) {
	// Handle PUT request for updating an existing user
}

func deleteUser(c *gin.Context) {
	// Handle DELETE request for deleting an existing user
}
