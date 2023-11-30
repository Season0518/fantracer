package components

import (
	"core/models"
	"core/pkg/mail"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

var userJoinedChan chan models.GroupIncreaseEvent

func errorHandler(err error) error {
	err = mail.SendMail(
		"FanTracer Service Stopped",
		fmt.Sprintf("在"+time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")+"时：服务出现异常:%v", err),
	)

	if err != nil {
		return fmt.Errorf("无法与smtp服务器通信")
	}

	return err
}

func ReadMessages(c *websocket.Conn, done chan struct{}) {
	defer close(done)

	targetGroup := []int64{688718641, 865444787}
	retry := 0

	userJoinedChan = make(chan models.GroupIncreaseEvent, 100)

	go SendWelcomeMessage(targetGroup, userJoinedChan)

	for {
		if retry > 3 {
			log.Panic(errorHandler(fmt.Errorf("WebSocket连接已断开")))
		}

		_, rawData, err := c.ReadMessage()
		if err != nil {
			retry++
			continue
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
		case "request":
			err = HandleRequestEvent(rawData, data)
			if err != nil {
				log.Println(err)
			}
		default:
			//fmt.Printf("其他类型通知,内容是:%v\n", data)
		}
	}
}
