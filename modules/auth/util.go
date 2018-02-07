package auth

import (
	"github.com/astaxie/beego/orm"
	"github.com/missdeer/kelly/modules/models"
)

func UserFollow(user *models.User, theUser *models.User) {
	if theUser.Read() == nil {
		var mutual bool
		tFollow := models.Follow{User: theUser, FollowUser: user}
		if err := tFollow.Read("User", "FollowUser"); err == nil {
			mutual = true
		}

		follow := models.Follow{User: user, FollowUser: theUser, Mutual: mutual}
		if err := follow.Insert(); err == nil && mutual {
			tFollow.Mutual = mutual
			tFollow.Update("Mutual")
		}

		if nums, err := user.FollowingUsers().Count(); err == nil {
			user.Following = int(nums)
			user.Update("Following")
		}

		if nums, err := theUser.FollowerUsers().Count(); err == nil {
			theUser.Followers = int(nums)
			theUser.Update("Followers")
		}
	}
}

func UserUnFollow(user *models.User, theUser *models.User) {
	num, _ := user.FollowingUsers().Filter("FollowUser", theUser.Id).Delete()
	if num > 0 {
		theUser.FollowingUsers().Filter("FollowUser", user.Id).Update(orm.Params{
			"Mutual": false,
		})

		if nums, err := user.FollowingUsers().Count(); err == nil {
			user.Following = int(nums)
			user.Update("Following")
		}

		if nums, err := theUser.FollowerUsers().Count(); err == nil {
			theUser.Followers = int(nums)
			theUser.Update("Followers")
		}
	}
}
