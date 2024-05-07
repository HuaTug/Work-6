package service

import (
	"context"

	"HuaTug.com/cache"
	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
		hlog.Info("用户重复注册")
		return err
	}
	key := "user_id"
	id := cache.GenerateID(key)
	passWord, err := utils.Crypt(req.Password)
	if err != nil {
		hlog.Info(err)
		return err
	}
	return db.CreateUser(v.ctx, &users.User{
		UserId:   id,
		Password: passWord,
		UserName: req.UserName,
	})
}
