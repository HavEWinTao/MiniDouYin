package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini-douyin/dao"
	"mini-douyin/models"
	. "mini-douyin/models"
	utils_func "mini-douyin/utils"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]models.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
	"854540702@qq.comsisisisi04200521": {
		Id:            2,
		Name:          "fanhongtao",
		FollowCount:   5,
		FollowerCount: 10,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

//用户注册
func Register(c *gin.Context) {
	//获取用户名、密码
	username := c.Query("username")
	password := c.Query("password")
	//新用户查重
	if _, err := dao.UsersLoginInfo(username); err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User already exist",
			},
		})
	} else {
		//创建新用户
		//密码SHA256加密
		atomic.AddInt64(&userIdSequence, 1)
		newUser := models.UserDao{
			UserId:        userIdSequence,
			UserName:      username,
			Password:      utils_func.GetSHAEncode(password),
			FollowCount:   0,
			FollowerCount: 0,
		}
		if err := dao.Register(newUser); err == nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   userIdSequence,
				Token:    username + utils_func.GetSHAEncode(password),
			})
		}

	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	fmt.Println("用户名: ", username)
	fmt.Println("密码: ", password)

	token := username + utils_func.GetSHAEncode(password)

	user, err := dao.FindOneSimple(username)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		if utils_func.GetSHAEncode(password) == user.Password {
			usersLoginInfo[token] = User{
				Id:            user.UserId,
				Name:          user.UserName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			}
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.UserId,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Password error"},
			})
		}
	}

}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
