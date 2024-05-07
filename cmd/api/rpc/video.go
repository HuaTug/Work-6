package rpc

import (
	"context"
	"time"

	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/kitex_gen/videos/videoservice"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
)

var VideoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		klog.Info(err)
	}

	c, err := videoservice.NewClient(
		"Video",
		/* 		client.WithMiddleware(middleware.CommonMiddleware),
		   		client.WithInstanceMW(middleware.ClientMiddleware), */
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		logrus.Info(err)
	}
	VideoClient = c
}

func FeedList(ctx context.Context, req *videos.FeedServiceRequest) (resp *videos.FeedServiceResponse, err error) {
	resp, err = VideoClient.FeedService(ctx, req)
	if err != nil {
		if resp == nil {
			return
		}
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to provide FeedService!"
		return resp, err
	}
	return resp, nil
}

func VideoFeedList(ctx context.Context, req *videos.VideoFeedListRequest) (resp *videos.VideoFeedListResponse, err error) {
	resp, err = VideoClient.VideoFeedList(ctx, req)
	if err != nil {
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to provide VideoFeedList Service!"
		return resp, err
	}
	return resp, err
}

func VideoSearch(ctx context.Context, req *videos.VideoSearchRequest) (resp *videos.VideoSearchResponse, err error) {
	resp, err = VideoClient.VideoSearch(ctx, req)
	if err != nil {
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to provide VideoSearch Service!"
		return resp, err
	}
	return resp, err
}

func VideoPopular(ctx context.Context, req *videos.VideoPopularRequest) (resp *videos.VideoPopularResponse, err error) {
	resp, err = VideoClient.VideoPopular(ctx, req)
	if err != nil {
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to provide VideoPopular Service!"
		return resp, err
	}
	return resp, err
}
