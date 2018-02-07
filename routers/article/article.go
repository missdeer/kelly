package article

import (
	"github.com/astaxie/beego"
	"github.com/missdeer/kelly/modules/models"
	"github.com/missdeer/kelly/routers/base"
)

type ArticleRouter struct {
	base.BaseRouter
}

func (this *ArticleRouter) loadArticle(article *models.Article) bool {
	uri := this.GetString(":slug")
	err := models.Articles().RelatedSel("User").Filter("IsPublish", true).Filter("Uri", uri).One(article)
	if err == nil {
		this.Data["Article"] = article
	} else {
		beego.Error(err)
		this.Abort("404")
	}
	return err != nil
}

func (this *ArticleRouter) Show() {
	this.TplName = "article/show.html"
	article := models.Article{}
	if this.loadArticle(&article) {
		return
	}
}
