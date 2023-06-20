package components

import (
	"chatbot/middleware"
	"core/models"
	"strconv"
	"time"
)

func buildWelComeMessage(newComers []models.GroupIncreaseEvent) []models.MessageBody {
	var messageChain []models.MessageBody

	for _, user := range newComers {
		messageChain = append(messageChain, models.MessageBody{
			Type: "at",
			Data: map[string]string{
				"qq": strconv.FormatInt(user.UserID, 10),
			},
		})
	}

	messageChain = append(messageChain, models.MessageBody{
		Type: "text",
		Data: map[string]string{
			"text": "欢迎新人！\\n加入糕老师粉丝群吧！ 请多多支持年糕老师，感谢🙏\\n",
		},
	})

	messageChain = append(messageChain, models.MessageBody{
		Type: "image",
		Data: map[string]string{
			"file":    "https://i.mjj.rip/2023/06/18/14cc580006eba64c40ac1826055cd2e9.jpeg",
			"subType": "0",
		},
	})

	messageChain = append(messageChain, models.MessageBody{
		Type: "image",
		Data: map[string]string{
			"file":    "https://i.mjj.rip/2023/06/18/88dd4b7fd95f3473038083e41f8342f0.jpeg",
			"subType": "0",
		},
	})

	return messageChain
}

func SendWelcomeMessage(groupId int64, userJoinedChan chan models.GroupIncreaseEvent) error {
	var newUsers []models.GroupIncreaseEvent
	var timer *time.Timer
	var err error

	for joinInfo := range userJoinedChan {
		newUsers = append(newUsers, joinInfo)
		if timer == nil {
			timer = time.AfterFunc(30*time.Second, func() {
				messageChain := buildWelComeMessage(newUsers)
				err = middleware.PostMessageSendEvent(groupId, messageChain)
				newUsers = nil // 清空 newUsers
				timer = nil    // 清空 timer
				if err != nil {
					return
				}
			})
		}
	}

	if err != nil {
		return err
	}

	return nil
}
