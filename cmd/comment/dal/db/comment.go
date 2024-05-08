package db //nolint:gofmt

import (
	"context"

	"HuaTug.com/kitex_gen/comments"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func CreateComment(ctx context.Context, comment *comments.Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}

func GetMaxId(ctx context.Context) (int64, error) {
	var Id int64
	if err := DB.WithContext(ctx).Model(&comments.Comment{}).Select("MAX(comment_id)").Scan(&Id).Error; err != nil {
		return Id, errors.Wrapf(err, "GetMaxId failed,err:%v", err)
	}
	return Id, nil
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
			return nil
	}
	if err := DB.WithContext(ctx).Model(&comments.Comment{}).Where("comment_id=? and video_id=?", req.CommentId, req.VideoId).
		Delete(&comments.Comment{}).Error; err != nil {
		return errors.Wrapf(err, "DeleteComment failed,err:%v", err)
	}
	return nil
}

func ListComment(ctx context.Context, req *comments.ListCommentRequest) ([]*comments.Comment, int64, error) {
	if req.PageSize == 0 {
		req.PageSize = 15
	}
	var total int64
	var comment []*comments.Comment
	if err := DB.WithContext(ctx).Model(&comments.Comment{}).Where("video_id=?", req.VideoId).Count(&total).
		Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).
		Find(&comment).Error;err!=nil{
			return comment,total,errors.Wrapf(err,"ListComment failed,err:%v",err)
		}
	return comment, total, nil
}
