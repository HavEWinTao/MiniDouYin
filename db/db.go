package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"mini-douyin/config"
)

var db *gorm.DB

func init() {
	var err error
	DbUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		config.MysqlConfig.User,
		config.MysqlConfig.Password,
		config.MysqlConfig.Host,
		config.MysqlConfig.Port,
		config.MysqlConfig.Database)
	db, err = gorm.Open("mysql", DbUrl)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		return
	}
	//最大打开的连接数
	db.DB().SetMaxOpenConns(config.MysqlConfig.MaxOpenConn)
	//最大空闲连接数
	db.DB().SetMaxIdleConns(config.MysqlConfig.MaxIdLeConn)
	//允许表名复数
	db.SingularTable(config.MysqlConfig.IsPlural)
	db.LogMode(config.MysqlConfig.Debug)
}

// GetDB 获取 gorm db，其他包调用此方法即可拿到 db
// 无需担心不同协程并发时使用这个 db 对象会公用一个连接，因为 db 在调用其方法时候会从数据库连接池获取新的连接
func GetDB() *gorm.DB {
	return db
}
