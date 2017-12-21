package apis

import (
	"Users/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/login", user.Login)
			//userGroup.POST("/register", user.Signup)
			//userGroup.POST("/logout", user.Signout)
			//userGroup.POST("/forgetPassword", user.ForgetPassword)
			//userGroup.POST("/resetPassword", user.ResetPassword)
			//userGroup.POST("/bindMobile", user.BindMobile)
		}
	}

	return router
}