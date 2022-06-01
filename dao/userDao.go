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

// UsersLoginInfo 查重sql
func UsersLoginInfo(username string) (UserDao, error) {
	DB := db.GetDB()
	var user UserDao
	err := DB.Where("user_name = ?", username).First(&user).Error
	return user, err
}

// FindOneSimple 查找基本信息
func FindOneSimple(username string) (UserDao, error) {
	DB := db.GetDB()
	var user UserDao
	err := DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
