package db

import (
	"context"

	"HuaTug.com/kitex_gen/favorites"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/kitex_gen/videos"
	"github.com/pkg/errors"
)

func FavoriteAction(ctx context.Context, favorite *favorites.Favorite) error {
	if err := DB.WithContext(ctx).Model(&favorites.Favorite{}).Create(favorite).Error; err != nil {
		return errors.Wrapf(err, "FavoriteAction failed,err:%v", err)
	}
	return nil
}

func UnFavoriteAction(ctx context.Context, userId, videoId int64) error {
	if err := DB.WithContext(ctx).Model(&favorites.Favorite{}).Where("user_id=? And video_id=?", userId, videoId).Delete(&favorites.Favorite{}).Error; err != nil {
		return errors.Wrapf(err, "UnFavoriteAction failed,err:%v", err)
	}
	return nil
}

func Judge(ctx context.Context, VideoId int64) bool {
	var count int64
	//ToDo :通过查阅资料 可以通过Count计数进行判断所需要查询的数据是否存在 若不用这种方式则会导致即使查询为空 也不会报错 因为err只是查询是否成功 而err==nil只表示查询语法没有出错 不只代表数据查询为空
	//这完善了点赞机制
	//ToDo question:如何让不同的人对视频进行点赞操作

	if DB.WithContext(ctx).Model(&videos.Video{}).Where("video_id=? And favorite_count!=?", VideoId, 0).Count(&count); count > 0 {
		return false
	}
	return true
}

func FavoriteExist(ctx context.Context, uid int64) ([]*favorites.Favorite, error) {
	var fav []*favorites.Favorite
	if err := DB.WithContext(ctx).Model(&favorites.Favorite{}).Where("user_id=?", uid).Find(&fav).Error; err != nil {
		return fav, errors.Wrapf(err, "FavoriteExist failed,err:%v", err)
	}
	return fav, nil
}

func UserExist(ctx context.Context, uid int64) ([]*favorites.User, error) {
	var user []*favorites.User
	if err := DB.WithContext(ctx).Model(&users.User{}).Where("user_id=?", uid).Find(&user).Error; err != nil {
		return user, errors.Wrapf(err, "UserExist failed,err:%v", err)
	}
	return user, nil
}
