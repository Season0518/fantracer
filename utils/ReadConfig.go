package utils

import (
	"github.com/Unknwon/goconfig"
)

func ReadBaseUrl() string {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil{
		panic("配置文件读取错误")
	}

	baseUrl, err := cfg.GetValue("Mirai","baseurl")
	if err != nil {
		panic("baseUrl非法, 无法链接mirai-http服务")
	}

	return baseUrl
}


func ReadVerifyKey() (string,error) {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil{
		panic("配置文件读取错误")
	}

	return cfg.GetValue("Mirai","verifykey")
}

func ReadBotAccount() (int,error) {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil{
		panic("配置文件读取错误")
	}

	return cfg.Int("Mirai","account")
} 

