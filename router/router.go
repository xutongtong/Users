package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xutongtong/Users/controller"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		v1.GET("/login", controller.Login)
		v1.GET("/add", controller.AddUser)
		v1.GET("/all", controller.GetAllUser)
		//userGroup.POST("/register", user.Signup)
		//userGroup.POST("/logout", user.Signout)
		//userGroup.POST("/forgetPassword", user.ForgetPassword)
		//userGroup.POST("/resetPassword", user.ResetPassword)
		//userGroup.POST("/bindMobile", user.BindMobile)
	}

	router.Run(":8080")
	return router
}