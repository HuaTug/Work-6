package service

import (
	"context"

	"HuaTug.com/cmd/comment/dal/db"
	"HuaTug.com/kitex_gen/comments"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{ctx: ctx}
}

func (v *DeleteCommentService) DeleteComment(ctx context.Context, req *comments.CommentDeleteRequest) (*comments.CommentDeleteResponse, error) {
	resp := new(comments.CommentDeleteResponse)
	if err := db.DeleteComment(v.ctx, req); err != nil {

		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to Delete Comment"
		return resp, errors.WithMessage(err, "dao.DeleteComment failed")
	}
	resp.Code = consts.StatusOK
	resp.Msg = "Success to Delete Comment"
	return resp, nil
}
