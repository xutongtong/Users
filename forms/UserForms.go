package forms

//SignForm

type SigninForm struct {
	Name      string
	Password  string
}

func (form SigninForm) validate() bool{
	if len(form.Name) == 0 {
		return true
	}

	return false
}

type SignUpForm struct {
	Mobile   string
	Code     string
	Password string
}

type ForgetPasswordForm struct {





}

type ResetPasswordForm struct {

}

type BindMobileForm struct {

}