package controllers

import (
	"fmt"

	//"Users/models"
	"Users/services"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var server = new(services.UserService)

// User Sign
func (uc UserController) Login(c *gin.Context) {
	fmt.Println("Start")

	name1, flag := c.Get("name")
	if flag {
		fmt.Println("not exist")
	}
	name := c.Param("name")
	password := c.Param("password")

	fmt.Println(name1)
	fmt.Printf("Name:%s", name)
	fmt.Println(password)

	//newUser := models.User{Name:name,Password:password}
	//
	//fmt.Println(newUser)
	//
	//server.Login(newUser)
	//
	//fmt.Println("End")
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