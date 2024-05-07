package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	jwt "HuaTug.com/cmd/mw"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Login(ctx context.Context, c *app.RequestContext) {
	var loginVar LoginParam
	var err error
	if err := c.Bind(&loginVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	jwt.AccessTokenJwtMiddleware.LoginHandler(ctx, c)
	jwt.RefreshTokenJwtMiddleware.LoginHandler(ctx, c)

	AccessToken := c.GetString("Access-Token")
	RefreshToken := c.GetString("Refresh-Token")
	resp := new(users.LoginUserResponse)

	resp, err = rpc.LoginUser(ctx, &users.LoginUserResquest{
		UserName: loginVar.UserName,
		Password: loginVar.PassWord,
	})
	hlog.Info(resp.User.UserName)
	if resp != nil {
		resp.Code = consts.StatusOK
		resp.Msg = "Login Success"
		resp.Token = AccessToken
		resp.RefreshToken = RefreshToken
	}
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
