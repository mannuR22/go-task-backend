package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mannuR22/go-task-backend.git/business"
)

func PrintHelloWorld(c *gin.Context) {

	response := business.HelloMessage()
	c.JSON(http.StatusOK, response)
}
