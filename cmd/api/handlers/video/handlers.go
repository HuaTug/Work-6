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

type UpLoadvideoParam struct {
	ContentType string `json:"content_type" form:"content_type"`
	ObjectName  string `json:"object_name" form:"object_name"`
	BucketName  string `json:"bucket_name" form:"bucket_name"`
	Title       string `json:"title" form:"title"`
	CoverUrl    string `json:"cover_url" form:"cover_url"`
}

type FeedListParam struct {
	LastTime string `json:"last_time" form:"last_time"`
}

type VideoFeedListParam struct {
	AuthorId int64 `form:"author_id" `
	PageNum  int64 `form:"page_num"`
	PageSize int64 `form:"page_size"`
}

type VideoSearchParam struct {
	Keyword  string `form:"keyword"`
	PageNum  int64  `form:"page_num"`
	PageSize int64  `form:"page_size"`
	FromDate string `form:"from_date"`
	ToDate   string `form:"to_date"`
}
