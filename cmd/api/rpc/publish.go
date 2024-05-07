package rpc

import (
	"context"
	"time"

	"HuaTug.com/kitex_gen/publishs"
	"HuaTug.com/kitex_gen/publishs/uploadvideoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
)

var uploadVideoClient uploadvideoservice.Client

func initUploadrpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		klog.Info(err)
	}

	c, err := uploadvideoservice.NewClient(
		"Publish",
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
	uploadVideoClient = c
}

func UploadVideo(ctx context.Context, req *publishs.UpLoadVideoRequest) (*publishs.UpLoadVideoResponse, error) {
	/* 	if uploadVideoClient == nil {
		hlog.Info("Client Error")
		return nil, nil
	} */
	resp, err := uploadVideoClient.UploadVideo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func VideoCreate(ctx context.Context, req *publishs.VideoCreateRequest) (resp *publishs.VideoCreateResponse, err error) {
	resp, err = uploadVideoClient.VideoCreate(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}