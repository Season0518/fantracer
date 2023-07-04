package utils

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var cfg *goconfig.ConfigFile

// var err error

func ReadMiraiBaseUrl() (string, error) {
	return cfg.GetValue("Mirai", "baseurl")
}

func ReadVerifyKey() (string, error) {
	return cfg.GetValue("Mirai", "verifykey")
}

func ReadBotAccount() (int64, error) {
	return cfg.Int64("Mirai", "account")
}

func ReadMySQLConfig() (string, string, string, error) {
	section, err := cfg.GetSection("MySQL")

	if err != nil {
		return "", "", "", err
	}

	return section["account"], section["password"], section["port"], err
}

func ReadWebsocketHost() (string, error) {
	return cfg.GetValue("CQHttp", "websocket")
}

func ReadCQBaseUrl() (string, error) {
	return cfg.GetValue("CQHttp", "baseurl")
}

func init() {
	configFilePath, err := ConvertToFullPath("config.ini")
	if err != nil {
		log.Printf("%v", err)
		return
	}

	cfg, err = goconfig.LoadConfigFile(configFilePath)
	if err != nil {
		log.Printf("配置文件读取失败,部分功能可能无法正常使用\n错误信息: %v", err)
		return
	}
}
