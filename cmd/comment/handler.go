package main

import (
	"context"

	"HuaTug.com/cmd/comment/service"
	"HuaTug.com/kitex_gen/comments"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type CommentServiceImpl struct {
}

func (v *CommentServiceImpl) CreateComment(ctx context.Context, req *comments.CreateCommentRequest) (*comments.CreateCommentResponse, error) {
	resp := new(comments.CreateCommentResponse)
	var err error
	resp, err = service.NewCreateCommentService(ctx).CreateComment(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.CreateComment failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		return resp, err
	}
	return resp, nil
}

func (v *CommentServiceImpl) ListComment(ctx context.Context, req *comments.ListCommentRequest) (*comments.ListCommentResponse, error) {
	resp := new(comments.ListCommentResponse)
	var err error
	resp, err = service.NewListCommentService(ctx).ListComment(ctx, req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.ListComment failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)  
		return resp, err
	}
	return resp, nil
}

func (v *CommentServiceImpl) DeleteComment(ctx context.Context, req *comments.CommentDeleteRequest) (*comments.CommentDeleteResponse, error) {
	resp := new(comments.CommentDeleteResponse)
	var err error
	resp, err = service.NewDeleteCommentService(ctx).DeleteComment(ctx, req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.DeleteComment failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)  
		return resp, err
	}
	return resp, nil
}
