package components

import (
	"core/models"
	"core/services/cqhttp"
	"core/utils"
	"fmt"
	"log"
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

func SendWelcomeMessage(groupIds []int64, userJoinedChan chan models.GroupIncreaseEvent) {
	var messageChain []models.MessageBody
	var err error

	newComersMap := make(map[int64][]models.GroupIncreaseEvent)
	timerMap := make(map[int64]*time.Timer)

	for joinInfo := range userJoinedChan {
		if utils.FindElement(groupIds, joinInfo.GroupID) == -1 {
			continue
		}
		newComersMap[joinInfo.GroupID] = append(newComersMap[joinInfo.GroupID], joinInfo)

		if timerMap[joinInfo.GroupID] == nil {
			timerMap[joinInfo.GroupID] = time.AfterFunc(30*time.Second, func(groupId int64) func() {
				return func() {
					messageChain, err = BuildWelComeMessage(newComersMap[groupId])
					if err != nil {
						log.Panic("Build welcome message failed: ", err)
					}
					err = cqhttp.PostMessageSendEvent(groupId, messageChain)
					timerMap[groupId] = nil
					newComersMap[groupId] = nil

					if err != nil {
						log.Println("Send welcome message failed: ", err)
					}
				}
			}(joinInfo.GroupID))

			for value, key := range timerMap {
				fmt.Println(value, key)
			}
		}
	}
}
