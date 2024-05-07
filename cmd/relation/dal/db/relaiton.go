package db

import (
	"context"

	"HuaTug.com/kitex_gen/relations"
	"HuaTug.com/kitex_gen/users"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sirupsen/logrus"
)

func Follow(ctx context.Context, Relation *relations.Relation) error {
	if err := DB.WithContext(ctx).Model(&relations.Relation{}).Create(Relation).Error; err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}

func UnFollow(ctx context.Context, Relation *relations.Relation) error {
	if err := DB.WithContext(ctx).Model(&relations.Relation{}).Create(Relation).Error; err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}

func FollowList(ctx context.Context, req *relations.RelationServicePageRequest) ([]*relations.Relation, error) {
	var list []*relations.Relation
	var count int64
	if err := DB.WithContext(ctx).Model(&relations.Relation{}).Where("relation=?", req.UserId).Count(&count).
		Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&list).Error; err != nil {
		hlog.Info("分页查询出错")
		hlog.Info(err)
	}
	return list, nil
}

func GetUser(ctx context.Context, uid int64) []*relations.User {
	var user []*relations.User
	if err := DB.WithContext(ctx).Model(&users.User{}).Where("user_id=?", uid).Find(&user); err != nil {
		hlog.Info(err)
	}
	return user
}