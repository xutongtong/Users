package controllers

import (
	"../forms"
	"../services"

	"github.com/gin-gonic/gin"
	"fmt"
)

type UserController struct{}

var userService = new(services.UserService)

// User Sign
func (uc UserController) Signin(c *gin.Context) {
	input.Validate(c.Params)

	userService.Signin(form)

	fmt.Println("ok")
	c.JSON(200, gin.H{"message": "Successed!"})
}

// User Signup
func (uc UserController) Signup(c *gin.Context) {

}

// User Signup
func (uc UserController) Signout(c *gin.Context) {

}

// User ForgetPassword
func (uc UserController) ForgetPassword(c *gin.Context) {

	return;
}

// User ResetPassword
func (uc UserController) ResetPassword(c *gin.Context) {

	return;
}

// User Bind Mobile
func (uc UserController) BindMobile(c *gin.Context) {

	return;
}