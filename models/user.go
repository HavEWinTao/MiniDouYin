package models

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
	IsFollow      bool
}

func (u UserDao) TableName() string {
	return "users"
}
