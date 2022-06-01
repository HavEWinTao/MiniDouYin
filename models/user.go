package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type UserDao struct {
	UserId        int64
	UserName      string
	Password      string
	FollowCount   int64
	FollowerCount int64
}

// UserClaims 用户JWT的payload结构
type UserClaims struct {
	// 用户账号
	Uid string `json:"uid"`
	// 用户密码(SHA256加密)
	Verification string `json:"verification"`
	// 用户权限等级
	Auth int `json:"auth"`
	jwt.StandardClaims
}

func (u UserDao) TableName() string {
	return "users"
}
