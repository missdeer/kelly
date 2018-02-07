package base

import (
	"github.com/missdeer/kelly/modules/mailer"
)

type TestRouter struct {
	BaseRouter
}

func (this *TestRouter) Get() {
	this.TplName = this.GetString(":tmpl")
	this.Data = mailer.GetMailTmplData(this.Locale.Lang, &this.User)
	this.Data["Code"] = "CODE"
}
