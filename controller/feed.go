package controller

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/models"
	"net/http"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, models.FeedResponse{
		Response:  models.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
