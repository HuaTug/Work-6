package rpc

import (
	"context"
	"time"

	"HuaTug.com/kitex_gen/comments"
	"HuaTug.com/kitex_gen/comments/commentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
)

var commentClient commentservice.Client

func initCommentRpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		klog.Info(err)
	}

	c, err := commentservice.NewClient(
		"Comment",
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
	commentClient = c
}

func CreateComment(ctx context.Context, req *comments.CreateCommentRequest) (*comments.CreateCommentResponse, error) {
	resp, err := commentClient.CreateComment(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func ListComment(ctx context.Context, req *comments.ListCommentRequest) (*comments.ListCommentResponse, error) {
	resp, err := commentClient.ListComment(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func DeleteComment(ctx context.Context, req *comments.CommentDeleteRequest) (*comments.CommentDeleteResponse, error) {
	resp, err := commentClient.DeleteComment(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
