// Code generated by Kitex v0.9.1. DO NOT EDIT.

package favoriteservice

import (
	favorites "HuaTug.com/kitex_gen/favorites"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteService(ctx context.Context, req *favorites.FavoriteRequest, callOptions ...callopt.Option) (r *favorites.FavoriteResponse, err error)
	ListFavorite(ctx context.Context, req *favorites.ListFavoriteRequest, callOptions ...callopt.Option) (r *favorites.ListFavoriteResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kFavoriteServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) FavoriteService(ctx context.Context, req *favorites.FavoriteRequest, callOptions ...callopt.Option) (r *favorites.FavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteService(ctx, req)
}

func (p *kFavoriteServiceClient) ListFavorite(ctx context.Context, req *favorites.ListFavoriteRequest, callOptions ...callopt.Option) (r *favorites.ListFavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListFavorite(ctx, req)
}
