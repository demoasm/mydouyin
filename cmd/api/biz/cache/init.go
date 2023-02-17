package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"mydouyin/cmd/api/biz/apimodel"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

// Init the redis client
func Init() string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	video := &apimodel.Video{VideoID: 1, Author: apimodel.User{UserID: 1}, PlayUrl: "1", CoverUrl: "1", FavoriteCount: 1, CommentCount: 1, IsFavorite: true, Title: "lala"}
	fmt.Println(video)
	videoStr, _ := json.Marshal(video)
	rdb.Set(ctx, "key", videoStr, time.Hour)
	value := rdb.Get(ctx, "key").Val()
	var getValue apimodel.Video
	err := json.Unmarshal([]byte(value), &getValue)
	if err != nil {
		return ""
	}
	fmt.Println(getValue.Title)
	return value
}

type RedisCache struct {
}

func NewRedisCache() *RedisCache {
	return &RedisCache{}
}

//func FirstInit() {
//	videoLists, nextTime, err := rpc.GetFeed(context.Background())
//}

// Feed Get the videos feed when start
func (r *RedisCache) Updating() {
	time.Sleep(3000)
	r.Updating()
}

// 第一次启动，需要从mysql获得30个视频
// 有新上传视频更新cache	->	数据结构是什么？调什么接口
// 剔除过时视频
