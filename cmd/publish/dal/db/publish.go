package db

import (
	"context"

	"HuaTug.com/kitex_gen/publishs"
	"HuaTug.com/kitex_gen/videos"
	"github.com/pkg/errors"
)

func VideoCreate(ctx context.Context, req *publishs.VideoCreateRequest) error {
	if err := DB.WithContext(ctx).Model(&videos.Video{}).Create(&req.Video); err != nil {
		return errors.Wrapf(err.Error, "VideoCreate failed,err:%v", err)
	}
	return nil
}
