package rpc

import (
	"context"
	"time"

	"HuaTug.com/kitex_gen/favorites"
	"HuaTug.com/kitex_gen/favorites/favoriteservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		klog.Info(err)
	}

	c, err := favoriteservice.NewClient(
		"Favorite",
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
	favoriteClient = c
}
func Favorite(ctx context.Context, req *favorites.FavoriteRequest) (resp *favorites.FavoriteResponse, err error) {
	resp=new(favorites.FavoriteResponse)
	resp,err=favoriteClient.FavoriteService(ctx,req)
	if err!=nil{
		return resp,err
	}
	return resp,nil
}

func ListFavorite(ctx context.Context,req *favorites.ListFavoriteRequest)(resp *favorites.ListFavoriteResponse,err error){
	resp=new(favorites.ListFavoriteResponse)
	resp,err=favoriteClient.ListFavorite(ctx,req)
	if err!=nil{
		return resp,err
	}
	return resp,nil
}