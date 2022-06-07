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
	"strconv"
	"strings"
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
	}
	//视频文件保存到七牛云
	err = utils.UploadVideo(data, user.Id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//封面保存到七牛云
	err = utils.UploadCover(finalName)
	if err != nil {
		fmt.Println(err.Error())
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
	token := c.Query("token")
	//查询用户
	user := usersLoginInfo[token]
	userID := c.Query("user_id")
	userIDInt64, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Println("userID转换失败")
	}
	videoDaoList, err := dao.SelectVideoByID(userIDInt64)
	if err != nil {
		fmt.Println(err.Error())
	}
	videoList := make([]models.Video, len(videoDaoList), len(videoDaoList))
	for i := 0; i < len(videoDaoList); i++ {
		videoList[i].Id = videoDaoList[i].VideoId
		videoList[i].Author = user
		videoList[i].PlayUrl = utils.GetVideoUrl(videoDaoList[i].PlayUrl)
		fmt.Println("playUrl: ", videoList[i].PlayUrl)
		videoList[i].CoverUrl = utils.GetCoverUrl(videoDaoList[i].CoverUrl)
		videoList[i].FavoriteCount = videoDaoList[i].FavoriteCount
		videoList[i].CommentCount = videoDaoList[i].CommentCount
		//TODO:is_favorite
		videoList[i].IsFavorite = true
		videoList[i].Title = videoDaoList[i].Title
	}
	c.JSON(http.StatusOK, models.VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
