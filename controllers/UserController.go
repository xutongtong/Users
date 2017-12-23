package controllers

import (
	"Users/forms"
	"Users/services"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var server = new(services.UserService)

// User Sign
func (uc UserController) Login(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	userLoginForm := forms.UserLoginForm{Name:name, Password:password}
	server.Login(userLoginForm)
}

// User Signup
//func (uc UserController) Signup(c *gin.Context) {
//
//}

// User Signup
//func (uc UserController) Signout(c *gin.Context) {
//
//}

// User ForgetPassword
//func (uc UserController) ForgetPassword(c *gin.Context) {
//
//	return;
//}

// User ResetPassword
//func (uc UserController) ResetPassword(c *gin.Context) {
//
//	return;
//}

// User Bind Mobile
//func (uc UserController) BindMobile(c *gin.Context) {
//
//	return;
//}