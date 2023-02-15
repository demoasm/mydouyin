package rpc

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/relation"
	"mydouyin/kitex_gen/relation/relationservice"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"
	"mydouyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := relationservice.NewClient(
		consts.RelationServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func CreateRelation(ctx context.Context, req *relation.CreateRelationRequest) error {
	resp, err := relationClient.CreateRelation(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMessage)
	}
	return nil
}

func DeleteRelation(ctx context.Context, req *relation.DeleteRelationRequest) error {
	resp, err := relationClient.DeleteRelation(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMessage)
	}
	return nil
}

func GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest) ([]*apimodel.User, error) {
	resp, err := relationClient.GetFollower(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	if len(resp.FollowerIds) < 1 {
		return []*apimodel.User{}, nil
	}
	ur, err := userClient.MGetUser(ctx, &douyinuser.MGetUserRequest{
		UserIds: resp.FollowerIds,
	})
	if err != nil {
		return nil, err
	}
	if ur.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(ur.BaseResp.StatusCode, ur.BaseResp.StatusMessage)
	}
	res := make([]*apimodel.User, 0, 30)
	for _, rpc_user := range ur.Users {
		user := apimodel.PackUser(rpc_user)
		r, err := relationClient.ValidIfFollowRequest(ctx, &relation.ValidIfFollowRequest{
			FollowId:   user.UserID,
			FollowerId: req.FollowId,
		})
		if err != nil || r.BaseResp.StatusCode != 0 {
			continue
		}
		user.IsFollow = r.IfFollow
		res = append(res, user)
	}
	return res, nil
}

func GetFollowList(ctx context.Context, req *relation.GetFollowListRequest) ([]*apimodel.User, error) {
	resp, err := relationClient.GetFollow(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	if len(resp.FollowIds) < 1 {
		return []*apimodel.User{}, nil
	}
	ur, err := userClient.MGetUser(ctx, &douyinuser.MGetUserRequest{
		UserIds: resp.FollowIds,
	})
	if err != nil {
		return nil, err
	}
	if ur.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(ur.BaseResp.StatusCode, ur.BaseResp.StatusMessage)
	}
	res := make([]*apimodel.User, 0, 30)
	for _, rpc_user := range ur.Users {
		user := apimodel.PackUser(rpc_user)
		user.IsFollow = true
		res = append(res, user)
	}
	return res, nil
}

func GetFriendList(ctx context.Context, req *relation.GetFollowerListRequest) ([]*apimodel.FriendUser, error) {
	resp, err := relationClient.GetFollower(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	if len(resp.FollowerIds) < 1 {
		return []*apimodel.FriendUser{}, nil
	}
	ur, err := userClient.MGetUser(ctx, &douyinuser.MGetUserRequest{
		UserIds: resp.FollowerIds,
	})
	if err != nil {
		return nil, err
	}
	if ur.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(ur.BaseResp.StatusCode, ur.BaseResp.StatusMessage)
	}
	res := make([]*apimodel.FriendUser, 0, 30)
	for _, rpc_user := range ur.Users {
		user := apimodel.PackFriendUser(rpc_user)
		r, err := relationClient.ValidIfFollowRequest(ctx, &relation.ValidIfFollowRequest{
			FollowId:   user.UserID,
			FollowerId: req.FollowId,
		})
		if err != nil || r.BaseResp.StatusCode != 0 {
			continue
		}
		user.IsFollow = r.IfFollow
		res = append(res, user)
	}
	return res, nil
}
