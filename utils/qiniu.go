package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"time"
)

// MyPutRet 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func UploadVideo(file *multipart.FileHeader, userID int64) {
	accessKey := "4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE"
	secretKey := "GhkvA1-NjMyYiyjNSUvFFy73R9evynSBAMWf3znH"
	bucket := "douyin-mini"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuabei
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	finalName := fmt.Sprintf("%d_%s", userID, file.Filename)
	key := "videos/" + finalName
	data, err := file.Open()
	if err != nil {
		return
	}
	defer func(data multipart.File) {
		err := data.Close()
		if err != nil {

		}
	}(data)
	var dataBytes []byte = make([]byte, file.Size)
	_, err = data.Read(dataBytes)
	if err != nil {
		return
	}
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(dataBytes), file.Size, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}

func GetVideoUrl(filename string) string {
	accessKey := "4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE"
	secretKey := "GhkvA1-NjMyYiyjNSUvFFy73R9evynSBAMWf3znH"
	domain := "http://rcg3q4uhi.hb-bkt.clouddn.com"
	key := "videos/" + filename
	mac := qbox.NewMac(accessKey, secretKey)
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	return privateAccessURL
}
