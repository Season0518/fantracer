package components

import (
	"chatbot/middleware"
	"core/models"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

var userJoinedChan chan models.GroupIncreaseEvent

func ReadMessages(c *websocket.Conn, done chan struct{}) {
	defer close(done)
	userJoinedChan = make(chan models.GroupIncreaseEvent, 100)
	go SendWelcomeMessage(688718641, userJoinedChan)

	for {
		rawData, data, err := middleware.AccessWebSocket(c)
		if err != nil {
			log.Println(err)
		}

		switch data["post_type"] {
		case "message":
			err = HandleMessageEvent(rawData, data)
			if err != nil {
				fmt.Println(err)
			}
		case "notice":
			err = HandleNoticeEvent(rawData, data)
			if err != nil {
				fmt.Println(err)
			}
		default:
			//fmt.Printf("其他类型通知,内容是:%v\n", data)
		}
	}
}
