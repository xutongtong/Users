package services

import (
	"fmt"

	"Users/db"
	. "Users/models"
)

type UserService struct {}

// User Signin
func (us UserService) Login(user User) {
	db := db.GetDB()

	db.Model(&user).Select("name, password, salt").Find("name = ?", user.Name)

	fmt.Println(user)
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


