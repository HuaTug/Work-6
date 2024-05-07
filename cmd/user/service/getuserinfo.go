package service

import (
	"context"

	"HuaTug.com/cache"
	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/users"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type GetUserInfoService struct {
	ctx context.Context
}

func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

func (v *GetUserInfoService) GetUserInfo(userId int64) (user *users.User, err error) {
	users, err := cache.CacheGetUser(userId)
	if err != nil {
		if user, err = db.GetUser(v.ctx, userId); err != nil {
			hlog.Info(err)
			return user, err
		} else {
			go cache.CacheSetUser(user)
			return user,nil
		}
	}
	return users, nil
}
