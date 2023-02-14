package rpc

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/kitex_gen/douyincomment"
	"mydouyin/kitex_gen/douyincomment/commentservice"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"
	"mydouyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var commentClient commentservice.Client

func initComment() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := commentservice.NewClient(
		consts.CommentServiceName,
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
	commentClient = c
}

// Create
func CreateComment(ctx context.Context, req *douyincomment.CreateCommentRequest) (int64, error) {
	resp, err := commentClient.CreateComment(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.CommentId, err
}

// Delete
func DeleteComment(ctx context.Context, req *douyincomment.DeleteCommentRequest) error {
	resp, err := commentClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// Get List
func GetVideoComments(ctx context.Context, req *douyincomment.GetVideoCommentsRequest) (comment_list []apimodel.Comment, err error) {
	resp, err := commentClient.GetVideoComments(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	comment_list = make([]apimodel.Comment, 0, 50)
	for _, rpc_comment := range resp.Comments {
		r, err := userClient.MGetUser(ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_comment.User}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		user := apimodel.PackUser(r.Users[0])
		comment := apimodel.PackComment(rpc_comment)
		comment.Commentor = *user
		comment_list = append(comment_list, *comment)
	}
	return
}
