package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type mysqlConfig struct {
	Host        string
	Port        int
	User        string
	Password    string
	Database    string
	MaxIdLeConn int
	MaxOpenConn int
	Debug       bool
	IsPlural    bool
	TablePrefix string
}

var MysqlConfig *mysqlConfig

func init() {
	//程序启动的时候 就会执行init方法
	MysqlConfig = new(mysqlConfig)
	curPath, _ := os.Getwd()
	fmt.Println(curPath)
	_, err := toml.DecodeFile("../config/config.toml", &MysqlConfig)
	if err != nil {
		panic(err)
	}
}
