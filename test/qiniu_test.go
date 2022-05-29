package test

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"mini-douyin/utils"
	"testing"
)

func TestUploadVideo(t *testing.T) {
	fasthttp.ListenAndServe(":8888", func(ctx *fasthttp.RequestCtx) {
		//根据参数名获取上传的文件
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			err.Error()
			return
		}
		utils.UploadVideo(fileHeader)
		fmt.Println(fileHeader.Filename, " upload!")
		return
	})
}

func TestGetVideoUrl(t *testing.T) {
	filename := "2_mmexport1647791690810.mp4"
	videoUrl := utils.GetVideoUrl(filename)
	fmt.Println(videoUrl)
}
