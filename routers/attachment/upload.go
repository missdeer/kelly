package attachment

import (
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/missdeer/kelly/modules/attachment"
	"github.com/missdeer/kelly/modules/models"
	"github.com/missdeer/kelly/routers/base"
	"github.com/missdeer/kelly/setting"
)

type UploadRouter struct {
	base.BaseRouter
}

func (this *UploadRouter) Post() {
	result := map[string]interface{}{
		"success": false,
	}

	defer func() {
		this.Data["json"] = &result
		this.ServeJSON()
	}()

	// check permition
	if !this.User.IsActive {
		return
	}

	// get file object
	file, handler, err := this.Ctx.Request.FormFile("image")
	if err != nil {
		return
	}
	defer file.Close()

	t := time.Now()

	image := models.Image{}
	image.User = &this.User

	// get mime type
	mime := handler.Header.Get("Content-Type")

	// save and resize image
	if err := attachment.SaveImage(&image, file, mime, handler.Filename, t); err != nil {
		beego.Error(err)
		return
	}

	result["link"] = image.LinkMiddle()
	result["success"] = true

}

func ImageFilter(ctx *context.Context) {
	if setting.QiniuEnabled && setting.UpYunEnabled {
		if time.Now().UnixNano()%2 == 0 {
			url := setting.QiniuUrl + "upload" + ctx.Request.RequestURI
			http.Redirect(ctx.ResponseWriter, ctx.Request, url, 302)
		} else {
			url := setting.UpYunUrl + "upload" + ctx.Request.RequestURI
			http.Redirect(ctx.ResponseWriter, ctx.Request, url, 302)
		}
		return
	}

	if setting.QiniuEnabled {
		url := setting.QiniuUrl + "upload" + ctx.Request.RequestURI
		http.Redirect(ctx.ResponseWriter, ctx.Request, url, 302)
		return
	}

	if setting.UpYunEnabled {
		url := setting.UpYunUrl + "upload" + ctx.Request.RequestURI
		http.Redirect(ctx.ResponseWriter, ctx.Request, url, 302)
	}
}
