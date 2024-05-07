package rpc

import (
	"context"
	"time"

	"HuaTug.com/kitex_gen/relations"
	"HuaTug.com/kitex_gen/relations/followservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
)

var relationClient followservice.Client

func initRealtionRpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		klog.Info(err)
	}

	c, err := followservice.NewClient(
		"Relation",
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
	relationClient = c
}

func Relation(ctx context.Context, req *relations.RelationServiceRequest) (resp *relations.RelationServiceResponse, err error) {
	resp = new(relations.RelationServiceResponse)
	resp,err=relationClient.RelationService(ctx,req)
	if err!=nil{
		return resp,err
	}
	return resp, nil
}

func RelationPage(ctx context.Context, req *relations.RelationServicePageRequest) (resp *relations.RelationServicePageResponse, err error) {
	resp = new(relations.RelationServicePageResponse)
	resp,err=relationClient.RelationServicePage(ctx,req)
	if err!=nil{
		return resp,err
	}
	return resp, nil
}
