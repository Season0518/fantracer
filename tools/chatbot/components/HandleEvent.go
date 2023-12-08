package components

import (
	"core/driver"
	"core/models"
	"core/services"
	"core/services/cqhttp"
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
			// 暂时弃用配置加载查询功能
			// Todo: 增加更全面的状态监控
			//mediaURL, err := utils.ReadMediaURL()
			//if err != nil {
			//	return err
			//}
			//welcomeText, err := utils.ReadWelcomeText()
			//if err != nil {
			//	return err
			//}
			//loadedMsg := fmt.Sprintf("目前已经载入的欢迎词是: \\n\\n%v\\n\\n目前已经载入的图片URL有: %v", welcomeText, mediaURL)

			//err = cqhttp.PostMessageSendEvent(debug, []models.MessageBody{
			//	{Type: "text", Data: map[string]string{"text": loadedMsg}},
			//})

			err = cqhttp.PostMessageSendEvent(debug, []models.MessageBody{
				{Type: "text", Data: map[string]string{"text": "Fantracer正常工作"}},
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
		// 当
		fmt.Printf("捕获到加群信息,内容是: %v\n", data)
		err := HandleIncreaseEvent(rawData)
		if err != nil {
			return err
		}
	} else if data["notice_type"] == "group_decrease" {
		// 当群员被管理员踢出时，拉入黑名单
		fmt.Printf("捕获到退群消息, 内容是: %v\n", data)
		err := HandleDecreaseEvent(rawData)
		if err != nil {
			return err
		}
	}
	return nil
}

func HandleRequestEvent(rawData []byte, data map[string]interface{}) error {
	if data["request_type"] == "group" {
		fmt.Printf("捕获到入群申请,内容是: %v\n", data)
		err := HandleJoinEvent(rawData)
		if err != nil {
			return err
		}
	}

	return nil
}

func HandleJoinEvent(data []byte) error {
	isUserBanned := func(joinEvent models.GroupJoinEvent) bool {
		var record []models.GroupBlackList
		err := services.QueryDB(fmt.Sprintf("user_id = %v", joinEvent.UserID), &record, driver.Engine)
		if err != nil {
			fmt.Printf("在链接数据库时出错, 错误: %v", err)
			return false
		}

		if len(record) != 0 {
			return true
		} else {
			return false
		}
	}

	var joinEvent models.GroupJoinEvent

	err := json.Unmarshal(data, &joinEvent)
	if err != nil {
		return err
	}

	userInfo, err := cqhttp.GetStrangerInfo(joinEvent.UserID, true)
	if err != nil {
		return err
	}

	// Todo: 进行权限校验，确保在自己是管理员的前提下进行操作。
	if userInfo.Level < 19 {
		// err = cqhttp.SetGroupAddRequest(joinEvent, false, "您的QQ状态异常，请联系Staff邀请进群")
		// if err != nil {
		// 	return err
		// }
		return nil
	} else if isUserBanned(joinEvent) {
		err = cqhttp.SetGroupAddRequest(joinEvent, false, "您已被Staff拉黑，有疑问请联系Staff")
		if err != nil {
			return err
		}
	} else {
		err = cqhttp.SetGroupAddRequest(joinEvent, true, "")
		if err != nil {
			return err
		}
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

func HandleDecreaseEvent(data []byte) error {
	var decreaseEvent models.GroupDecreaseEvent

	err := json.Unmarshal(data, &decreaseEvent)
	if err != nil {
		return err
	}

	if decreaseEvent.SubType != "kick" {
		return nil
	}

	blackList := models.GroupBlackList{
		GroupID:    decreaseEvent.GroupID,
		UserID:     decreaseEvent.UserID,
		Time:       decreaseEvent.Time,
		OperatorID: decreaseEvent.OperatorID,
		SubType:    decreaseEvent.SubType,
	}

	err = services.InsertDB(blackList, driver.Engine)
	if err != nil {
		return err
	}

	return nil
}
