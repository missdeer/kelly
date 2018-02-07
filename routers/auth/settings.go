package auth

import (
	"github.com/astaxie/beego"
	"github.com/missdeer/kelly/modules/auth"
	"github.com/missdeer/kelly/routers/base"
)

// SettingsRouter serves user settings.
type SettingsRouter struct {
	base.BaseRouter
}

// Profile implemented user profile settings page.
func (this *SettingsRouter) Profile() {
	this.TplName = "settings/profile.html"

	// need login
	if this.CheckLoginRedirect() {
		return
	}

	form := auth.ProfileForm{Locale: this.Locale}
	form.SetFromUser(&this.User)
	this.SetFormSets(&form)

	formPwd := auth.PasswordForm{}
	this.SetFormSets(&formPwd)
}

// ProfileSave implemented save user profile.
func (this *SettingsRouter) ProfileSave() {
	this.TplName = "settings/profile.html"

	if this.CheckLoginRedirect() {
		return
	}

	action := this.GetString("action")

	if this.IsAjax() {
		switch action {
		case "send-verify-email":
			if this.User.IsActive {
				this.Data["json"] = false
			} else {
				auth.SendActiveMail(this.Locale, &this.User)
				this.Data["json"] = true
			}

			this.ServeJSON()
			return
		}
		return
	}

	profileForm := auth.ProfileForm{Locale: this.Locale}
	profileForm.SetFromUser(&this.User)

	pwdForm := auth.PasswordForm{User: &this.User}

	this.Data["Form"] = profileForm

	switch action {
	case "save-profile":
		if this.ValidFormSets(&profileForm) {
			if err := profileForm.SaveUserProfile(&this.User); err != nil {
				beego.Error("ProfileSave: save-profile", err)
			}
			this.FlashRedirect("/settings/profile", 302, "ProfileSave")
			return
		}

	case "change-password":
		if this.ValidFormSets(&pwdForm) {
			// verify success and save new password
			if err := auth.SaveNewPassword(&this.User, pwdForm.Password); err == nil {
				this.FlashRedirect("/settings/profile", 302, "PasswordSave")
				return
			} else {
				beego.Error("ProfileSave: change-password", err)
			}
		}

	default:
		this.Redirect("/settings/profile", 302)
		return
	}

	if action != "save-profile" {
		this.SetFormSets(&profileForm)
	}
	if action != "change-password" {
		this.SetFormSets(&pwdForm)
	}
}
