package db

import (
	"context"

	"HuaTug.com/kitex_gen/publishs"
	"HuaTug.com/kitex_gen/videos"
)

func VideoCreate(ctx context.Context, req *publishs.VideoCreateRequest) error {
	return DB.WithContext(ctx).Model(&videos.Video{}).Create(&req.Video).Error
}
