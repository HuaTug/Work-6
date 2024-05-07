package cache

import "github.com/sirupsen/logrus"

func CacheChangeUserCount(userid, op int64, types string) {
	users, err := CacheGetUser(userid)
	//fmt.Println(users.UserID, users.UserName, users.FollowCount)
	if err != nil {
		logrus.Printf("user:%v miss cache", userid)
		return
	}
	switch types {
	case "follow":
		users.FollowCount += op
	case "follower":
		users.FollowerCount += op
	case "like":
		users.FavoriteCount += op
	case "unlike":
		users.FavoriteCount += op
	}
	CacheSetUser(users)
}
