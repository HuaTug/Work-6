package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/comments"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func DeleteComment(ctx context.Context,c *app.RequestContext){
	var delete DeleteCommentParam
	if err:=c.Bind(&delete);err!=nil{
		SendResponse(c,errno.ConvertErr(err),nil)
	}
	resp,err:=rpc.DeleteComment(ctx,&comments.CommentDeleteRequest{
		VideoId: delete.VideoId,
		CommentId: delete.CommentId,
	})
	if err!=nil{
		SendResponse(c,errno.ConvertErr(err),nil)
		return
	}
	SendResponse(c,errno.Success,resp)
}