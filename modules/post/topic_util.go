package post

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/missdeer/kelly/modules/models"
)

func ListCategories(cats *[]models.Category) (int64, error) {
	return models.Categories().OrderBy("-order").All(cats)
}

func ListTopics(topics *[]models.Topic) (int64, error) {
	return models.Topics().OrderBy("-Followers").All(topics)
}

func ListTopicsOfCat(topics *[]models.Topic, cat *models.Category) (int64, error) {
	var list orm.ParamsList
	var where string
	if cat != nil {
		where = " WHERE category_id = ?"
	}

	sql := fmt.Sprintf(`SELECT topic_id
		FROM post%s
		GROUP BY topic_id
		ORDER BY COUNT(topic_id) DESC LIMIT 8`, where)

	rs := orm.NewOrm().Raw(sql)

	if cat != nil {
		rs = rs.SetArgs(cat.Id)
	}

	cnt, err := rs.ValuesFlat(&list)
	if err != nil {
		beego.Error("models.ListTopicsOfCat ", err)
		return 0, err
	}
	if cnt > 0 {
		nums, err := models.Topics().Filter("Id__in", list).All(topics)
		if err != nil {
			beego.Error("models.ListTopicsOfCat ", err)
			return 0, err
		}
		return nums, err
	}
	return 0, nil
}
