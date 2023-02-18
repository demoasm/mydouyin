package videohandel

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"
	"net/http"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type VideoHandel struct {
	Root         string
	RelativePath string
	Signal       chan error
	CommandQueue chan *command
	client       *oss.Client
	bucket       *oss.Bucket
}

type commandState int

const (
	begin commandState = iota
	finishCoverUpLoad
	allFinish
)

type command struct {
	videoName string
	userID    int64
	title     string
	ctx       context.Context
	state     commandState
}

var VH *VideoHandel

func Init() {
	VH = new(VideoHandel)
	VH.Root = "/home/mao/Desktop/mydouyin/static/"
	VH.Root = consts.StaticRoot + "static/"
	VH.RelativePath = "static/"
	VH.Signal = make(chan error)
	VH.CommandQueue = make(chan *command, 20)

	// 初始化OSS
	// yourEndpoint填写Bucket对应的Endpoint 例https://oss-cn-hangzhou.aliyuncs.com
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。
	var err error
	VH.client, err = oss.New(consts.EndPoint, consts.AKID, consts.AKS)
	if err != nil {
		panic(fmt.Sprintf("init videohandler error:%v", err))
	}

	// 填写存储空间名称
	VH.bucket, err = VH.client.Bucket(consts.Bucket)
	if err != nil {
		panic(fmt.Sprintf("init videohandler error:%v", err))
	}
	go VH.listen()
}

func (vh *VideoHandel) CommitCommand(VideoName string, UserID int64, Title string, ctx context.Context) {
	vh.CommandQueue <- &command{
		videoName: VideoName,
		userID:    UserID,
		title:     Title,
		ctx:       ctx,
		state:     begin,
	}
}

func (vh *VideoHandel) listen() {
	for {
		cmd := <-vh.CommandQueue
		log.Printf("[********VideoHandler********] recover command:%v", cmd)
		err := vh.execCommand(cmd)
		if err != nil || cmd.state != allFinish {
			log.Printf("[********VideoHandler********] command exec fail, error:%v", err)
			vh.CommandQueue <- cmd
		} else {
			log.Printf("[********VideoHandler********] command exec success!!!")
		}
	}
}

// 执行指令，视频上传成功后service提交指令给videohandler，handler只执行生成封面、入库等操作
func (vh *VideoHandel) execCommand(cmd *command) error {
	//执行指令，生成封面
	// 截图格式
	cover_name := "cover/" + time.Now().Format("2006-01-02-15:04:05") + ".jpg"
	switch cmd.state {
	case begin:
		style := "video/snapshot,t_0,f_jpg,w_800,h_600"
		// 根据视频名直接获取截图url
		signedURL, err := vh.bucket.SignURL(cmd.videoName, oss.HTTPGet, 600, oss.Process(style))
		if err != nil {
			return err
		}
		pic, err := http.Get(signedURL)
		if err != nil {
			return err
		}
		defer pic.Body.Close()
		reader := bufio.NewReader(pic.Body)

		err = vh.bucket.PutObject(cover_name, reader)

		if err != nil {
			return err
		}
		cmd.state = finishCoverUpLoad
		fallthrough
	case finishCoverUpLoad:
		//调rpc写库
		resp, err := rpc.CreateVideo(cmd.ctx, &douyinvideo.CreateVideoRequest{
			Author:   cmd.userID,
			PlayUrl:  cmd.videoName,
			CoverUrl: cover_name,
			Title:    cmd.title,
		})
		if err != nil {
			return err
		}
		if resp.BaseResp.StatusCode != 0 {
			return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		}
		cmd.state = allFinish
	}
	return nil
}

func (vh *VideoHandel) UpLoadVideo(data *multipart.FileHeader) (videoName string, err error) {
	// 获取文件流
	// 视频文件object名称
	var filepoint multipart.File
	filepoint, err = data.Open()
	if err != nil {
		return
	}
	defer filepoint.Close()
	// 上传视频
	videoName = "videos/" + time.Now().Format("2006-01-02-15:04:05") + ".mp4"
	err = vh.bucket.PutObject(videoName, filepoint)
	return
}
