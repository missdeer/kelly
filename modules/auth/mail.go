package auth

import (
	"fmt"

	"github.com/beego/i18n"

	"github.com/missdeer/kelly/modules/mailer"
	"github.com/missdeer/kelly/modules/models"
	"github.com/missdeer/kelly/modules/utils"
)

// Send user register mail with active code
func SendRegisterMail(locale i18n.Locale, user *models.User) {
	code := CreateUserActiveCode(user, nil)

	subject := locale.Tr("mail.register_success_subject")

	data := mailer.GetMailTmplData(locale.Lang, user)
	data["Code"] = code
	body := utils.RenderTemplate("mail/auth/register_success.html", data)

	msg := mailer.NewMailMessage([]string{user.Email}, subject, body)
	msg.Info = fmt.Sprintf("UID: %d, send register mail", user.Id)

	// async send mail
	mailer.SendAsync(msg)
}

// Send user reset password mail with verify code
func SendResetPwdMail(locale i18n.Locale, user *models.User) {
	code := CreateUserResetPwdCode(user, nil)

	subject := locale.Tr("mail.reset_password_subject")

	data := mailer.GetMailTmplData(locale.Lang, user)
	data["Code"] = code
	body := utils.RenderTemplate("mail/auth/reset_password.html", data)

	msg := mailer.NewMailMessage([]string{user.Email}, subject, body)
	msg.Info = fmt.Sprintf("UID: %d, send reset password mail", user.Id)

	// async send mail
	mailer.SendAsync(msg)
}

// Send email verify active email.
func SendActiveMail(locale i18n.Locale, user *models.User) {
	code := CreateUserActiveCode(user, nil)

	subject := locale.Tr("mail.verify_your_email_subject")

	data := mailer.GetMailTmplData(locale.Lang, user)
	data["Code"] = code
	body := utils.RenderTemplate("mail/auth/active_email.html", data)

	msg := mailer.NewMailMessage([]string{user.Email}, subject, body)
	msg.Info = fmt.Sprintf("UID: %d, send email verify mail", user.Id)

	// async send mail
	mailer.SendAsync(msg)
}
