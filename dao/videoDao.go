package dao

import (
	//"mini-douyin/db"
	"mini-douyin/db"
	. "mini-douyin/models"
)

func SaveVideo(video VideoDao) error {
	DB := db.GetDB()
	err := DB.Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}
