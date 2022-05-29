package controller

import "mini-douyin/models"

var DemoVideos = []models.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://rcg3q4uhi.hb-bkt.clouddn.com/videos/2_share_1fde94aa2f4803329cb4be99f15bdace.mp4?e=1653841853&token=4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE:3JNwy6RzbaWoGETR4W7NhqvxiAE=",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        UserFan,
		PlayUrl:       "http://43.138.10.134:8080/static/2_share_3c9f4558b4f64e3769bc8118521e09d3.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 4,
		CommentCount:  56,
		IsFavorite:    true,
		Title:         "有趣的生物书",
	},
}

var DemoComments = []models.Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = models.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var UserFan = models.User{
	Id:            2,
	Name:          "樊tastic",
	FollowCount:   10,
	FollowerCount: 10,
	IsFollow:      false,
}
