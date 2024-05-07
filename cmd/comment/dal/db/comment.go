package db //nolint:gofmt

import (
	"context"
	"errors"

	"HuaTug.com/kitex_gen/comments"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sirupsen/logrus"
)

func CreateComment(ctx context.Context, comment *comments.Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}

func GetMaxId(ctx context.Context) int64 {
	var Id int64
	if err := DB.WithContext(ctx).Model(&comments.Comment{}).Select("MAX(comment_id)").Scan(&Id).Error; err != nil {
		hlog.Info(err)
		return Id
	}
	return Id
}

func Exist(ctx context.Context, req *comments.CreateCommentRequest) bool {
	var count int64
	if DB.WithContext(ctx).Model(&comments.Comment{}).Where("comment_id=? and video_id=?", req.IndexId, req.VideoId).
		Count(&count); count == 0 {
		logrus.Info("没有对应的视频评论")
		return false
	}
	return true
}

func DeleteComment(ctx context.Context, req *comments.CommentDeleteRequest) error {
	var count int64
	if DB.WithContext(ctx).Model(&comments.Comment{}).Where("comment_id=? and video_id=?", req.CommentId, req.VideoId).
		Count(&count); count == 0 {
		return errors.New("don't Exist the Comment!")
	}
	return DB.WithContext(ctx).Model(&comments.Comment{}).Where("comment_id=? and video_id=?", req.CommentId, req.VideoId).
		Delete(&comments.Comment{}).Error
}

func ListComment(ctx context.Context, req *comments.ListCommentRequest) ([]*comments.Comment, int64, error) {
	if req.PageSize == 0 {
		req.PageSize = 15
	}
	var total int64
	var comment []*comments.Comment
	err := DB.WithContext(ctx).Model(&comments.Comment{}).Where("video_id=?", req.VideoId).Count(&total).
		Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).
		Find(&comment).Error
	return comment, total, err
}
