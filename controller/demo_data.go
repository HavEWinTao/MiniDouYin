package controller

import "mini-douyin/models"

var DemoVideos = []models.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://rcg3q4uhi.hb-bkt.clouddn.com/videos/2_share_1fde94aa2f4803329cb4be99f15bdace.mp4?e=1653931201&token=4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE:PHjJJun0UlpQQvQof1BdJbCmk-0=",
		CoverUrl:      "http://rcg3q4uhi.hb-bkt.clouddn.com/covers/2_share_1fde94aa2f4803329cb4be99f15bdace.jpg?e=1653931418&token=4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE:gHhj4QA-t47RGxASQoMzqf5qTII=",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "有趣的生物书",
	},
	{
		Id:            2,
		Author:        UserFan,
		PlayUrl:       "http://rcg3q4uhi.hb-bkt.clouddn.com/videos/2_share_995bcd6a971dd3bc255d776f843b759e.mp4?e=1653995625&token=4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE:7puZ7L2ryLqG1coFnCjxNSum0e0=",
		CoverUrl:      "http://rcg3q4uhi.hb-bkt.clouddn.com/covers/2_share_995bcd6a971dd3bc255d776f843b759e.jpg?e=1653931521&token=4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE:lx-M9ii256r3VnMdy9Ec1rtl2pk=",
		FavoriteCount: 4,
		CommentCount:  56,
		IsFavorite:    true,
		Title:         "绿西瓜",
	},
	{
		Id:            3,
		Author:        UserFan,
		PlayUrl:       "http://rcg3q4uhi.hb-bkt.clouddn.com/videos/2_share_4ad98ca7966f17c23ba4038e7971bbc6.mp4?e=1654596049&token=4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE:mK31kN8ZhPCK_XUfy6JUkfLluKw=",
		CoverUrl:      "http://rcg3q4uhi.hb-bkt.clouddn.com/covers/2_share_4ad98ca7966f17c23ba4038e7971bbc6.jpg?e=1653931453&token=4uYOoO9F304PStnjT5V2CYNLXrk8rAHDnT6E13WE:ADVqx85CQQXOaVtU8d6_OIUascg=",
		FavoriteCount: 4,
		CommentCount:  56,
		IsFavorite:    true,
		Title:         "猪猪",
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
