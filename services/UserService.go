package services

import (
	"../db"
	"../forms"
	. "../models"
	"fmt"
)

type UserService struct {}

// User Signin
func (us UserService) Signin(name, password string) (User, error) {
	db := db.GetDB()
	user := User{}

	db.Model(&user).Select("name, password, salt").Find("name = ?", name)

	if user.Password == password {
		fmt.Println("abc")
	}

	fmt.Println(user)

	return user, nil
}

// User Signup
func (us UserService) Signup(form forms.SignupForm) (user User, err error) {
	//db := db.GetDB()




	return user, nil
}

// User Signout
func (us UserService) Signout() (err error) {


	return nil
}

// User Rest Password
func (us UserService) ForgetPassword(form forms.ForgetPasswordForm) (user User, err error) {


	return user, nil
}

// User Rest Password
func (us UserService) ResetPassword(form forms.ResetPasswordForm) (user User, err error) {


	return user, nil
}

// User bind mobile
func (us UserService) BindMobile(form forms.BindMobileForm) (user User, err error) {


	return user, nil
}


