package test

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"mini-douyin/utils"
	"testing"
)

func TestUploadVideo(t *testing.T) {
	err := fasthttp.ListenAndServe(":8888", func(ctx *fasthttp.RequestCtx) {
		//根据参数名获取上传的文件
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			err.Error()
			return
		}
		utils.UploadVideo(fileHeader, 2)
		fmt.Println(fileHeader.Filename, " upload!")
		return
	})
	if err != nil {
		return
	}
}

func TestGetVideoUrl(t *testing.T) {
	filename := "2_share_995bcd6a971dd3bc255d776f843b759e.mp4"
	videoUrl := utils.GetVideoUrl(filename)
	fmt.Println(videoUrl)
}

func TestGetCoverUrl(t *testing.T) {
	filename := "2_share_995bcd6a971dd3bc255d776f843b759e.jpg"
	videoUrl := utils.GetCoverUrl(filename)
	fmt.Println(videoUrl)
}
