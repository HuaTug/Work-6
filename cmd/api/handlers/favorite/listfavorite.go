package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/favorites"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func ListFavorite(ctx context.Context, c *app.RequestContext) {
	var err error
	var userId int64
	v,_:=c.Get("user_id")
	userId=utils.Transfer(v)
	resp:=new(favorites.ListFavoriteResponse)
	if resp,err=rpc.ListFavorite(ctx,&favorites.ListFavoriteRequest{
		UserId: userId,
	});err!=nil{
		SendResponse(c,errno.ConvertErr(err),nil)
	}else{
		SendResponse(c,errno.Success,resp)
	}
}
