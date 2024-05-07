package service

import (
	"context"
	"time"

	"HuaTug.com/cache"
	"HuaTug.com/cmd/comment/dal/db"
	"HuaTug.com/kitex_gen/comments"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

func (v *CreateCommentService) CreateComment(req *comments.CreateCommentRequest) (*comments.CreateCommentResponse, error) {
	key := "comment_id"
	Id := cache.GenerateID(key)
	resp := new(comments.CreateCommentResponse)
	comment := &comments.Comment{
		CommentId: Id,
		VideoId:   req.VideoId,
		Comment:   req.Comment,
		UserId:    req.UserId,
		Time:      time.Now().Format(time.DateTime),
		IndexId:   req.IndexId,
	}
	if req.ActionType == 1 && req.IndexId != 0 {
		if flag := db.Exist(v.ctx, req); !flag {
			comment.IndexId = 0
			if err := db.CreateComment(v.ctx, comment); err != nil {
				hlog.Info(err)
				resp.Code = consts.StatusBadRequest
				resp.Msg = "Fail to Create Comment"
				return resp, err
			}
			resp.Code = consts.StatusOK
			resp.Msg = "Success to Create Comment"
			hlog.Info("新插入一条评论成功")
		} else {
			if err := db.CreateComment(v.ctx, comment); err != nil {
				hlog.Info(err)
				resp.Code = consts.StatusBadRequest
				resp.Msg = "Fail to Create Comment"
				return resp, err
			}
		}
	} else {
		if err := db.CreateComment(v.ctx, comment); err != nil {
			hlog.Info(err)
			resp.Code = consts.StatusBadRequest
			resp.Msg = "Fail to Create Comment"
			return resp, err
		}
	}
	//获取此时这张表中的最大主键值
	commentId := db.GetMaxId(v.ctx) 
	hlog.Info(commentId)
	go func() {
		err := cache.CacheSetCommentVideo(req.VideoId, commentId)
		if err != nil {
			hlog.Info(err)
		}
	}()
	return resp, nil
}
