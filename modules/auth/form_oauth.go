package auth

import (
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
)

// OAuth connect Register form
type OAuthRegisterForm struct {
	UserName   string      `valid:"Required;AlphaDash;MinSize(5);MaxSize(30)"`
	Email      string      `valid:"Required;Email;MaxSize(80)"`
	Password   string      `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe string      `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	Locale     i18n.Locale `form:"-"`
}

func (form *OAuthRegisterForm) Valid(v *validation.Validation) {

	// Check if passwords of two times are same.
	if form.Password != form.PasswordRe {
		v.SetError("PasswordRe", "auth.repassword_not_match")
		return
	}

	e1, e2, _ := CanRegistered(form.UserName, form.Email)

	if !e1 {
		v.SetError("UserName", "auth.username_already_taken")
	}

	if !e2 {
		v.SetError("Email", "auth.email_already_taken")
	}
}

func (form *OAuthRegisterForm) Labels() map[string]string {
	return map[string]string{
		"UserName":   "auth.login_username",
		"Email":      "auth.login_email",
		"Password":   "auth.login_password",
		"PasswordRe": "auth.retype_password",
	}
}

func (form *OAuthRegisterForm) Helps() map[string]string {
	return map[string]string{
		"UserName": form.Locale.Tr("valid.min_length_is", 5) + ", " + form.Locale.Tr("valid.only_contains", "a-z 0-9 - _"),
	}
}

func (form *OAuthRegisterForm) Placeholders() map[string]string {
	return map[string]string{
		"UserName":   "auth.plz_enter_username",
		"Email":      "auth.plz_enter_email",
		"Password":   "auth.plz_enter_password",
		"PasswordRe": "auth.plz_reenter_password",
	}
}

// OAuth connect Login form
type OAuthLoginForm struct {
	UserName string `valid:"Required"`
	Password string `form:"type(password)" valid:"Required"`
}

func (form *OAuthLoginForm) Labels() map[string]string {
	return map[string]string{
		"UserName": "auth.username_or_email",
		"Password": "auth.login_password",
	}
}
