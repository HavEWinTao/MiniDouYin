package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"mini-douyin/dao"
	"mini-douyin/models"
	"mini-douyin/utils"
	"net/http"
	"path/filepath"
	"strings"
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
	saveFile := filepath.Join("./temp/", finalName)
	//保存视频文件(temp文件夹，用来截取封面)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1, //视频文件保存失败
			StatusMsg:  err.Error(),
		})
		return
	}
	err = utils.CutCover(finalName)
	if err != nil {
		return
	}
	title := c.PostForm("title")
	coverName := strings.Split(filename, ".")[0] + ".jpg"
	coverName = fmt.Sprintf("%d_%s", user.Id, coverName)
	video := models.VideoDao{
		VideoId:       rand.Int63(),
		AuthorId:      user.Id,
		PlayUrl:       finalName,
		CoverUrl:      coverName,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		UploadTime:    time.Now(),
	}
	//视频文件保存到七牛云
	err = utils.UploadVideo(data, user.Id)
	if err != nil {
		return
	}
	//封面保存到七牛云
	err = utils.UploadCover(finalName)
	if err != nil {
		return
	}
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
