package components

import (
	"core/models"
	"core/services/cqhttp"
	"core/utils"
	"fmt"
	"log"
	"time"
)

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
