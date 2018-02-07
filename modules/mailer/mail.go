package mailer

import (
	"github.com/missdeer/kelly/modules/models"
	"github.com/missdeer/kelly/setting"
)

// Create New mail message use MailFrom and MailUser
func NewMailMessage(To []string, subject, body string) Message {
	msg := NewHtmlMessage(To, setting.MailFrom, subject, body)
	msg.User = setting.MailUser
	return msg
}

func GetMailTmplData(lang string, user *models.User) map[interface{}]interface{} {
	data := make(map[interface{}]interface{}, 10)
	data["AppName"] = setting.AppName
	data["AppVer"] = setting.AppVer
	data["AppUrl"] = "https://" + setting.AppHost + "/" //setting.AppUrl
	data["AppLogo"] = setting.AppLogo
	data["IsProMode"] = setting.IsProMode
	data["Lang"] = lang
	data["ActiveCodeLives"] = setting.ActiveCodeLives
	data["ResetPwdCodeLives"] = setting.ResetPwdCodeLives
	if user != nil {
		data["User"] = user
	}
	return data
}
