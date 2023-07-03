package components

import (
	"core/models"
	"encoding/json"
	"fmt"
)

func HandleMessageEvent(rawData []byte, data map[string]interface{}) error {
	//fmt.Printf("这是一条消息,内容是: %v\n", data)

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
	var groupId int64

	// 设置欢迎ID
	groupId = 688718641

	err := json.Unmarshal(data, &increaseEvent)
	if err != nil {
		return err
	}

	if increaseEvent.GroupID == groupId {
		userJoinedChan <- increaseEvent
		if err != nil {
			return err
		}
	}

	return nil
}
