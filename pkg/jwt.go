package config

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
)

var (
	AccessTokenJwtMiddleware  *jwt.HertzJWTMiddleware
	RefreshTokenJwtMiddleware *jwt.HertzJWTMiddleware

	AccessTokenExpireTime  = time.Hour * 24
	RefreshTokenExpireTime = time.Hour * 24 * 30

	AccessTokenIdentityKey  = "user_id"
	RefreshTokenIdentityKey = "user_id"
)

type LoginParam struct {
	UserName string `form:"user_name" json:"username"`
	PassWord string `form:"password"  json:"password"`
}

func AccessTokenJwtInit() {
	var err error
	AccessTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:                         []byte("access_token_key_123456"),
		TokenLookup:                 "query:token,form:token",
		Timeout:                     AccessTokenExpireTime,
		IdentityKey:                 AccessTokenIdentityKey,
		WithoutDefaultTokenHeadName: true,

		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginRequest LoginParam
			if err := c.Bind(&loginRequest); err != nil {
				return nil, err
			}
			resp, err := rpc.LoginUser(ctx, &users.LoginUserResquest{
				UserName: loginRequest.UserName,
				Password: loginRequest.PassWord,
			})
			if err != nil {
				//c.JSON(http.StatusBadRequest, "登录失败")
				hlog.Info(err)
				return nil, err
			}
			hlog.Info(resp.User.UserName)
			hlog.Info(resp.User.UserId)
			c.Set("user_id", resp.User.UserId)
			//生成refresh_token , 并且设置键值对映射
			/* 			_, refretoken, _ := utils2.GenerateToken(users.UserID, users.UserName)
			   			c.Set("refresh", refretoken) */
			return resp.User.UserId, nil
			//return nil, nil
		},

		// data为Authenticator返回的interface{}
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					AccessTokenIdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, message string, time time.Time) {
			hlog.CtxInfof(ctx, "Login Successfully. IP: "+c.ClientIP())
			c.Set("Access-Token", message)
		},

		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
	})

	if err != nil {
		panic(err)
	}
	hlog.Infof("Access-Token Jwt Initialized Successfully")
}

func GenerateAccessToken(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get(RefreshTokenIdentityKey)
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(v)
	c.Header("New-Access-Token", tokenString)
}

func IsAccessTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := AccessTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}
	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}
	c.Set("JWT_PAYLOAD", claims)
	identity := AccessTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(AccessTokenJwtMiddleware.IdentityKey, identity)
	}
	if !AccessTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return false
	}
	return true

}

func ExtractUserIdWhenAuthorized(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	data, exist := c.Get(AccessTokenJwtMiddleware.IdentityKey)
	if !exist {
		return nil, errors.New("Service Error")
	}
	return data, nil
}

func CovertJWTPayloadToString(ctx context.Context, c *app.RequestContext) (string, error) {
	data, err := ExtractUserIdWhenAuthorized(ctx, c)
	if err != nil {
		return ``, err
	}
	return data.(map[string]interface{})["Uid"].(string), nil
}

func RefreshTokenJwtInit() {
	var err error
	RefreshTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:                         []byte("refresh_token_key_abcdef"),
		TokenLookup:                 "query:Refresh-Token,header:Refresh-Token",
		Timeout:                     RefreshTokenExpireTime,
		IdentityKey:                 RefreshTokenIdentityKey,
		WithoutDefaultTokenHeadName: true,

		// 只在LoginHandler触发
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			uid, exist := c.Get("user_id")
			if !exist {
				return nil, errors.New("Auth Fail!")
			}
			if v, ok := uid.(int64); !ok {
				return nil, errors.New("Fail to convert")
			} else {
				return v, nil
			}
		},

		/* 		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &PayloadIdentityData{
				Uid: claims[RefreshTokenJwtMiddleware.IdentityKey].(string),
			}
		}, */

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					RefreshTokenJwtMiddleware.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, message string, time time.Time) {
			c.Set("Refresh-Token", message)
		},

		/* 		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			c.Set("user_id", data.(*PayloadIdentityData).Uid)
			return true
		}, */
	})
	if err != nil {
		panic(err)
	}
	hlog.Infof("Refresh-Token Jwt Initialized Successfully")
}

func IsRefreshTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := RefreshTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}
	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}
	c.Set("JWT_PAYLOAD", claims)
	identity := RefreshTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(RefreshTokenJwtMiddleware.IdentityKey, identity)
	}
	if !RefreshTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return false
	}
	return true
}

func GenerateRefreshToken(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get(AccessTokenIdentityKey)
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(v)
	c.Header("New-Refresh-Token", tokenString)
}
