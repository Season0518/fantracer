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

func ReadMySQLConfig() (string,string,string,error) {
	section,err := cfg.GetSection("MySQL")

	if err != nil {
		return "","","",err
	}

	return section["account"],section["password"],section["port"],err
}

func init() {
	cfg,err = goconfig.LoadConfigFile("config.ini")
	if err != nil {
		panic("配置文件读取失败")
	}
}

