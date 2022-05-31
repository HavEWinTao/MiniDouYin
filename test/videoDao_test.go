package test

import (
	"fmt"
	"math/rand"
	"mini-douyin/dao"
	"mini-douyin/models"
	"os"
	"testing"
)

func TestSaveVideo(t *testing.T) {
	curPath, _ := os.Getwd()
	fmt.Println(curPath)
	user := models.User{
		Id:            2,
		Name:          "fanhongtao",
		FollowCount:   5,
		FollowerCount: 10,
		IsFollow:      true,
	}
	//nowTime := time.Now().Format("2006-01-02 15:04:05")
	video := models.VideoDao{
		VideoId:       rand.Int63(),
		AuthorId:      user.Id,
		PlayUrl:       "http://43.138.10.134:8080/static/",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "test title",
	}
	err := dao.SaveVideo(video)
	if err != nil {
		return
	}
}

func TestSelectVideoByID(t *testing.T) {
	var userId int64
	userId = 2
	VideoDaoList, err := dao.SelectVideoByID(userId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < len(VideoDaoList); i++ {
		fmt.Println(VideoDaoList[i])
	}
}
