package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini-douyin/dao"
	"mini-douyin/models"
	"mini-douyin/utils"
	"net/http"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//TODO:就是测试一下url
	videoDaoList, err := dao.SelectVideos()
	if err != nil {
		fmt.Println(err.Error())
	}
	videoList := make([]models.Video, len(videoDaoList), len(videoDaoList))
	for i := 0; i < len(videoDaoList); i++ {
		videoList[i].Id = videoDaoList[i].VideoId
		videoList[i].Author, err = dao.SelectUserByID(videoDaoList[i].AuthorId)
		if err != nil {
			fmt.Println(err.Error())
		}
		videoList[i].PlayUrl = utils.GetVideoUrl(videoDaoList[i].PlayUrl)
		fmt.Println("playUrl: ", videoList[i].PlayUrl)
		videoList[i].CoverUrl = utils.GetCoverUrl(videoDaoList[i].CoverUrl)
		videoList[i].FavoriteCount = videoDaoList[i].FavoriteCount
		videoList[i].CommentCount = videoDaoList[i].CommentCount
		//TODO:is_favorite
		videoList[i].IsFavorite = true
		videoList[i].Title = videoDaoList[i].Title
	}
	c.JSON(http.StatusOK, models.FeedResponse{
		Response:  models.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
