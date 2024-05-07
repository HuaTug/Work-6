package service

import (
	"context"
	"errors"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/users"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (v *QueryUserService) QueryUserInfo(req *users.QueryUserRequest) (user []*users.User, total int64,err error) {
	var count int64
	if user,count,err= db.QueryUser(v.ctx, req.Keyword,req.Page,req.PageSize); err != nil {
		return user,count,errors.New("Query Failed")
	}
	return user,count,nil
}
