package db

import (
	"context"
	"mydouyin/pkg/consts"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Author        int64  `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

// CreateVideo create video info
func CreateVideo(ctx context.Context, videos []*Video) error {
	return DB.WithContext(ctx).Create(videos).Error
}

// MGetVideos multiple get list of video info
func MGetVideos(ctx context.Context, videoIDs []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(videoIDs) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", videoIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetFeed multiple get list of video info
func GetFeed(ctx context.Context, latest_time string) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("created_at <= ?", latest_time).Limit(30).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetVideos multiple get list of video info
func MGetVideosbyAuthor(ctx context.Context, authorID int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("author = ?", authorID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
