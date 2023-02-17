package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	videohandel "mydouyin/cmd/api/biz/videoHandel"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/errno"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{
		ctx: ctx,
	}
}

func (s *PublishService) PublishVideo(req apimodel.PublishVideoRequest, user *apimodel.User) (*apimodel.PublishVideoResponse, error) {
	resp := new(apimodel.PublishVideoResponse)
	videourl, coverurl, err := videohandel.VH.UpLoadFile(req.Data)
	if err != nil {
		return resp, err
	}
	rpc_resp, err := rpc.CreateVideo(s.ctx, &douyinvideo.CreateVideoRequest{
		Author:   user.UserID,
		PlayUrl:  videourl,
		CoverUrl: coverurl,
		Title:    req.Title,
	})
	if err != nil {
		<-videohandel.VH.Signal
		return resp, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		<-videohandel.VH.Signal
		return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	// 接受并发错误
	err = <-videohandel.VH.Signal
	if err != nil {
		// TODO:理论上不会出现err(划掉)
		// 真的会有ERROR😅😅😅
		return resp, err
	}
	return resp, nil
}

func (s *PublishService) GetPublishList(req apimodel.GetPublishListRequest, user *apimodel.User) (*apimodel.GetPublishListResponse, error) {
	resp := new(apimodel.GetPublishListResponse)
	rpc_resp, err := rpc.GetList(s.ctx, &douyinvideo.GetListRequest{
		UserId: user.UserID,
	})
	if err != nil {
		return nil, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	resp.VideoList = make([]apimodel.Video, 0, 50)
	for _, rpc_video := range rpc_resp.VideoList {
		r, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_video.Author}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		author := apimodel.PackUser(r.Users[0])
		video := apimodel.PackVideo(rpc_video)
		video.Author = *author
		resp.VideoList = append(resp.VideoList, *video)
	}
	return resp, nil
}
