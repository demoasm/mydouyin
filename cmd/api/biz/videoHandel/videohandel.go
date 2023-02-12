package videohandel

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type VideoHandel struct {
	Root         string
	RelativePath string
}

var VH *VideoHandel

func Init() {
	VH = new(VideoHandel)
	VH.Root = "/home/mao/Desktop/douyin/static/"
	VH.RelativePath = "static/"
}

func (vh *VideoHandel) UpLoadFile(file *multipart.FileHeader) (videourl, coverurl string, err error) {
	filepoint, err := file.Open()
	if err != nil {
		return
	}
	defer filepoint.Close()
	name := time.Now().Format("2006-01-02 15:04:05") + ".mp4"
	newfile, err := os.Create(vh.Root + "video/" + name)
	if err != nil {
		return
	}
	defer newfile.Close()

	var context []byte = make([]byte, 1024)
	for {
		n, err := filepoint.Read(context)
		newfile.Write(context[:n])
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return "", "", err
			}
		}
	}

	//截取封面
	cover_name := time.Now().Format("2006-01-02 15:04:05") + ".jpg"
	// cmd := exec.Command("ffmpeg", "-i", vh.Root+"video/"+name, "-vf", "select=eq(n,100)", "-vframes", "1", snapshotPath)
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(vh.Root+"video/"+name).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		Run()
	if err != nil {
		return
	}
	coverfile, err := os.Create(vh.Root + "cover/" + cover_name)
	if err != nil {
		return
	}
	coverfile.Write(buf.Bytes())
	return vh.RelativePath + "video/" + name, vh.RelativePath + "cover/" + cover_name, nil
}
