package service

import (
	"context"
	"sync"

	"HuaTug.com/cache"
	relation "HuaTug.com/cache"
	"HuaTug.com/cmd/favorite/dal/db"
	"HuaTug.com/kitex_gen/favorites"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
		return resp, err
	}
	go func() {
		defer wg.Done()
		relation.CacheChangeUserCount(toUid, sub, "unlike")
	}()
	wg.Wait()
	resp.Code = consts.StatusOK
	resp.Msg = "Success to UnFavorite"
	return resp, nil
}
