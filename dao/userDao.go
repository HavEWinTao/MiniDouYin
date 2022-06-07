package dao

import (
	"errors"
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

// SelectUserByID 根据id查找用户
func SelectUserByID(id int64) (User, error) {
	DB := db.GetDB()
	var user UserDao
	DB.Where("user_id = ?", id).Find(&user)
	if DB.Error != nil {
		var ret User
		return ret, errors.New("查询用户失败")
	}
	return User{
		Id:            user.UserId,
		Name:          user.UserName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      true,
	}, nil
}
