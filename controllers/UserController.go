package controllers

import (
	//"fmt"

	"Users/forms"
	"Users/services"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var server = new(services.UserService)

// User Sign
func (uc UserController) Login(c *gin.Context) {
	//fmt.Println("Start")
	userLoginForm := forms.UserLoginForm{Name:"xutt", Password:"123456"}
	server.Login(userLoginForm)

	//name1, flag := c.Get("name")
	//if flag {
	//	fmt.Println("not exist")
	//}
	//name := c.Param("name")
	//password := c.Param("password")

	//fmt.Println(name1)
	//fmt.Printf("Name:%s", name)
	//fmt.Println(password)
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