package components

import (
	"chatbot/middleware"
	"core/models"
	"math/rand"
	"strconv"
	"time"
)

func BuildWelComeMessage(newComers []models.GroupIncreaseEvent) []models.MessageBody {
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
			"text": "欢迎新人！\\n\\n快乐小妙招，爱上小年糕🌟\\n\\n在你面前的正是——长春首个地偶团体成员/线上线下主打反差/温柔腼腆内敛小女孩/生吃坟墓第一人/一人驯服数千igao/东北最火地偶/超绝可爱の年糕公主殿下！\\n\\n加入糕老师粉丝群吧！ 请多多支持年糕老师，感谢🙏",
		},
	})

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	posterURLs := []string{}
	messageChain = append(messageChain, models.MessageBody{
		Type: "image",
		Data: map[string]string{
			"file":    posterURLs[r.Intn(len(posterURLs))],
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
				messageChain := BuildWelComeMessage(newUsers)
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
