package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mannuR22/go-task-backend.git/business"
	models "github.com/mannuR22/go-task-backend.git/models"
)

func CreateList(c *gin.Context) {

	var userList models.ListModel
	if err := c.ShouldBindJSON(&userList); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.ListCreatedVM{})
		return
	}
	listVM, err := business.CreateList(userList)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, listVM)
}

func ReadList(c *gin.Context) {

	refId := c.Param("refId")

	userList, err := business.ReadList(refId)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, userList)
}

func UpdateListItems(c *gin.Context) {
	refId := c.Param("refId")
	var userListItems models.ListModel
	if err := c.ShouldBindJSON(&userListItems); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.ListCreatedVM{})
		return
	}
	err := business.UpdateListItems(refId, userListItems)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "Unique Items have been Added")
}

func DeleteListItems(c *gin.Context) {
	refId := c.Param("refId")
	var userListItems models.ListModel
	if err := c.ShouldBindJSON(&userListItems); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.ListCreatedVM{})
		return
	}
	err := business.DeleteListItems(refId, userListItems)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "Item(s) deleted successfully")
}
