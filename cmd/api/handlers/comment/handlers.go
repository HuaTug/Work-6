package handlers

import (
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type CreateCommentParam struct {
	VideoId    int64  `form:"video_id"`
	Comment    string `form:"comment"`
	IndexId    int64  `form:"index_id"`
	ActionType int64  `form:"action_type"`
}

type ListCommentParam struct {
	PageNum  int64 `form:"page_num"`
	PageSize int64 `form:"page_size"`
	VideoId  int64 `form:"video_id"`
}

type DeleteCommentParam struct{
	VideoId int64 `form:"video_id"`
	CommentId int64 `form:"comment_id"`
}