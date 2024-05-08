package service

import (
	"context"

	"HuaTug.com/cmd/comment/dal/db"
	"HuaTug.com/kitex_gen/comments"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
)

type ListCommentService struct {
	ctx context.Context
}

func NewListCommentService(ctx context.Context) *ListCommentService {
	return &ListCommentService{ctx: ctx}
}

func (v *ListCommentService) ListComment(ctx context.Context, req *comments.ListCommentRequest) (*comments.ListCommentResponse, error) {
	var comment []*comments.Comment
	var total int64
	resp := new(comments.ListCommentResponse)
	var err error
	comment, total, err = db.ListComment(v.ctx, req)
	if err != nil {
		return resp, errors.WithMessage(err, "dao.ListComment failed")
	} else {
		resp.Total = total
		resp.Comments = comment
		resp.Code = consts.StatusOK
		resp.Msg = "Success to List Comments"
	}
	return resp, nil
}
