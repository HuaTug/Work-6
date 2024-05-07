package service

import (
	"context"

	"HuaTug.com/cache"
	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/kitex_gen/videos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)


type FeedListService struct{
	ctx context.Context
}

func NewFeedListService(ctx context.Context)*FeedListService{
	return &FeedListService{ctx:ctx}
}

//这里的v指向方法，用于传递ctx上下文
func (v *FeedListService)FeedList(req *videos.FeedServiceRequest)([]*videos.Video,error){
	if video,err:=db.Feedlist(v.ctx,req);err!=nil{
		hlog.Info(err)
		return video,err
	}else{
		
		cache.Insert(video)
		for _,s:=range video{
			err:=cache.RangeAdd(s.FavoriteCount,s.VideoId)
			if err!=nil{
			hlog.Info(err)
			}
		}
		return video,nil
	}
}