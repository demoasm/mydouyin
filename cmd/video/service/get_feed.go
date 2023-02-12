package service
import(
	"context"
	"mydouyin/cmd/video/dal/db"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/cmd/video/pack"
)
type GetFeedService struct {
	ctx context.Context
}

// NewGetFeedService new GetFeedService
func NewGetFeedService(ctx context.Context) *GetFeedService {
	return &GetFeedService{ctx: ctx}
}

// GetFeedService.
func (s *GetFeedService) GetFeed(req *douyinvideo.GetFeedRequest) (int64, []*douyinvideo.Video, error) {
	videos, err := db.GetFeed(s.ctx, req.LatestTime)
	if err != nil {
		return 0, nil, err
	}
	var index int = len(videos) - 1 
	return videos[index].CreatedAt.Unix(), pack.Videos(videos) ,nil
}

