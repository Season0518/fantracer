package utils

import (
	"github.com/Unknwon/goconfig"
)

var cfg *goconfig.ConfigFile
var err error 

func ReadBaseUrl() (string,error) {
	return cfg.GetValue("Mirai","baseurl")
}

func ReadVerifyKey() (string,error) {
	return cfg.GetValue("Mirai","verifykey")
}

func ReadBotAccount() (int,error) {
	return cfg.Int("Mirai","account")
} 

func init() {
	cfg,err = goconfig.LoadConfigFile("config.ini")
	if err != nil {
		panic("配置文件读取失败")
	}
}

