package main

import (
	"log"
	douyinvideo "mydouyin/kitex_gen/douyinvideo/videoservice"
)

func main() {
	svr := douyinvideo.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
