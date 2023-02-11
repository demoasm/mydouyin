package service

import (
	"context"
	"mydouyin/cmd/video/dal/db"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/cmd/video/pack"
)

type GetListService struct {
	ctx context.Context
}

// NewGetFeedService new GetFeedService
func NewGetListService(ctx context.Context) *GetListService {
	return &GetListService{ctx: ctx}
}

// GetFeedService.
func (s *GetListService) GetList(req *douyinvideo.GetListRequest) ([]*douyinvideo.Video, error) {
	videos, err := db.MGetVideosbyAuthor(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	} 
	return pack.Videos(videos) ,nil
}
