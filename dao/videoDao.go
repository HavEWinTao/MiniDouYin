package dao

import (
	"errors"
	"fmt"
	"mini-douyin/db"
	. "mini-douyin/models"
)

func SaveVideo(video VideoDao) error {
	DB := db.GetDB()
	DB.Create(&video)
	if DB.Error != nil {
		return errors.New("向数据库中保存视频信息失败")
	}
	return nil
}

func SelectVideoByID(id int64) ([]VideoDao, error) {
	DB := db.GetDB()
	var videos []VideoDao
	DB.Where("author_id=?", id).Find(&videos)
	if DB.Error != nil {
		fmt.Println(DB.Error.Error())
		return nil, errors.New("查询视频列表失败")
	}
	return videos, nil
}
