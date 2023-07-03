package main

import (
	"chatbot/components"
	"core/driver"
	"log"
)

func main() {
	//var test models.GroupIncreaseEvent
	//test.UserID = 2185154974

	//fmt.Println(utils.SerializeCQCode(components.BuildWelComeMessage([]models.GroupIncreaseEvent{test})))
	err := driver.InitWS()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go components.ReadMessages(driver.Conn, done)

	<-done
}
