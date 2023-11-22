package main

import (
	"chatbot/components"
	"core/driver"
	"log"
)

var err error

func main() {
	err = driver.InitCfg()
	if err != nil {
		log.Fatal(err)
	}

	err = driver.InitWS()
	if err != nil {
		log.Fatal(err)
	}

	err = driver.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("WebSocket连接成功")
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
