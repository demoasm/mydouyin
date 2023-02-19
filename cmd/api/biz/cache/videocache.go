package cache

import (
	"context"
	"encoding/json"
	"log"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinvideo"
	"strconv"
	"time"
)

type VideoCache struct {
	ctx    context.Context
	ticker time.Ticker
}

var VC *VideoCache

func initVideoCache() {
	VC = new(VideoCache)
	VC.ctx = context.Background()
	VC.ticker = *time.NewTicker(5 * time.Minute)
	// rdeisClient.Del(VC.ctx, "video_cache").Result()
	VC.pullVideoList()
	go func() {
		for {
			<-VC.ticker.C
			VC.pullVideoList()
			VC.ticker.Reset(5 * time.Minute)
		}
	}()
}

func (c *VideoCache) pullVideoList() {
	log.Println("[*********VideoCache.pullVideoList()*********]")
	videoList := make([]*apimodel.Video, 0, 120)
	lasted_time := strconv.FormatInt(time.Now().Unix(), 10)
	for i := 0; i < 3; i++ {
		r, err := rpc.GetFeed(c.ctx, &douyinvideo.GetFeedRequest{
			LatestTime: lasted_time,
			UserId:     -1,
		})
		if err != nil {
			return
		}
		if r.BaseResp.StatusCode != 0 {
			return
		}
		videoList = append(videoList, apimodel.PackVideos(r.VideoList)...)
		if len(apimodel.PackVideos(r.VideoList)) < 30 {
			break
		}
		lasted_time = strconv.FormatInt(r.NextTime, 10)
	}
	val, err := json.Marshal(videoList)
	if err != nil {
		return
	}
	_, err = rdeisClient.Set(c.ctx, "video_cache", val, time.Hour).Result()
	if err != nil {
		return
	}
}

func (c *VideoCache) GetVideoList(lasted_time string) (videoList []*apimodel.Video, hit bool, err error) {
	res, err := rdeisClient.Get(c.ctx, "video_cache").Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(res), &videoList)
	if err != nil {
		return
	}
	hit = (len(videoList) > 0)
	return
}
