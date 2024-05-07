package service

import (
	"context"
	"sync"

	"HuaTug.com/cache"
	relation "HuaTug.com/cache"
	"HuaTug.com/cmd/favorite/dal/db"
	"HuaTug.com/kitex_gen/favorites"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type FavoriteService struct {
	ctx context.Context
}

const (
	add = int64(1)
	sub = int64(-1)
)

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{ctx: ctx}
}

func (s *FavoriteService) Favorite(ctx context.Context, req *favorites.FavoriteRequest) (*favorites.FavoriteResponse, error) {
	var wg sync.WaitGroup
	resp := new(favorites.FavoriteResponse)
	key := "favorite_id"
	Id := cache.GenerateID(key)
	videoId, _ := cache.CacheGetCommentVideo(req.CommentId)
	favorite := &favorites.Favorite{
		FavoriteId: Id,
		VideoId:    videoId,
		UserId:     req.UserID,
		CommentId:  req.CommentId,
		VideoType:  req.VideoType,
		//这里的VideoType是一种标识，用于标记是否要去对视频进行点赞 如果为1，则表示要对该视频进行点赞 否则不对视频点赞
	}
	/* 	if err != nil {
		hlog.Info(err)
	} */
	var err error
	errChan := make(chan error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = db.FavoriteAction(s.ctx, favorite)
		errChan <- err
	}()

	//ToDo:实现对视频的点赞缓存操作
	go func() {
		defer wg.Done()
		relation.CacheChangeUserCount(2, add, "like")
	}()
	wg.Wait()
	if err, ok := <-errChan; ok && err != nil {
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to Favorite"
		return resp, err
	}
	resp.Code = consts.StatusOK
	resp.Msg = "Success to Favorite"
	return resp, nil
}
