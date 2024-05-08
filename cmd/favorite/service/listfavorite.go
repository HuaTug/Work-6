package service

import (
	"context"
	"sync"

	"HuaTug.com/cmd/favorite/dal/db"
	"HuaTug.com/kitex_gen/favorites"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
)

type ListFavoriteService struct {
	ctx context.Context
}

func NewListFavoriteService(ctx context.Context) *ListFavoriteService {
	return &ListFavoriteService{ctx: ctx}
}

func (s *ListFavoriteService) ListFavorite(ctx context.Context, req *favorites.ListFavoriteRequest) (*favorites.ListFavoriteResponse, error) {
	resp := new(favorites.ListFavoriteResponse)
	var favs []*favorites.Favorite
	var users []*favorites.User
	var wg sync.WaitGroup
	var err1,err2 error
	wg.Add(2)
	go func() {
		defer wg.Done()
		users,err1 = db.UserExist(s.ctx, req.UserId)
	}()
	go func() {
		defer wg.Done()
		favs ,err2= db.FavoriteExist(s.ctx, req.UserId)
	}()
	wg.Wait()
	if err1!=nil{
		return resp,errors.WithMessage(err1,"dao.UserExist failed")
	}
	if err2!=nil{
		return resp,errors.WithMessage(err2,"dao.FavoriteExist failed")
	}
	resp.Favs = favs
	resp.Users = users
	if len(users) == 0 {
		hlog.Info("users is nil")
	}
	resp.Code = consts.StatusOK
	resp.Msg = "Success to ListFavorite"
	return resp, nil
}
