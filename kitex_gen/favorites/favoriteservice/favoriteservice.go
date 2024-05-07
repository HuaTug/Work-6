// Code generated by Kitex v0.9.1. DO NOT EDIT.

package favoriteservice

import (
	favorites "HuaTug.com/kitex_gen/favorites"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"FavoriteService": kitex.NewMethodInfo(
		favoriteServiceHandler,
		newFavoriteServiceFavoriteServiceArgs,
		newFavoriteServiceFavoriteServiceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListFavorite": kitex.NewMethodInfo(
		listFavoriteHandler,
		newFavoriteServiceListFavoriteArgs,
		newFavoriteServiceListFavoriteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	favoriteServiceServiceInfo                = NewServiceInfo()
	favoriteServiceServiceInfoForClient       = NewServiceInfoForClient()
	favoriteServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return favoriteServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return favoriteServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorites.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "favorites",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteServiceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorites.FavoriteServiceFavoriteServiceArgs)
	realResult := result.(*favorites.FavoriteServiceFavoriteServiceResult)
	success, err := handler.(favorites.FavoriteService).FavoriteService(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteServiceArgs() interface{} {
	return favorites.NewFavoriteServiceFavoriteServiceArgs()
}

func newFavoriteServiceFavoriteServiceResult() interface{} {
	return favorites.NewFavoriteServiceFavoriteServiceResult()
}

func listFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorites.FavoriteServiceListFavoriteArgs)
	realResult := result.(*favorites.FavoriteServiceListFavoriteResult)
	success, err := handler.(favorites.FavoriteService).ListFavorite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceListFavoriteArgs() interface{} {
	return favorites.NewFavoriteServiceListFavoriteArgs()
}

func newFavoriteServiceListFavoriteResult() interface{} {
	return favorites.NewFavoriteServiceListFavoriteResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteService(ctx context.Context, req *favorites.FavoriteRequest) (r *favorites.FavoriteResponse, err error) {
	var _args favorites.FavoriteServiceFavoriteServiceArgs
	_args.Req = req
	var _result favorites.FavoriteServiceFavoriteServiceResult
	if err = p.c.Call(ctx, "FavoriteService", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFavorite(ctx context.Context, req *favorites.ListFavoriteRequest) (r *favorites.ListFavoriteResponse, err error) {
	var _args favorites.FavoriteServiceListFavoriteArgs
	_args.Req = req
	var _result favorites.FavoriteServiceListFavoriteResult
	if err = p.c.Call(ctx, "ListFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}