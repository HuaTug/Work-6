package rpc

import (
	"context"
	"errors"
	"time"

	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/kitex_gen/users/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		klog.Info(err)
	}

	c, err := userservice.NewClient(
		"User",
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
	userClient = c
}

func CreateUser(ctx context.Context, req *users.CreateUserRequest) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return errors.New("Fail to create User!")
	}
	return nil
}

func LoginUser(ctx context.Context, req *users.LoginUserResquest) (resp *users.LoginUserResponse, err error) {
	resp, err = userClient.LoginUser(ctx, req)
	if err != nil {
		return resp, errors.New("Fail to Login!")
	}
	return resp, nil
}

func CheckUser(ctx context.Context, req *users.LoginUserResquest) (users.User, error) {
	var user users.User
	resp, err := userClient.LoginUser(ctx, req)
	if err != nil {
		return user, err
	}
	return *resp.User, nil
}

func QueryUser(ctx context.Context, req *users.QueryUserRequest) (resp *users.QueryUserResponse, err error) {
	resp, err = userClient.QueryUser(ctx, req)
	if err != nil {
		return resp, errors.New("Fail to Query")
	}
	return resp, nil
}

func DeleteUser(ctx context.Context, req *users.DeleteUserRequest) (resp *users.DeleteUserResponse, err error) {
	resp, err = userClient.DeleteUser(ctx, req)
	if err != nil {
		return resp, errors.New("Fail to Delete")
	}
	return resp, nil
}

func GetUserInfo(ctx context.Context, req *users.GetUserInfoRequest) (resp *users.GetUserInfoResponse, err error) {
	resp, err = userClient.GetUserInfo(ctx, req)
	if err != nil {
		return resp, errors.New("Fail to GetInfo")
	}
	return resp, nil
}

func UpdateUser(ctx context.Context, req *users.UpdateUserRequest) (resp *users.UpdateUserResponse, err error) {
	resp, err = userClient.UpdateUser(ctx, req)
	if err != nil {
		return resp, errors.New("Fail to Update")
	}
	return resp, nil
}
