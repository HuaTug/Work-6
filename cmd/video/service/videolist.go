package service

import (
	"context"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/kitex_gen/videos"
)

type VideoListService struct {
	ctx context.Context
}

func NewVideoListService(ctx context.Context) *VideoListService {
	return &VideoListService{ctx: ctx}
}

func (v *VideoListService)VideoList(req *videos.VideoFeedListRequest)(video []*videos.Video,count int64,err error){
	if video,count,err=db.Videolist(v.ctx,req);err!=nil{
		return video,count,err
	}
	return video,count,err
}