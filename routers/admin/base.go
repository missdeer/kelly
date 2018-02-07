package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/missdeer/kelly/modules/models"
)

type AdminRouter struct {
	BaseAdminRouter
}

func (this *AdminRouter) ModelGet() {
	id := this.GetString("id")
	model := this.GetString("model")
	result := map[string]interface{}{
		"success": false,
	}

	var data []orm.ParamsList

	defer func() {
		if len(data) > 0 {
			result["success"] = true
			result["data"] = data[0]
		}
		this.Data["json"] = result
		this.ServeJSON()
	}()

	var qs orm.QuerySeter

	switch model {
	case "User":
		qs = models.Users()
	}

	qs = qs.Filter("Id", id).Limit(1)

	switch model {
	case "User":
		qs.ValuesList(&data, "Id", "UserName")
	}
}

func (this *AdminRouter) ModelSelect() {
	search := this.GetString("search")
	model := this.GetString("model")
	result := map[string]interface{}{
		"success": false,
	}

	var data []orm.ParamsList

	defer func() {
		if len(data) > 0 {
			result["success"] = true
			result["data"] = data
		}
		this.Data["json"] = result
		this.ServeJSON()
	}()

	if len(search) < 3 {
		return
	}

	switch model {
	case "User":
		models.Users().Filter("UserName__icontains", search).Limit(10).ValuesList(&data, "Id", "UserName")
	}
}
