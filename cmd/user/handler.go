// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"

	"HuaTug.com/cmd/user/service"
	"HuaTug.com/kitex_gen/users"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *users.CreateUserRequest) (resp *users.CreateUserResponse, err error) {
	resp = new(users.CreateUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		return resp, nil
	}

	return resp, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *users.UpdateUserRequest) (resp *users.UpdateUserResponse, err error) {
	resp = new(users.UpdateUserResponse)
	if err := service.NewUpdateUserService(ctx).UpdateUser(req); err != nil {
		hlog.Info(err)
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail To Update User"
		return resp, err
	}
	resp.Code = consts.StatusOK
	resp.Msg = "Update User Success"
	return resp, nil
}

func (s *UserServiceImpl) LoginUser(ctx context.Context, req *users.LoginUserResquest) (resp *users.LoginUserResponse, err error) {
	resp = new(users.LoginUserResponse)
	var user users.User
	user, err = service.NewLoginUserService(ctx).LoginUsers(req)
	if err != nil {
		return resp, err
	}

	resp.User = &user
	return resp, nil
}

func (s *UserServiceImpl) CheckUser(ctx context.Context, req *users.LoginUserResquest) (user users.User, err error) {
	user, err = service.NewLoginUserService(ctx).LoginUsers(req)
	if err != nil {
		hlog.Info(err)
		return user, err
	}
	return user, nil

}
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *users.QueryUserRequest) (resp *users.QueryUserResponse, err error) {
	resp = new(users.QueryUserResponse)
	resp.Users, resp.Totoal, err = service.NewQueryUserService(ctx).QueryUserInfo(req)
	if err != nil {
		hlog.Info(err)
		return resp, err
	}
	resp.Code = 200
	resp.Msg = "Query Success"
	return resp, nil
}

func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *users.GetUserInfoRequest) (resp *users.GetUserInfoResponse, err error) {
	resp = new(users.GetUserInfoResponse)
	var user *users.User
	hlog.Info(req.UserId)
	if user, err = service.NewGetUserInfoService(ctx).GetUserInfo(req.UserId); err != nil {
		hlog.Info(err)
		return resp, err
	}
	resp.User = user
	return resp, nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *users.DeleteUserRequest) (resp *users.DeleteUserResponse, err error) {
	resp = new(users.DeleteUserResponse)
	if err := service.NewDeleteUSerService(ctx).DeleteUser(req.UserId); err != nil {
		hlog.Info(err)
		resp.Msg = "Fail to Delete User!"
		resp.Code = 500
		return resp, err
	}
	resp.Code = 200
	resp.Msg = "Delete User Success!"
	return resp, nil
}
