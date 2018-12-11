package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Get(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
	return
}
