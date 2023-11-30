package main

import (
	"chatbot/components"
	"core/driver"
	"core/pkg/mail"
	"log"

	"github.com/jonboulle/clockwork"
)

var err error

func main() {
	err = driver.InitCfg()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("配置文件加载成功")

	// Todo: 处理ws服务器异常: Error during message reading: websocket: close 1006 (abnormal closure): unexpected EOF
	err := driver.InitWS()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("WebSocket连接成功")

	err = driver.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("数据库连接成功")

	err = mail.InitNotify(clockwork.NewRealClock())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("邮件系统初始化成功")

	done := make(chan struct{})
	go components.ReadMessages(driver.Conn, done)
	go func() {
		err := components.SendUpdateMessage()
		if err != nil {
			log.Println(err)
		}
	}()

	<-done
}
