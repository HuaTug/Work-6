package service

import (
	"context"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/kitex_gen/videos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type VideoSearchService struct{
	ctx context.Context
}

func NewVideoSearchService(ctx context.Context)*VideoSearchService{
	return &VideoSearchService{ctx :ctx}
}

func (v *VideoSearchService)VideoSearch(req *videos.VideoSearchRequest)(video []*videos.Video,count int64,err error){
	if video,count,err=db.Videosearch(v.ctx,req);err!=nil{
		hlog.Info(err)
		return video,count,err
	}
	return video,count,err
}
