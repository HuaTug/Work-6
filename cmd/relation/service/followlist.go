package service

import (
	"context"

	"HuaTug.com/cmd/relation/dal/db"
	"HuaTug.com/kitex_gen/relations"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type FollowListService struct {
	ctx context.Context
}

func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{ctx: ctx}
}

func (s *FollowListService) FollowList(ctx context.Context, req *relations.RelationServicePageRequest) (resp *relations.RelationServicePageResponse, err error) {
	resp = new(relations.RelationServicePageResponse)
	if resp.Relaitons, err = db.FollowList(ctx, req); err != nil {
		resp.Code = consts.StatusBadRequest
		resp.Msg = "Fail to FollowList"
		return resp, err
	}
	resp.Users = db.GetUser(ctx, req.UserId)
	resp.Code = consts.StatusOK
	resp.Msg = "Success to FollowList"
	return resp, nil
}
