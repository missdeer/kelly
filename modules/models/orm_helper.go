package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func CheckIsExist(qs orm.QuerySeter, field string, value interface{}, skipId int) bool {
	qs = qs.Filter(field, value)
	if skipId > 0 {
		qs = qs.Exclude("Id", skipId)
	}
	return qs.Exist()
}

func CountObjects(qs orm.QuerySeter) (int64, error) {
	cnt, err := qs.Count()
	if err != nil {
		beego.Error("models.CountObjects ", err)
		return 0, err
	}
	return cnt, err
}

func ListObjects(qs orm.QuerySeter, objs interface{}) (int64, error) {
	nums, err := qs.All(objs)
	if err != nil {
		beego.Error("models.ListObjects ", err)
		return 0, err
	}
	return nums, err
}
