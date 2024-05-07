package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func UpdateUser(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get("user_id")
	userId := utils.Transfer(v)
	var update UpdateParam
	if err := c.Bind(&update); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	password,_:=utils.Crypt(update.PassWord)
	resp, err := rpc.UpdateUser(ctx, &users.UpdateUserRequest{
		UserId:   userId,
		UserName: update.UserName,
		Password: password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, resp)
}
