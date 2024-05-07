package service

import (
	"context"

	"HuaTug.com/cmd/user/dal/db"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type DeleteUserService struct {
	ctx context.Context
}

func NewDeleteUSerService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

func (v *DeleteUserService) DeleteUser(userId int64) error {
	if err := db.DeleteUser(v.ctx, userId); err != nil {
		hlog.Info(err)
		return err
	}
	return nil
}
