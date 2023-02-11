package main

import (
	"context"
	douyinvideo "mydouyin/kitex_gen/douyinvideo"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *douyinvideo.CreateVideoRequest) (resp *douyinvideo.CreateVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *douyinvideo.GetFeedRequest) (resp *douyinvideo.GetFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// GetList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetList(ctx context.Context, req *douyinvideo.GetListRequest) (resp *douyinvideo.GetListResponse, err error) {
	// TODO: Your code here...
	//我改了一点
	return
}

// MGetVideoUser implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideoUser(ctx context.Context, req *douyinvideo.MGetVideoRequest) (resp *douyinvideo.MGetVideoResponse, err error) {
	// TODO: Your code here...
	return
}
