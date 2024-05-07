package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/relations"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func RelationService(ctx context.Context, c *app.RequestContext) {
	var relationservice RelationParam
	if err:=c.Bind(&relationservice);err!=nil{
		SendResponse(c,errno.ConvertErr(err),nil)
	}
	var userId int64
	v,_:=c.Get("user_id")
	userId=utils.Transfer(v)
	resp:=new(relations.RelationServiceResponse)
	var err error
	resp,err=rpc.Relation(ctx,&relations.RelationServiceRequest{
		ActionType: relationservice.ActionType,
		ToUserId: relationservice.ToUserId,
		UserId: userId,
	})
	if err!=nil{
		SendResponse(c,errno.ConvertErr(err),resp)
	}
	SendResponse(c,errno.Success,resp)
}
