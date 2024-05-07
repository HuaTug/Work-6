package service

import (
	"context"
	"strconv"

	"HuaTug.com/cache"
	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/kitex_gen/videos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type VideoPopularService struct {
	ctx context.Context
}

func NewVideoPopularService(ctx context.Context) *VideoPopularService {
	return &VideoPopularService{ctx: ctx}
}

func (v *VideoPopularService) VideoPopular(req *videos.VideoPopularRequest) (video []*videos.Video, err error) {
	res, err := cache.RangeList("Rank")
	if err != nil {
		hlog.Info(err)
		return
	}
	var temp *videos.Video
	for i := 0; i < len(res); i++ {
		s, err := strconv.Atoi(res[i])
		if err != nil {
			hlog.Info(err)
			return video, err
		}
		temp, _ = db.FindVideo(v.ctx, int64(s))
		video = append(video, temp)
	}
	return video,nil
}
