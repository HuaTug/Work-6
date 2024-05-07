package service

import (
	"context"

	"HuaTug.com/cmd/publish/dal/db"
	"HuaTug.com/kitex_gen/publishs"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type VideoCreateService struct {
	ctx context.Context
}

func NewVideoCreateService(ctx context.Context) *VideoCreateService {
	return &VideoCreateService{ctx: ctx}
}

func (v *VideoCreateService)VideoCreate(ctx context.Context,req *publishs.VideoCreateRequest)error{
	err:=db.VideoCreate(v.ctx,req)
	if err!=nil{
		hlog.Info(err)
		return err
	}
	return nil	
}