package service

import (
	"context"
	"sync"

	"HuaTug.com/cmd/favorite/dal/db"
	"HuaTug.com/kitex_gen/favorites"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
	wg.Add(2)
	go func() {
		defer wg.Done()
		users = db.UserExist(s.ctx, req.UserId)
	}()
	go func() {
		defer wg.Done()
		favs = db.FavoriteExist(s.ctx, req.UserId)
	}()
	wg.Wait()
	resp.Favs = favs
	resp.Users = users
	if len(users) == 0 {
		hlog.Info("users is nil")
	}
	resp.Code = consts.StatusOK
	resp.Msg = "Success to ListFavorite"
	return resp, nil
}
