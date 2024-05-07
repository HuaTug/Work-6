package main

import (
	"context"

	"HuaTug.com/cmd/relation/service"
	"HuaTug.com/kitex_gen/relations"
)

type RelationServiceImpl struct {}

func (v *RelationServiceImpl) RelationService(ctx context.Context, req *relations.RelationServiceRequest) (resp *relations.RelationServiceResponse, err error) {
	resp = new(relations.RelationServiceResponse)
	resp, err = service.NewFollowService(ctx).Following(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

func (v *RelationServiceImpl) RelationServicePage(ctx context.Context, req *relations.RelationServicePageRequest) (resp *relations.RelationServicePageResponse, err error) {
	resp = new(relations.RelationServicePageResponse)
	resp, err = service.NewFollowListService(ctx).FollowList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
