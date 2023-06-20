package middleware

import (
	"bytes"
	"core/models"
	"core/utils"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PostMessageSendEvent(groupId int64, messageChain []models.MessageBody) error {
	router := "/send_group_msg"

	cqMessage := utils.SerializeCQCode(messageChain)
	jsonStr := fmt.Sprintf(`{"group_id": %d, "message": "%s", "auto_escape": %v}`, groupId, cqMessage, false)
	fmt.Println(jsonStr)

	baseUrl, err := utils.ReadCQBaseUrl()
	if err != nil {
		return err
	}

	resp, err := http.Post(baseUrl+router, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))

	return nil
}
