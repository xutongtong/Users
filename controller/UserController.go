package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/xutongtong/Users/model"
)

// User Sign
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, model.UserAdd())
}

func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, model.UserList())
}