// Code generated by Kitex v0.9.1. DO NOT EDIT.

package videoservice

import (
	videos "HuaTug.com/kitex_gen/videos"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FeedService(ctx context.Context, req *videos.FeedServiceRequest, callOptions ...callopt.Option) (r *videos.FeedServiceResponse, err error)
	VideoFeedList(ctx context.Context, req *videos.VideoFeedListRequest, callOptions ...callopt.Option) (r *videos.VideoFeedListResponse, err error)
	VideoSearch(ctx context.Context, req *videos.VideoSearchRequest, callOptions ...callopt.Option) (r *videos.VideoSearchResponse, err error)
	VideoPopular(ctx context.Context, req *videos.VideoPopularRequest, callOptions ...callopt.Option) (r *videos.VideoPopularResponse, err error)
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
	return &kVideoServiceClient{
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

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) FeedService(ctx context.Context, req *videos.FeedServiceRequest, callOptions ...callopt.Option) (r *videos.FeedServiceResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FeedService(ctx, req)
}

func (p *kVideoServiceClient) VideoFeedList(ctx context.Context, req *videos.VideoFeedListRequest, callOptions ...callopt.Option) (r *videos.VideoFeedListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoFeedList(ctx, req)
}

func (p *kVideoServiceClient) VideoSearch(ctx context.Context, req *videos.VideoSearchRequest, callOptions ...callopt.Option) (r *videos.VideoSearchResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoSearch(ctx, req)
}

func (p *kVideoServiceClient) VideoPopular(ctx context.Context, req *videos.VideoPopularRequest, callOptions ...callopt.Option) (r *videos.VideoPopularResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoPopular(ctx, req)
}
