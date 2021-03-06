package controller

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/models"
	"net/http"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, models.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, models.UserListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		UserList: []models.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, models.UserListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		UserList: []models.User{DemoUser},
	})
}
