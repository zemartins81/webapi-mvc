package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jcmartins81/webapi-with-go/database"
	"github.com/jcmartins81/webapi-with-go/models"
)

func CreateUser(c *gin.Context) {
	db := database.GetDataBase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
}
