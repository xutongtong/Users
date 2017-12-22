package services

import (
	"Users/db"
	//"Users/models"
	"Users/forms"
	"fmt"
	"Users/models"
)

type UserService struct {}


// User Signin
func (us UserService) Login(form forms.UserLoginForm) {
	fmt.Println("Start Login")

	db.GetDB()

	user := models.User{Name:form.Name, Password:form.Password}

	orm.AutoMigrate(&models.User{})
	orm.NewRecord(user)
	orm.First(&user, 1)
	orm.Exec("SELECT * FROM users")
	//
	fmt.Println(user)
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


