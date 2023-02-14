package pack

import (
	"mydouyin/cmd/favorite/dal/db"
	"mydouyin/kitex_gen/douyinfavorite"
)

// Favorite pack video info
func Favorite(f *db.Favorite) *douyinfavorite.Favorite {
	if f == nil {
		return nil
	}
	return &douyinfavorite.Favorite{
		FavoriteId: int64(f.ID),
		UserId:  	int64(f.UserId),	
		VideoId: 	int64(f.VideoId),

	}
}

func FavoriteToVideoids(favorites []*db.Favorite) []int64 {
	vids := make([]int64, 0)
	for i:= 0 ;i < len(favorites); i++{
		if temp := int64(favorites[i].VideoId); temp != 0 {
			vids = append(vids, temp)
		}
	}
	return vids
}

