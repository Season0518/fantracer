package main

import (
	"chatbot/components"
	"core/driver"
	"log"
)

func main() {
	err := driver.InitWS()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("WebSocket连接成功")
	done := make(chan struct{})
	go components.ReadMessages(driver.Conn, done)

	<-done
}
