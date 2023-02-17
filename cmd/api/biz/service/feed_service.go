package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/errno"
	"time"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) GetFeed(req apimodel.GetFeedRequest, userId int64) (*apimodel.GetFeedResponse, error) {
	resp := new(apimodel.GetFeedResponse)
	var err error
	rpc_resp, err := rpc.GetFeed(s.ctx, &douyinvideo.GetFeedRequest{
		LatestTime: req.LatestTime,
		UserId:     userId,
	})
	if err != nil {
		resp.NextTime = time.Now().Unix()
		return resp, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		resp.NextTime = time.Now().Unix()
		return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	resp.VideoList = make([]apimodel.Video, 0, 30)
	favorites := make([]*douyinfavorite.Favorite, 0)
	for _, rpc_video := range rpc_resp.VideoList {
		favorite := new(douyinfavorite.Favorite)
		favorite.UserId = userId
		favorite.VideoId = rpc_video.VideoId
		favorites = append(favorites, favorite)
	}
	isFavorites, err := rpc.GetIsFavorite(s.ctx, &douyinfavorite.GetIsFavoriteRequest{FavoriteList: favorites})

	if err != nil {
		resp.NextTime = time.Now().Unix()
		return resp, err
	}

	if len(rpc_resp.VideoList) != len(isFavorites.IsFavorites) {
		resp.NextTime = time.Now().Unix()
		return resp, errno.ServiceErr
	}

	for i := 0; i < len(rpc_resp.VideoList); i++ {
		r, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_resp.VideoList[i].Author}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		author := apimodel.PackUser(r.Users[0])
		video := apimodel.PackVideo(rpc_resp.VideoList[i])
		video.Author = *author
		video.IsFavorite = isFavorites.IsFavorites[i]
		resp.VideoList = append(resp.VideoList, *video)
	}
	resp.NextTime = rpc_resp.NextTime
	return resp, nil
}
