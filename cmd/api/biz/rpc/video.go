package rpc

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/kitex_gen/douyinvideo/videoservice"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"
	"mydouyin/pkg/mw"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideo() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := videoservice.NewClient(
		consts.VideoServiceName,
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
	videoClient = c
}

//publish video create video info
func PublishVideo(ctx context.Context, req *douyinvideo.CreateVideoRequest) error {
	resp, err := videoClient.CreateVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

//GetFeed get feed by time
func GetFeed(ctx context.Context, req *douyinvideo.GetFeedRequest) (feed []apimodel.Video, next_time int64, err error) {
	resp, err := videoClient.GetFeed(ctx, req)
	if err != nil {
		return nil, time.Now().Unix(), err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, time.Now().Unix(), errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	feed = make([]apimodel.Video, 0, 30)
	favorites := make([]*douyinfavorite.Favorite, 0)
	for _, rpc_video := range resp.VideoList {
		favorite := new(douyinfavorite.Favorite)
		favorite.UserId = req.UserId
		favorite.VideoId = rpc_video.VideoId
		favorites = append(favorites, favorite)
	}
	isFavorites, err := favoriteClient.GetIsFavorite(ctx, &douyinfavorite.GetIsFavoriteRequest{FavoriteList: favorites})

	if err != nil {
		return nil, time.Now().Unix(), err
	}

	if len(resp.VideoList) != len(isFavorites.IsFavorites) {
		return nil, time.Now().Unix(), errno.ServiceErr
	}

	for i := 0; i < len(resp.VideoList); i++ {
		r, err := userClient.MGetUser(ctx, &douyinuser.MGetUserRequest{UserIds: []int64{resp.VideoList[i].Author}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		author := apimodel.PackUser(r.Users[0])
		video := apimodel.PackVideo(resp.VideoList[i])
		video.Author = *author
		video.IsFavorite = isFavorites.IsFavorites[i]
		feed = append(feed, *video)
	}
	next_time = resp.NextTime
	return feed, next_time, nil
}

//GetPublishList get video list by author
func GetPublishList(ctx context.Context, req *douyinvideo.GetListRequest) (video_list []apimodel.Video, err error) {
	resp, err := videoClient.GetList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	video_list = make([]apimodel.Video, 0, 50)
	for _, rpc_video := range resp.VideoList {
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
