package post

import (
	"github.com/astaxie/beego/validation"
	"github.com/missdeer/kelly/modules/models"
	"github.com/missdeer/kelly/modules/utils"
)

type TopicAdminForm struct {
	Create    bool   `form:"-"`
	Id        int    `form:"-"`
	Name      string `valid:"Required;MaxSize(30)"`
	Intro     string `form:"type(textarea)" valid:"Required"`
	NameZhCn  string `valid:"Required;MaxSize(30)"`
	IntroZhCn string `form:"type(textarea)" valid:"Required"`
	Slug      string `valid:"Required;MaxSize(100)"`
	Followers int    ``
	Order     int    ``
	Image     string `valid:""`
}

func (form *TopicAdminForm) Valid(v *validation.Validation) {
	qs := models.Topics()

	if models.CheckIsExist(qs, "Name", form.Name, form.Id) {
		v.SetError("Name", "admin.field_need_unique")
	}

	if models.CheckIsExist(qs, "NameZhCn", form.NameZhCn, form.Id) {
		v.SetError("NameZhCn", "admin.field_need_unique")
	}

	if models.CheckIsExist(qs, "Slug", form.Slug, form.Id) {
		v.SetError("Slug", "admin.field_need_unique")
	}
}

func (form *TopicAdminForm) SetFromTopic(topic *models.Topic) {
	utils.SetFormValues(topic, form)
}

func (form *TopicAdminForm) SetToTopic(topic *models.Topic) {
	utils.SetFormValues(form, topic, "Id")
}

type CategoryAdminForm struct {
	Create bool   `form:"-"`
	Id     int    `form:"-"`
	Name   string `valid:"Required;MaxSize(30)"`
	Slug   string `valid:"Required;MaxSize(100)"`
	Order  int    ``
}

func (form *CategoryAdminForm) Valid(v *validation.Validation) {
	qs := models.Categories()

	if models.CheckIsExist(qs, "Name", form.Name, form.Id) {
		v.SetError("Name", "admin.field_need_unique")
	}

	if models.CheckIsExist(qs, "Slug", form.Slug, form.Id) {
		v.SetError("Slug", "admin.field_need_unique")
	}
}

func (form *CategoryAdminForm) SetFromCategory(cat *models.Category) {
	utils.SetFormValues(cat, form)
}

func (form *CategoryAdminForm) SetToCategory(cat *models.Category) {
	utils.SetFormValues(form, cat, "Id")
}
