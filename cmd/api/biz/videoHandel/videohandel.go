package videohandel

import (
	"bufio"
	"fmt"
	"mime/multipart"
	"mydouyin/pkg/consts"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type VideoHandel struct {
	Root         string
	RelativePath string
	Signal       chan error
	client       *oss.Client
	bucket       *oss.Bucket
}

var VH *VideoHandel

func Init() {
	VH = new(VideoHandel)
	VH.Root = "/mnt/d/Documents/demos/mydouyin/static/"
	VH.Root = consts.StaticRoot + "static/"
	VH.RelativePath = "static/"
	VH.Signal = make(chan error)

	// 初始化OSS
	// yourEndpoint填写Bucket对应的Endpoint 例https://oss-cn-hangzhou.aliyuncs.com
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。
	var err error
	VH.client, err = oss.New(consts.EndPoint, consts.AKID, consts.AKS)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称
	VH.bucket, err = VH.client.Bucket(consts.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

// 异步上传

func (vh *VideoHandel) UpLoad(file *multipart.FileHeader, videoObject string, picObject string) {
	// 获取文件流
	filepoint, err := file.Open()
	if err != nil {
		return
	}
	defer filepoint.Close()
	// 上传视频
	err = vh.bucket.PutObject(videoObject, filepoint)
	if err != nil {
		vh.Signal <- err
		return
	}
	// 截图格式
	style := "video/snapshot,t_0,f_jpg,w_800,h_600"
	// 根据视频名直接获取截图url
	signedURL, err := vh.bucket.SignURL(videoObject, oss.HTTPGet, 600, oss.Process(style))
	if err != nil {
		vh.Signal <- err
		return
	}
	pic, err := http.Get(signedURL)
	if err != nil {
		vh.Signal <- err
		return
	}
	defer pic.Body.Close()
	reader := bufio.NewReader(pic.Body)
	// 再次上传截图
	err = vh.bucket.PutObject(picObject, reader)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	hlog.Infof("upload Success")
	vh.Signal <- nil
}

func (vh *VideoHandel) UpLoadFile(file *multipart.FileHeader) (videourl, coverurl string, err error) {

	// 视频文件object名称
	name := time.Now().Format("2006-01-02-15:04:05") + ".mp4"
	// 构建Object名称
	var viewObjBuilder strings.Builder
	viewObjBuilder.WriteString("videos/")
	viewObjBuilder.WriteString(name)
	viewObjectName := viewObjBuilder.String()

	// 封面文件object名称
	cover_name := time.Now().Format("2006-01-02-15:04:05") + ".jpg"
	// 构建Object名称
	var picObjBuilder strings.Builder
	picObjBuilder.WriteString("cover/")
	picObjBuilder.WriteString(cover_name)
	picObj := picObjBuilder.String()

	// 构建前缀url
	var resultPrefixBuilder strings.Builder
	resultPrefixBuilder.WriteString("https://")
	resultPrefixBuilder.WriteString(consts.Bucket)
	resultPrefixBuilder.WriteString(".")
	resultPrefixBuilder.WriteString(consts.EndPoint)
	resultPrefixBuilder.WriteString("/")
	resultPrefix := resultPrefixBuilder.String()
	fmt.Println("***********************" + resultPrefix + viewObjectName)

	// 开启协程上传
	go vh.UpLoad(file, viewObjectName, picObj)

	return resultPrefix + viewObjectName, resultPrefix + picObj, nil

}
