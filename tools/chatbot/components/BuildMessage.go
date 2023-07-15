package components

import (
	"core/models"
	"core/utils"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func BuildWelComeMessage(newComers []models.GroupIncreaseEvent) ([]models.MessageBody, error) {
	var messageChain []models.MessageBody

	for _, user := range newComers {
		messageChain = append(messageChain, models.MessageBody{
			Type: "at",
			Data: map[string]string{
				"qq": strconv.FormatInt(user.UserID, 10),
			},
		})
	}

	welcomeMsg, err := utils.ReadWelcomeText()
	if err != nil {
		return nil, err
	}

	messageChain = append(messageChain, models.MessageBody{
		Type: "text",
		Data: map[string]string{
			"text": welcomeMsg,
		},
	})

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	posterURLs, err := utils.ReadMediaURL()
	if err != nil {
		return nil, err
	}

	messageChain = append(messageChain, models.MessageBody{
		Type: "image",
		Data: map[string]string{
			"file":    posterURLs[r.Intn(len(posterURLs))],
			"subType": "0",
		},
	})

	return messageChain, nil
}

func BuildUpdateMessage(info models.PostInfo, record models.PostRecord) ([]models.MessageBody, error) {
	var messageChain []models.MessageBody

	messageChain = append(messageChain, models.MessageBody{
		Type: "at",
		Data: map[string]string{
			"qq": "all",
		},
	})
	messageChain = append(messageChain, models.MessageBody{
		Type: "text",
		Data: map[string]string{
			"text": fmt.Sprintf("公主在%s发布了内容！\\n\\n%s\\n\\n快点击%s围观吧！", info.Platform, func(s string) string {
				r := []rune(s)
				if len(r) < 10 {
					return string(r)
				} else {
					return string(r[:10]) + "..."
				}
			}(record.Text), record.Refer),
			"subType": "0",
		},
	})
	return messageChain, nil
}

func BuildFailedMessage(err error) ([]models.MessageBody, error) {
	var messageChain []models.MessageBody

	messageChain = append(messageChain, models.MessageBody{
		Type: "text",
		Data: map[string]string{
			"text":    fmt.Sprintf("程序抛出异常: %v", err),
			"subType": "0",
		},
	})

	return messageChain, nil
}
