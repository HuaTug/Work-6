package service

import (
	"context"

	"HuaTug.com/cache"
	"HuaTug.com/cmd/relation/dal/db"
	"HuaTug.com/kitex_gen/relations"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type FollowService struct {
	ctx context.Context
}

const (
	add = int64(1)
	sub = int64(-1)
)

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{ctx: ctx}
}

func (s *FollowService) Following(ctx context.Context, req *relations.RelationServiceRequest) (resp *relations.RelationServiceResponse, err error) {
	key := "relation_id"
	Id := cache.GenerateID(key)
	resp = new(relations.RelationServiceResponse)
	relation := &relations.Relation{
		RelationId: Id,
		FollowId: req.ToUserId,
		UserId:     req.UserId,
	}
	if req.ActionType == 1 {
		if err := db.Follow(ctx, relation); err != nil {
			resp.Code = consts.StatusBadRequest
			resp.Msg = "Fail to Follow"
			return resp, err
		}
		go cache.CacheChangeUserCount(req.UserId, add, "follow")

		go cache.CacheChangeUserCount(req.ToUserId, sub, "follower")

		resp.Code = consts.StatusOK
		resp.Msg = "Success"
		return resp, nil

	} else {
		if err := db.UnFollow(ctx, relation); err != nil {
			resp.Code = consts.StatusBadRequest
			resp.Msg = "Fail to UnFollow"
			return resp, err
		}
		go cache.CacheChangeUserCount(req.UserId, sub, "follow")

		go cache.CacheChangeUserCount(req.ToUserId, sub, "follower")

		resp.Code = consts.StatusOK
		resp.Msg = "Success"
		return resp, nil
	}
}
