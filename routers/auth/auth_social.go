package auth

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/beego/social-auth"
	"github.com/missdeer/kelly/modules/auth"
	"github.com/missdeer/kelly/modules/models"
	"github.com/missdeer/kelly/modules/utils"
	"github.com/missdeer/kelly/routers/base"
	"github.com/missdeer/kelly/setting"
)

type socialAuther struct {
}

func (p *socialAuther) IsUserLogin(ctx *context.Context) (int, bool) {
	if id := auth.GetUserIdFromSession(ctx.Input.CruSession); id > 0 {
		return id, true
	}
	return 0, false
}

func (p *socialAuther) LoginUser(ctx *context.Context, uid int) (string, error) {
	user := models.User{Id: uid}
	if user.Read() == nil {
		auth.LoginUser(&user, ctx, true)
	}
	return auth.GetLoginRedirect(ctx), nil
}

var SocialAuther social.SocialAuther = new(socialAuther)

func OAuthRedirect(ctx *context.Context) {
	redirect, err := setting.SocialAuth.OAuthRedirect(ctx)
	if err != nil {
		beego.Error("OAuthRedirect", err)
	}

	if len(redirect) > 0 {
		ctx.Redirect(302, redirect)
	}
}

func OAuthAccess(ctx *context.Context) {
	redirect, _, err := setting.SocialAuth.OAuthAccess(ctx)
	if err != nil {
		beego.Error("OAuthAccess", err)
	}

	if len(redirect) > 0 {
		ctx.Redirect(302, redirect)
	}
}

type SocialAuthRouter struct {
	base.BaseRouter
}

func (this *SocialAuthRouter) canConnect(socialType *social.SocialType) bool {
	if st, ok := setting.SocialAuth.ReadyConnect(this.Ctx); !ok {
		return false
	} else {
		*socialType = st
	}
	return true
}

func (this *SocialAuthRouter) Connect() {
	this.TplName = "auth/connect.html"

	if this.CheckLoginRedirect(false) {
		return
	}

	var socialType social.SocialType
	if !this.canConnect(&socialType) {
		this.Redirect(setting.SocialAuth.LoginURL, 302)
		return
	}

	formL := auth.OAuthLoginForm{}
	this.SetFormSets(&formL)

	formR := auth.OAuthRegisterForm{Locale: this.Locale}
	this.SetFormSets(&formR)

	this.Data["Action"] = this.GetString("action")
	this.Data["Social"] = socialType
}

func (this *SocialAuthRouter) ConnectPost() {
	this.TplName = "auth/connect.html"

	if this.CheckLoginRedirect(false) {
		return
	}

	var socialType social.SocialType
	if !this.canConnect(&socialType) {
		this.Redirect(setting.SocialAuth.LoginURL, 302)
		return
	}

	p, ok := social.GetProviderByType(socialType)
	if !ok {
		this.Redirect(setting.SocialAuth.LoginURL, 302)
		return
	}

	var form interface{}

	formL := auth.OAuthLoginForm{}
	this.SetFormSets(&formL)

	formR := auth.OAuthRegisterForm{Locale: this.Locale}
	this.SetFormSets(&formR)

	action := this.GetString("action")
	if action == "connect" {
		form = &formL
	} else {
		form = &formR
	}

	this.Data["Action"] = action
	this.Data["Social"] = socialType

	// valid form and put errors to template context
	if this.ValidFormSets(form) == false {
		return
	}

	var user models.User

	switch action {
	case "connect":
		key := "auth.login." + formL.UserName + this.Ctx.Input.IP()
		if times, ok := utils.TimesReachedTest(key, setting.LoginMaxRetries); ok {
			this.Data["ErrorReached"] = true
		} else if auth.VerifyUser(&user, formL.UserName, formL.Password) {
			goto connect
		} else {
			utils.TimesReachedSet(key, times, setting.LoginFailedBlocks)
		}

	default:
		if err := auth.RegisterUser(&user, formR.UserName, formR.Email, formR.Password); err == nil {

			auth.SendRegisterMail(this.Locale, &user)

			goto connect

		} else {
			beego.Error("Register: Failed ", err)
		}
	}

failed:
	this.Data["Error"] = true
	return

connect:
	if loginRedirect, _, err := setting.SocialAuth.ConnectAndLogin(this.Ctx, socialType, user.Id); err != nil {
		beego.Error("ConnectAndLogin:", err)
		goto failed
	} else {
		this.Redirect(loginRedirect, 302)
		return
	}

	switch action {
	case "connect":
		this.FlashRedirect("/settings/profile", 302, "ConnectSuccess", p.GetName())
	default:
		this.FlashRedirect("/settings/profile", 302, "RegSuccess")
	}
}
