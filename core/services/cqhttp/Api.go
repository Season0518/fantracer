package cqhttp

import (
	"bytes"
	"core/driver"
	"core/models"
	"core/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetHttpData(accessToken string, route string, params map[string]string) ([]byte, error) {
	baseUrl := driver.Base.Bot.Http

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+route, nil)
	if err != nil {
		return []byte{}, err
	}

	// 设置HTTP请求参数
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	//读取返回值的byte字节流
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func PostMessageSendEvent(groupId int64, messageChain []models.MessageBody) error {
	router := "/send_group_msg"

	cqMessage := utils.SerializeCQCode(messageChain)
	jsonStr := fmt.Sprintf(`{"group_id": %d, "message": "%s", "auto_escape": %v}`, groupId, cqMessage, false)
	fmt.Println(jsonStr)

	baseUrl := driver.Base.Bot.Http

	resp, err := http.Post(baseUrl+router, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("发送了一条欢迎消息，响应为:%s\n", string(body))

	return nil
}

func GetGroupInfo(groupID int64, noCache bool) (models.GroupInfo, error) {
	route := "/get_group_info"
	rawData, err := GetHttpData("", route, map[string]string{
		"group_id": strconv.FormatInt(groupID, 10),
		"no_cache": strconv.FormatBool(noCache),
	})

	if err != nil {
		return models.GroupInfo{}, err
	}

	var groupInfo models.GroupInfo
	err = SerializeRespData(rawData, &groupInfo)
	if err != nil {
		return models.GroupInfo{}, err
	}

	groupInfo.InfoRetrievedAt = time.Now().Unix()

	return groupInfo, err
}

func GetMemberList(groupID int64, noCache bool) ([]models.MemberInfo, error) {
	route := "/get_group_member_list"
	rawData, err := GetHttpData("", route, map[string]string{
		"group_id": strconv.FormatInt(groupID, 10),
		"no_cache": strconv.FormatBool(noCache),
	})

	if err != nil {
		return nil, err
	}

	var groupMembers []models.MemberInfo
	err = SerializeRespData(rawData, &groupMembers)
	if err != nil {
		return nil, err
	}

	return groupMembers, err
}

func SetGroupAddRequest(joinEvent models.GroupJoinEvent, approve bool, reason string) error {
	route := "/set_group_add_request"
	_, err := GetHttpData("", route, map[string]string{
		"flag":     joinEvent.Flag,
		"sub_type": joinEvent.SubType,
		"approve":  strconv.FormatBool(approve),
		"reason":   reason,
	})

	if err != nil {
		return err
	}

	return nil
}

func GetStrangerInfo(userID int64, noCache bool) (models.UserInfo, error) {
	route := "/get_stranger_info"
	rawData, err := GetHttpData("", route, map[string]string{
		"user_id":  strconv.FormatInt(userID, 10),
		"no_cache": strconv.FormatBool(noCache),
	})
	if err != nil {
		return models.UserInfo{}, err
	}

	var UserInfo models.UserInfo
	err = SerializeRespData(rawData, &UserInfo)

	if err != nil {
		return models.UserInfo{}, err
	}

	return UserInfo, nil
}
