package db

import (
	"context"
	"mydouyin/pkg/consts"

	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (f *Favorite) TableName() string {
	return consts.FavoriteTableName
}

// CreateVideo create video info
func CreateFavorite(ctx context.Context, favorites []*Favorite) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(favorites).Error; err != nil {
			return err
		}
		for _, f := range favorites {
			if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", f.VideoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
				return err
			}
		}
		// 返回 nil 提交事务
		return nil
	})
	return nil
}

func CancleFavorite(ctx context.Context, favorites []*Favorite) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, f := range favorites {
			var favorite Favorite
			if err := tx.Where("user_id = ? and video_id = ?", f.UserId, f.VideoId).Delete(&favorite).Error; err != nil {
				return err
			}
		}
		for _, f := range favorites {
			if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", f.VideoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
				return err
			}
		}
		// 返回 nil 提交事务
		return nil
	})
	return nil
}

func QueryFavoriteById(ctx context.Context, userID int64, videoID int64) (bool, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userID, videoID).Find(&res).Error; err != nil {
		return false, err
	}
	if len(res) < 1 {
		return false, nil
	}
	return true, nil
}

func GetFavoriteList(ctx context.Context, userID int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}


