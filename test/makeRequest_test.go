package test

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"os"
	"path"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	//var file *multipart.FileHeader
	//待上传文件
	uploadFile := "../public/2_mmexport1647791690810.mp4"
	//新建一个缓冲，用于存放文件内容
	bodyBufer := &bytes.Buffer{}
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)
	//从bodyWriter生成fileWriter,并将文件内容写入fileWriter,多个文件可进行多次
	fileWriter, err := bodyWriter.CreateFormFile("file", path.Base(uploadFile))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file, err := os.Open(uploadFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	//不要忘记关闭打开的文件
	defer file.Close()
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println(err.Error())
	} //关闭bodyWriter停止写入数据
	bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	request := fasthttp.AcquireRequest()
	request.Header.SetContentType(contentType)
	//直接将构建好的数据放入post的body中
	request.SetBody(bodyBufer.Bytes())
	request.Header.SetMethod("POST")
	request.SetRequestURI("http://localhost:8888/")
	response := fasthttp.AcquireResponse()
	err = fasthttp.Do(request, response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(response.Body()))
}
