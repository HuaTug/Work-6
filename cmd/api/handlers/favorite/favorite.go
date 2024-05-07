package handlers

import (
	"context"
	"errors"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/favorites"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func FavoriteService(ctx context.Context, c *app.RequestContext) {
	var Fav FavoriteParam
	if err := c.Bind(&Fav); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	var userId int64
	if v, ok := c.Get("user_id"); !ok {
		SendResponse(c, errors.New("Fail to Get UserId"), nil)
	} else {
		userId = utils.Transfer(v)
	}
	if resp, err := rpc.Favorite(ctx, &favorites.FavoriteRequest{
		ActionType: Fav.ActionType,
		CommentId:  Fav.CommentId,
		VideoType:  Fav.VideoType,
		UserID:     userId,
	}); err != nil {
		hlog.Info(err)
		SendResponse(c, errno.ConvertErr(err), resp)
	} else {
		SendResponse(c, errno.Success, resp)
	}
}
