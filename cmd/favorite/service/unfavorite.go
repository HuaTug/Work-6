package service

import (
	"context"
	"sync"

	"HuaTug.com/cmd/favorite/dal/db"
	"HuaTug.com/config/cache"
	"HuaTug.com/kitex_gen/favorites"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
)

type UnFavoriteService struct {
	ctx context.Context
}

func NewUnFavoriteService(ctx context.Context) *UnFavoriteService {
	return &UnFavoriteService{ctx: ctx}
}

func (s *UnFavoriteService) UnFavorite(ctx context.Context, req *favorites.FavoriteRequest) (*favorites.FavoriteResponse, error) {
	resp := new(favorites.FavoriteResponse)
	var err error
	var wg sync.WaitGroup
	VideoId, _ := cache.CacheGetCommentVideo(req.CommentId)
	toUid, err := cache.CacheGetAuthor(VideoId)
	if err != nil {
		hlog.Info(err)
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = db.UnFavoriteAction(s.ctx, req.UserID, VideoId)
	}()
	if err != nil {
		hlog.Info(err)
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to UnFavorite"
		return resp, errors.WithMessage(err,"dao.UnFavoriteAction failed")
	}
	go func() {
		defer wg.Done()
		cache.CacheChangeUserCount(toUid, sub, "unlike")
	}()
	wg.Wait()
	resp.Code = consts.StatusOK
	resp.Msg = "Success to UnFavorite"
	return resp, nil
}
