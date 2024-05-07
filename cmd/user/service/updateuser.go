package service

import (
	"context"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/users"
)

type UpdateUserService struct {
	ctx context.Context
}

func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

func (v *UpdateUserService) UpdateUser(req *users.UpdateUserRequest) (err error) {
	user:=&users.User{
		UserId: req.UserId,
		UserName: req.UserName,
		Password: req.Password,
	}
	if err:=db.UpdateUser(v.ctx,user);err!=nil{
		return err
	}
	return nil
}
