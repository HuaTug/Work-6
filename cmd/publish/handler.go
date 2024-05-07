package main

import (
	"context"

	"HuaTug.com/cmd/publish/service"
	"HuaTug.com/kitex_gen/publishs"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type PublishServiceImpl struct{}

func (s *PublishServiceImpl) UploadVideo(ctx context.Context, req *publishs.UpLoadVideoRequest) (resp *publishs.UpLoadVideoResponse, err error) {
	resp = new(publishs.UpLoadVideoResponse)
	err = service.NewUploadService(ctx).UploadFile(req)
	if err != nil {
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to Upload file"
		return resp, err
	}

	return resp, nil
}

func (s *PublishServiceImpl) VideoCreate(ctx context.Context, req *publishs.VideoCreateRequest) (resp *publishs.VideoCreateResponse, err error) {
	resp = new(publishs.VideoCreateResponse)
	err = service.NewVideoCreateService(ctx).VideoCreate(ctx, req)
	if err != nil {
		hlog.Info(err)
		resp.Code=consts.StatusBadRequest
		resp.Msg="Fail to Create Video"
		return resp,err
	}
	return resp,nil
}
