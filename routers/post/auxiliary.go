package post

import (
	"github.com/missdeer/kelly/routers/base"
)

// HomeRouter serves home page.
type AuxiliaryRouter struct {
	base.BaseRouter
}

func (this *AuxiliaryRouter) About() {
	this.TplName = "post/about.html"
}

func (this *AuxiliaryRouter) FAQ() {
	this.TplName = "post/faq.html"
}

func (this *AuxiliaryRouter) Contact() {
	this.TplName = "post/contact.html"
}

func (this *AuxiliaryRouter) Err401() {
	this.TplName = "post/error/401.html"
}

func (this *AuxiliaryRouter) Err403() {
	this.TplName = "post/error/403.html"
}

func (this *AuxiliaryRouter) Err404() {
	this.TplName = "post/error/404.html"
}

func (this *AuxiliaryRouter) Err500() {
	this.TplName = "post/error/500.html"
}

func (this *AuxiliaryRouter) Err503() {
	this.TplName = "post/error/503.html"
}
