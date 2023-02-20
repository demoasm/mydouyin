// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relation "mydouyin/kitex_gen/relation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateRelation(ctx context.Context, req *relation.CreateRelationRequest, callOptions ...callopt.Option) (r *relation.BaseResp, err error)
	DeleteRelation(ctx context.Context, req *relation.DeleteRelationRequest, callOptions ...callopt.Option) (r *relation.BaseResp, err error)
	GetFollow(ctx context.Context, req *relation.GetFollowListRequest, callOptions ...callopt.Option) (r *relation.GetFollowListResponse, err error)
	GetFollower(ctx context.Context, req *relation.GetFollowerListRequest, callOptions ...callopt.Option) (r *relation.GetFollowerListResponse, err error)
	GetFriend(ctx context.Context, req *relation.GetFriendRequest, callOptions ...callopt.Option) (r *relation.GetFriendResponse, err error)
	ValidIfFollowRequest(ctx context.Context, req *relation.ValidIfFollowRequest, callOptions ...callopt.Option) (r *relation.ValidIfFollowResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
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

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) CreateRelation(ctx context.Context, req *relation.CreateRelationRequest, callOptions ...callopt.Option) (r *relation.BaseResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateRelation(ctx, req)
}

func (p *kRelationServiceClient) DeleteRelation(ctx context.Context, req *relation.DeleteRelationRequest, callOptions ...callopt.Option) (r *relation.BaseResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteRelation(ctx, req)
}

func (p *kRelationServiceClient) GetFollow(ctx context.Context, req *relation.GetFollowListRequest, callOptions ...callopt.Option) (r *relation.GetFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollow(ctx, req)
}

func (p *kRelationServiceClient) GetFollower(ctx context.Context, req *relation.GetFollowerListRequest, callOptions ...callopt.Option) (r *relation.GetFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollower(ctx, req)
}

func (p *kRelationServiceClient) GetFriend(ctx context.Context, req *relation.GetFriendRequest, callOptions ...callopt.Option) (r *relation.GetFriendResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriend(ctx, req)
}

func (p *kRelationServiceClient) ValidIfFollowRequest(ctx context.Context, req *relation.ValidIfFollowRequest, callOptions ...callopt.Option) (r *relation.ValidIfFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ValidIfFollowRequest(ctx, req)
}
