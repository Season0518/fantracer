package components

import (
	"core/models"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

var userJoinedChan chan models.GroupIncreaseEvent

func ReadMessages(c *websocket.Conn, done chan struct{}) {
	defer close(done)

	targetGroup := []int64{688718641, 865444787}

	userJoinedChan = make(chan models.GroupIncreaseEvent, 100)

	go SendWelcomeMessage(targetGroup, userJoinedChan)

	for {
		_, rawData, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		var data map[string]interface{}
		err = json.Unmarshal(rawData, &data)
		if err != nil {
			log.Println(err)
		}

		switch data["post_type"] {
		case "message":
			err = HandleMessageEvent(rawData, data)
			if err != nil {
				log.Println(err)
			}
		case "notice":
			err = HandleNoticeEvent(rawData, data)
			if err != nil {
				log.Println(err)
			}
		default:
			//fmt.Printf("其他类型通知,内容是:%v\n", data)
		}
	}
}
