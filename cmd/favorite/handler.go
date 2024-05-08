package main

import (
	"context"

	"HuaTug.com/cmd/favorite/service"
	"HuaTug.com/kitex_gen/favorites"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type FavoriteServiceImpl struct {
}

func (v *FavoriteServiceImpl) FavoriteService(ctx context.Context, req *favorites.FavoriteRequest) (resp *favorites.FavoriteResponse, err error) {
	resp = new(favorites.FavoriteResponse)
	if req.ActionType == 1 {
		resp, err = service.NewFavoriteService(ctx).Favorite(ctx, req)
		if err != nil {
			hlog.CtxErrorf(ctx, "service.FavoriteService failed,original error:%v", errors.Cause(err))
			hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
			return resp, err
		}
		return resp, nil
	} else {
		resp, err = service.NewUnFavoriteService(ctx).UnFavorite(ctx, req)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
}

func (v *FavoriteServiceImpl) ListFavorite(ctx context.Context, req *favorites.ListFavoriteRequest) (resp *favorites.ListFavoriteResponse, err error) {
	resp = new(favorites.ListFavoriteResponse)
	resp, err = service.NewListFavoriteService(ctx).ListFavorite(ctx, req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.ListFavorite failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)  
		return resp, err
	}
	return resp, nil
}
