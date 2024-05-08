package service

import (
	"context"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/users"
	"github.com/pkg/errors"
)

type LoginuserService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *LoginuserService {
	return &LoginuserService{ctx: ctx}
}

func (v *LoginuserService) LoginUsers(req *users.LoginUserResquest) (user users.User, err error) {
	if user, err, _ = db.CheckUser(v.ctx, req.UserName, req.Password); err != nil {
		return user, errors.WithMessage(err, "dao.CheckUser failed")
	}
	return user, nil
}
