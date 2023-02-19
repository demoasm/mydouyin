package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"strconv"
	"time"
)

type RedisCache struct {
	Ctx         context.Context
	Rdb         *redis.Client
	HeadTime    string
	TailTime    string
	SetName     string
	IsAllLoaded bool
}

var RC *RedisCache

// Init the redis client
func Init() {
	RC = new(RedisCache)
	RC.Ctx = context.Background()
	RC.SetName = "videos"
	RC.IsAllLoaded = false
	RC.HeadTime = time.Now().Format("2006-01-02 15:04:05")
	RC.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//RC.InitGetVideos(strconv.Itoa(int(time.Now().Unix())), 1)
	//video := &apimodel.Video{VideoID: 1, Author: apimodel.User{UserID: 1}, PlayUrl: "1", CoverUrl: "1", FavoriteCount: 1, CommentCount: 1, IsFavorite: true, Title: "lala"}
	//fmt.Println(video)
	//videoStr, _ := json.Marshal(video)
	//rdb.Set(ctx, "key", videoStr, time.Hour)
	//value := rdb.Get(ctx, "key").Val()
	//var getValue apimodel.Video
	//err := json.Unmarshal([]byte(value), &getValue)
	//if err != nil {
	//	return ""
	//}
	//fmt.Println(getValue.Title)
	//return value
	//fmt.Println(RC)
	//current := time.Now().Format("200601.02150405")
	//fmt.Println(current)
	//num := RC.Rdb.ZCount(RC.Ctx, RC.SetName, strconv.Itoa(2000), current).Val()
	//fmt.Println(num)

	// read from mysql to cache routine
	// go RC.InitGetVideos(RC.SystemTime)

}

// GetVideosFromRPC 从RPC获取MySQL数据
func GetVideosFromRPC(latestTime string, userId int64) (int64, []apimodel.Video) {
	rpc_resp, err := rpc.GetFeed(RC.Ctx, &douyinvideo.GetFeedRequest{
		LatestTime: latestTime,
		UserId:     userId,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	videoList := make([]apimodel.Video, 0, 30)
	favorites := make([]*douyinfavorite.Favorite, 0)
	for _, rpc_video := range rpc_resp.VideoList {
		favorite := new(douyinfavorite.Favorite)
		favorite.UserId = userId
		favorite.VideoId = rpc_video.VideoId
		favorites = append(favorites, favorite)
	}
	isFavorites, err := rpc.GetIsFavorite(RC.Ctx, &douyinfavorite.GetIsFavoriteRequest{FavoriteList: favorites})
	if err != nil {
		log.Fatal("error in rpc GetIsFavorite")
	}
	for i := 0; i < len(rpc_resp.VideoList); i++ {
		r, err := rpc.MGetUser(RC.Ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_resp.VideoList[i].Author}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		author := apimodel.PackUser(r.Users[0])
		video := apimodel.PackVideo(rpc_resp.VideoList[i])
		video.Author = *author
		video.IsFavorite = isFavorites.IsFavorites[i]
		videoList = append(videoList, *video)
	}
	return rpc_resp.NextTime, videoList
}

// AddVideoToCache 向redis里新加入一条数据
func (cache *RedisCache) AddVideoToCache(video apimodel.Video) {
	current, _ := json.Marshal(video)
	score, _ := strconv.ParseFloat(video.UploadTime, 64)
	cache.Rdb.ZAdd(cache.Ctx, cache.SetName, &redis.Z{Score: score, Member: current})
}

// UploadVideosToCache 将一组数据放到缓存中
func (cache *RedisCache) UploadVideosToCache(videoList []apimodel.Video) {
	for i := 0; i < len(videoList); i++ {
		current, _ := json.Marshal(videoList[i])
		score, _ := strconv.ParseFloat(videoList[i].UploadTime, 64)
		cache.Rdb.ZAdd(cache.Ctx, cache.SetName, &redis.Z{Score: score, Member: current})
	}
}

// InitGetVideos 系统启动时调用，从MySQL加载到redis中
func (cache *RedisCache) InitGetVideos(latestTime string, userId int64) {
	_, lists := GetVideosFromRPC(latestTime, userId)
	cache.UploadVideosToCache(lists)
}

// AddTimeBefore 向MySQL请求新的数据
func (cache *RedisCache) AddTimeBefore(reqTime string, userId int64) {
	if !cache.IsAllLoaded {
		count := cache.Rdb.ZCount(cache.Ctx, cache.SetName, "0", reqTime).Val()
		if count < 30 {
			_, lists := GetVideosFromRPC(reqTime, userId)
			// 所有历史数据已经读完
			if len(lists) < 30 {
				cache.IsAllLoaded = true
			}
			cache.UploadVideosToCache(lists)
		}
	}
}

// GetVideosFromCache 从redis里加载视频信息
func (cache *RedisCache) GetVideosFromCache(reqTime string) []apimodel.Video {
	option := redis.ZRangeBy{Max: reqTime, Count: 30}
	values, err := cache.Rdb.ZRangeByScore(cache.Ctx, cache.SetName, &option).Result()
	if err != nil {
		log.Fatal("err in sorted set range")
	}
	videos := make([]apimodel.Video, len(values))
	for i := 0; i < len(values); i++ {
		err = json.Unmarshal([]byte(values[i]), &videos[i])
		if err != nil {
			return videos
		}
	}
	return videos
}

// 第一次启动，从mysql加载最新30个视频
// 新视频做完数据库和OSS后直接写入cache
// 用户刷新获取旧视频，逐渐累积进入cache

// TODO: 缓存清除
// when the number > 240
// set the systemTime to the middle of all cache
// in order to save the memory

// 怎么去拥有一道彩虹🌈
// 怎么去拥抱一夏天的风🌀
// 天上的星星笑地上的人⭐
// 总是不能懂不能觉得足够🌃

// 如果我爱上你的笑容🥰
// 要怎么收藏要怎么拥有😐
// 如果你快乐再不是为我🧋
// 会不会放手其实才是拥有👋

// ———《知足》
