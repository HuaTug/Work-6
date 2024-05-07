package handlers

import (
	"context"
	"errors"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/comments"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func CreateComment(ctx context.Context, c *app.RequestContext) {
	var create CreateCommentParam
	var userId int64
	if v, flag := c.Get("user_id"); !flag {
		SendResponse(c, errors.New("Fail to Get the user_id"), nil)
	}else{
		userId=utils.Transfer(v)
	}
	if err := c.Bind(&create); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	resp, err := rpc.CreateComment(ctx, &comments.CreateCommentRequest{
		VideoId:    create.VideoId,
		IndexId:    create.IndexId,
		ActionType: create.ActionType,
		Comment:    create.Comment,
		UserId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
