package db

import (
	"context"

	"sync"

	"HuaTug.com/kitex_gen/videos"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Feedlist(ctx context.Context, req *videos.FeedServiceRequest) ([]*videos.Video, error) {
	var video []*videos.Video
	if err := DB.WithContext(ctx).Model(&videos.Video{}).Where("publish_time<?", req.LastTime).Find(&video); err != nil {
		return video, errors.Wrapf(err.Error, "FeedList failed,err:%v", err)
	}
	return video, nil
}

func Videolist(ctx context.Context, req *videos.VideoFeedListRequest) ([]*videos.Video, int64, error) {
	var video []*videos.Video
	var count int64
	if err := DB.Model(&videos.Video{}).Where("author_id=?", req.AuthorId).Count(&count).Limit(int(req.PageSize)).
		Offset(int((req.PageNum - 1) * req.PageSize)).Find(&video); err != nil {
		logrus.Info(err)
		return video, count, errors.Wrapf(err.Error, "VideoList failed,err:%v", err)
	}
	return video, count, nil
}

func Videosearch(ctx context.Context, req *videos.VideoSearchRequest) ([]*videos.Video, int64, error) {
	var wg sync.WaitGroup
	var video2 []*videos.Video
	var count int64
	var err error
	if req.Keyword != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = DB.Model(&videos.Video{}).
				Where("title like ? And publish_time<? &&publish_time>?", "%"+req.Keyword+"%", req.ToDate, req.FromDate).
				Count(&count).
				Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).
				Find(&video2).Error
		}()
		if err != nil {
			return video2, count, errors.Wrapf(err,"VideoSearch failed,err:%v",err)
		}
		wg.Wait()
	}
	return video2, count, nil
}

func FindVideo(ctx context.Context, videoId int64) (video *videos.Video, err error) {
	if err := DB.Model(&videos.Video{}).Where("video_id=?", videoId).Find(&video); err != nil {
		return video, errors.Wrapf(err.Error,"FindVideo failed,err:%v",err)
	}
	return video, nil
}
