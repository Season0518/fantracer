package components

import (
	"chatbot/middleware"
	"core/models"
	"core/utils"
	"encoding/json"
	"fmt"
)

func HandleMessageEvent(rawData []byte, data map[string]interface{}) error {
	//fmt.Printf("这是一条消息,内容是: %v\n", data)
	var groupMessage models.GroupMessageEvent
	var debug int64 = 865444787

	err := json.Unmarshal(rawData, &groupMessage)
	if err != nil {
		return err
	}

	if groupMessage.GroupID == debug {
		//fmt.Printf("DEBUG: %v\n", rawData)
		fmt.Printf("DEBUG: %v\n", groupMessage.Message)
		if groupMessage.Message == "/poster" {
			var increaseEvent models.GroupIncreaseEvent
			increaseEvent.EventUniversal = groupMessage.EventUniversal
			increaseEvent.UserID = groupMessage.UserID
			increaseEvent.GroupID = groupMessage.GroupID
			data, err := json.Marshal(increaseEvent)
			err = HandleIncreaseEvent(data)
			if err != nil {
				return err
			}
		} else if groupMessage.Message == "/content" {
			mediaURL, err := utils.ReadMediaURL()
			if err != nil {
				return err
			}
			welcomeText, err := utils.ReadWelcomeText()
			if err != nil {
				return err
			}
			loadedMsg := fmt.Sprintf("目前已经载入的欢迎词是: \\n\\n%v\\n\\n目前已经载入的图片URL有: %v", welcomeText, mediaURL)

			err = middleware.PostMessageSendEvent(debug, []models.MessageBody{
				{Type: "text", Data: map[string]string{"text": loadedMsg}},
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func HandleNoticeEvent(rawData []byte, data map[string]interface{}) error {
	if data["notice_type"] == "group_increase" {
		fmt.Printf("捕获到加群信息,内容是: %v\n", data)
		err := HandleIncreaseEvent(rawData)
		if err != nil {
			//log.Println(err)
			return err
		}
	} else {
		//fmt.Printf("这是一条通知,内容是:%v\n", data)
	}
	return nil
}

func HandleIncreaseEvent(data []byte) error {
	var increaseEvent models.GroupIncreaseEvent

	err := json.Unmarshal(data, &increaseEvent)
	if err != nil {
		return err
	}

	userJoinedChan <- increaseEvent
	if err != nil {
		return err
	}

	return nil
}
