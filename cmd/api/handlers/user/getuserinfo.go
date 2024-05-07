package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get("user_id")
	userId := utils.Transfer(v)
	resp, err := rpc.GetUserInfo(ctx, &users.GetUserInfoRequest{
		UserId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), resp)
		return
	}
	SendResponse(c, errno.Success, resp)
}
