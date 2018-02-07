package post

import (
	"fmt"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/missdeer/kelly/modules/models"
	"github.com/missdeer/kelly/modules/utils"
	"github.com/missdeer/kelly/setting"
)

func ListPostsOfCategory(cat *models.Category, posts *[]models.Post) (int64, error) {
	return models.Posts().Filter("Category", cat).RelatedSel().OrderBy("-Updated").All(posts)
}

func ListPostsOfTopic(topic *models.Topic, posts *[]models.Post) (int64, error) {
	return models.Posts().Filter("Topic", topic).RelatedSel().OrderBy("-Updated").All(posts)
}

var mentionRegexp = regexp.MustCompile(`\B@([\d\w-_]*)`)

func FilterMentions(user *models.User, content string) {
	matches := mentionRegexp.FindAllStringSubmatch(content, -1)
	mentions := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) > 1 {
			mentions = append(mentions, m[1])
		}
	}
	// var users []*User
	// num, err := Users().Filter("UserName__in", mentions).Filter("Follow__User", user.Id).All(&users)
	// if err == nil && num > 0 {
	// TODO mention email to user
	// }
}

func PostBrowsersAdd(uid int, ip string, post *models.Post) {
	var key string
	if uid == 0 {
		key = ip
	} else {
		key = utils.ToStr(uid)
	}
	key = fmt.Sprintf("PCA.%d.%s", post.Id, key)
	if setting.Cache.Get(key) != nil {
		return
	}
	_, err := models.Posts().Filter("Id", post.Id).Update(orm.Params{
		"Browsers": orm.ColValue(orm.ColAdd, 1),
	})
	if err != nil {
		beego.Error("PostCounterAdd ", err)
	}
	setting.Cache.Put(key, true, 60)
}

func PostReplysCount(post *models.Post) {
	cnt, err := post.Comments().Count()
	if err == nil {
		post.Replys = int(cnt)
		err = post.Update("Replys")

		_, err := models.Posts().Filter("Id", post.Id).Update(orm.Params{
			"TodayReplys": orm.ColValue(orm.ColAdd, 1),
		})
		if err != nil {
			beego.Error("PostTodayReply ", err)
		}
	}
	if err != nil {
		beego.Error("PostReplysCount ", err)
	}
}
