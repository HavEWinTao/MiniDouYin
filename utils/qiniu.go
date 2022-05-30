package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io/ioutil"
	"mime/multipart"
	"strings"
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

//七牛云的配置
var accessKey = "4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE"
var secretKey = "GhkvA1-NjMyYiyjNSUvFFy73R9evynSBAMWf3znH"
var bucket = "douyin-mini"
var domain = "http://rcg3q4uhi.hb-bkt.clouddn.com"

// UploadVideo 将视频以字节流的形式保存在七牛云
func UploadVideo(file *multipart.FileHeader, userID int64) error {
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
		return errors.New("打开文件失败")
	}
	defer func(data multipart.File) {
		err := data.Close()
		if err != nil {

		}
	}(data)
	var dataBytes []byte = make([]byte, file.Size)
	_, err = data.Read(dataBytes)
	if err != nil {
		return errors.New("读取文件数据失败")
	}
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(dataBytes), file.Size, &putExtra)
	if err != nil {
		fmt.Println(err)
		return errors.New("文件上传失败")
	}
	fmt.Println(ret)
	return nil
}

// GetVideoUrl 根据文件名获得视频的播放连接
func GetVideoUrl(filename string) string {
	key := "videos/" + filename
	mac := qbox.NewMac(accessKey, secretKey)
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	return privateAccessURL
}

// UploadCover 将封面上传到七牛云中
//保存在covers文件夹中
func UploadCover(filename string) error {
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
	filename = strings.Split(filename, ".")[0] + ".jpg"
	finalName := "./temp/" + filename
	key := "covers/" + filename
	data, err := ioutil.ReadFile(finalName)
	if err != nil {
		return errors.New("读取文件数据失败")
	}
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), int64(len(data)), &putExtra)
	if err != nil {
		fmt.Println(err)
		return errors.New("文件上传失败")
	}
	fmt.Println(ret)
	return nil
}

// GetCoverUrl 根据数据库中的视频名获得封面的连接
func GetCoverUrl(filename string) string {
	key := "covers/" + filename
	mac := qbox.NewMac(accessKey, secretKey)
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	return privateAccessURL
}
