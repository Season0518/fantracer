package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Unknwon/goconfig"
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
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	executableDir := filepath.Dir(executablePath)
	configFilePath := filepath.Join(executableDir, "config.ini")
	cfg, err = goconfig.LoadConfigFile(configFilePath)
	if err != nil {
		//log.Println(err)
		log.Printf("配置文件读取失败,部分功能可能无法正常使用\n错误信息: %v", err)
	}
}
