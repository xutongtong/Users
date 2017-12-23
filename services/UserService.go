package services

import (
	"Users/db"
	//"Users/models"
	"Users/forms"
	"fmt"
	"Users/models"
)

type UserService struct {}

var orm = db.GetDB()

// User Signin
func (us UserService) Login(form forms.UserLoginForm) {
	fmt.Println("Start Login")
	user := models.User{Name:form.Name, Password:form.Password, Mobile:"18680663925", Salt:"123456", Source:"self"}

	fmt.Println(user)
	//orm.AutoMigrate(&models.User{})
	orm.NewRecord(user)
	orm.Create(user)
	orm.NewRecord(user)


	//user1 := models.User{}
	//orm.First(&user1, 2)
	//fmt.Println(user1)
	//
	//user2 := models.User{}
	//orm.First(&user2, 3)
	//fmt.Println(user2)
	//db.Model(&models.User{}).Select("name, password, salt").Find("name = ?", form.Name)
}

// User Signup
//func (us UserService) Signup(form forms.SignupForm) (user User, err error) {
//	//db := db.GetDB()
//
//
//
//
//	return user, nil
//}
//
//// User Signout
//func (us UserService) Signout() (err error) {
//
//
//	return nil
//}
//
//// User Rest Password
//func (us UserService) ForgetPassword(form forms.ForgetPasswordForm) (user User, err error) {
//
//
//	return user, nil
//}
//
//// User Rest Password
//func (us UserService) ResetPassword(form forms.ResetPasswordForm) (user User, err error) {
//
//
//	return user, nil
//}
//
//// User bind mobile
//func (us UserService) BindMobile(form forms.BindMobileForm) (user User, err error) {
//
//
//	return user, nil
//}


