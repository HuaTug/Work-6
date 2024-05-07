package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/relations"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func RelationServicePage(ctx context.Context, c *app.RequestContext) {
	var relationservice RelationPageParam
	if err := c.Bind(&relationservice); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	var userId int64
	v, _ := c.Get("user_id")
	userId = utils.Transfer(v)
	resp := new(relations.RelationServicePageResponse)
	var err error
	resp, err = rpc.RelationPage(ctx, &relations.RelationServicePageRequest{
		PageNum: relationservice.PageNum,
		PageSize: relationservice.PageSize,
		UserId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), resp)
	}
	SendResponse(c, errno.Success, resp)
}
