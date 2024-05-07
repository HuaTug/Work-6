package main

import (
	"context"

	"HuaTug.com/cmd/video/service"
	"HuaTug.com/kitex_gen/videos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type VideoServiceImpl struct{}

func (s *VideoServiceImpl) FeedService(ctx context.Context, req *videos.FeedServiceRequest) (resp *videos.FeedServiceResponse, err error) {
	resp = new(videos.FeedServiceResponse)
	var video []*videos.Video
	if video, err = service.NewFeedListService(ctx).FeedList(req); err != nil {
		hlog.Info(err)
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to Get VideoFeed!"
		resp.VideoList = video
		return resp, err
	}
	resp.Code = consts.StatusOK
	resp.Msg = "Get VideoFeed Success"
	resp.VideoList = video
	return resp, nil
}

func (s *VideoServiceImpl) VideoFeedList(ctx context.Context, req *videos.VideoFeedListRequest) (resp *videos.VideoFeedListResponse, err error) {
	resp = new(videos.VideoFeedListResponse)
	var video []*videos.Video
	var count int64
	if video, count, err = service.NewVideoListService(ctx).VideoList(req); err != nil {
		hlog.Info(err)
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to Get VideoList!"
		resp.VideoList = video
		resp.Count = count
		return resp, err
	}
	resp.Code = consts.StatusOK
	resp.Msg = "Get VideoList Success"
	resp.VideoList = video
	resp.Count = count
	return resp, nil
}

func (s *VideoServiceImpl) VideoSearch(ctx context.Context, req *videos.VideoSearchRequest) (resp *videos.VideoSearchResponse, err error) {
	resp = new(videos.VideoSearchResponse)
	var video []*videos.Video
	var count int64
	if video, count, err = service.NewVideoSearchService(ctx).VideoSearch(req); err != nil {
		hlog.Info(err)
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to Get VideoFeed!"
		resp.VideoSearch = video
		resp.Count = count

		return resp, err
	}

	resp.Code = consts.StatusOK
	resp.Msg = "Get VideoFeed Success"
	resp.VideoSearch = video
	resp.Count = count

	return resp, nil
}

func (s *VideoServiceImpl) VideoPopular(ctx context.Context, req *videos.VideoPopularRequest) (resp *videos.VideoPopularResponse, err error) {
	resp = new(videos.VideoPopularResponse)
	var video []*videos.Video
	if video, err = service.NewVideoPopularService(ctx).VideoPopular(req); err != nil {
		hlog.Info(err)
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to Get VideoFeed!"
		resp.Popular=video
		return resp, err
	}
	resp.Code = consts.StatusOK
	resp.Msg = "Get VideoFeed Success"
	resp.Popular=video
	return resp, nil
}
