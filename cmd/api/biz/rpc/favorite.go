package rpc

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinfavorite/favoriteservice"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"
	"mydouyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := favoriteservice.NewClient(
		consts.FavoriteServiceName,
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
	favoriteClient = c
}

// publish video create video info
func FavoriteAction(ctx context.Context, req *douyinfavorite.FavoriteActionRequest) error {
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetFavoriteList(ctx context.Context, req *douyinfavorite.GetListRequest) (video_list []apimodel.Video, err error) {
	vids, err := favoriteClient.GetList(ctx, req)
	if err != nil {
		return nil, err
	}
	if vids.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(vids.BaseResp.StatusCode, vids.BaseResp.StatusMessage)
	}
	video_list = make([]apimodel.Video, 0, 50)
	videos, err := videoClient.MGetVideoUser(ctx, &douyinvideo.MGetVideoRequest{VideoIds: vids.VideoIds})
	if err != nil {
		return nil, err
	}
	if len(videos.Videos) < 1 {
		return make([]apimodel.Video, 0), nil
	} else {
		for _, rpc_video := range videos.Videos {
			r, err := userClient.MGetUser(ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_video.Author}})
			if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
				continue
			}
			author := apimodel.PackUser(r.Users[0])
			video := apimodel.PackVideo(rpc_video)
			video.Author = *author
			video_list = append(video_list, *video)
		}
		return video_list, nil
	}
}
