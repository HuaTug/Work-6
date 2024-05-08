package service

import (
	"context"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/config/cache"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/utils"
	"github.com/pkg/errors"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (v *CreateUserService) CreateUser(req *users.CreateUserRequest) error {
	var err error
	var flag bool
	//var wg sync.WaitGroup
	if _, err, flag = db.CheckUser(v.ctx, req.UserName, req.Password); !flag {
		return errors.WithMessage(err, "User duplicate registration")
	}
	key := "user_id"
	id := cache.GenerateID(key)
	passWord, err := utils.Crypt(req.Password)
	if err != nil {
		return errors.WithMessage(err, "Password fail to crypt")
	}
	return db.CreateUser(v.ctx, &users.User{
		UserId:   id,
		Password: passWord,
		UserName: req.UserName,
	})
}
