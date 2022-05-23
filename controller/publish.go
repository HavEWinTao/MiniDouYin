package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"mini-douyin/dao"
	"mini-douyin/models"
	"net/http"
	"path/filepath"
	"time"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	//读取视频数据时产生了错误
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	//保存视频文件
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1, //视频文件保存失败
			StatusMsg:  err.Error(),
		})
		return
	}
	title := c.PostForm("title")
	video := models.VideoDao{
		VideoId:       rand.Int63(),
		AuthorId:      user.Id,
		PlayUrl:       "http://43.138.10.134:8080/static/" + finalName,
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
		UploadTime:    time.Now(),
	}
	//视频文件的相关信息保存到数据库
	if err := dao.SaveVideo(video); err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 2, //数据库信息保存失败
			StatusMsg:  err.Error(),
		})
		return
	}
	//视频长传成功啦
	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, models.VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
