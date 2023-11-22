package components

import (
	"core/driver"
	"core/models"
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

	// Todo: 现阶段不支持处理多群聊的情况，因此这里直接写死。
	welcomeMsg := driver.Base.Greeting[0].Text
	if len(welcomeMsg) == 0 {
		return nil, fmt.Errorf("欢迎词不能为空")
	}

	messageChain = append(messageChain, models.MessageBody{
		Type: "text",
		Data: map[string]string{
			"text": welcomeMsg,
		},
	})

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Todo: 欢迎海报属性修改为可选项
	posterURLs := driver.Base.Greeting[0].MediaURL
	if posterURLs == nil {
		return nil, fmt.Errorf("欢迎海报不能为空")
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
			"text":    fmt.Sprintf("公主在%s发布了内容！\\n\\n%s\\n\\n快点击%s围观吧！", info.Platform, record.Text, record.Refer),
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
