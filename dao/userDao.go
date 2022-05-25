package dao

import (
	"mini-douyin/db"
	. "mini-douyin/models"
)

// Register 用户注册
func Register(user UserDao) error {
	DB := db.GetDB()
	if err := DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindOneSimple 查找基本信息
func FindOneSimple(username string) error {
	DB := db.GetDB()
	if err := DB.Where("user_name = ?", username).Error; err != nil {
		return err
	}
	return nil
}
