package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/publishs"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func UploadVideo(ctx context.Context, c *app.RequestContext) {
	var UploadVideo UpLoadvideoParam
	if err := c.Bind(&UploadVideo); err != nil {
		hlog.Info(err)
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	userId, _ := c.Get("user_id")
	resp := new(publishs.UpLoadVideoResponse)
	if v, ok := userId.(float64); ok { //ToDo :先将接口类型转化为interface接口类型后，再去转化为int64类型，而非直接断言从interface接口类型转化为int64类型
		uid := int64(v)
		file, err := c.FormFile("file")
		if err != nil {
			SendResponse(c, errno.ConvertErr(err), nil)
			return
		}
		if resp, err = rpc.UploadVideo(ctx, &publishs.UpLoadVideoRequest{
			ContentType: UploadVideo.ContentType,
			ObjectName:  UploadVideo.ObjectName,
			BucketName:  UploadVideo.BucketName,
			UserId: uid,
			Path: file.Filename,
			Title: UploadVideo.Title,
			CoverUrl: UploadVideo.CoverUrl,
		}); err != nil {
			SendResponse(c, errno.ConvertErr(err), resp)
			return
		}
		SendResponse(c,errno.Success,nil)
	}
}
