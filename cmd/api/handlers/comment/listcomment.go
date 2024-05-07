package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/comments"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func ListComment(ctx context.Context, c *app.RequestContext) {
	var list ListCommentParam
	if err := c.Bind(&list); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	if resp, err := rpc.ListComment(ctx, &comments.ListCommentRequest{
		PageNum:  list.PageNum,
		PageSize: list.PageSize,
		VideoId:  list.VideoId,
	}); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	} else {
		SendResponse(c, errno.Success, resp)
	}
}
