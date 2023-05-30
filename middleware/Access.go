package middleware

import (
	"encoding/json"
	"fantracer/models"
	miraiRequest "fantracer/models/mirai/Request"
	miraiResponse "fantracer/models/mirai/Response"
	"fantracer/utils"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func FetchGroupMembers(sessionKey string,group int,memberIds int) ([]models.Member,error) {
	router := "/latestMemberList"

	baseUrl,err := utils.ReadBaseUrl() 
	if err != nil {
		log.Panicln("baseUrl非法, 无法链接mirai-http服务")
	}

	requestBody := miraiRequest.Universial {
		SessionKey: sessionKey,
		Target: group,
		MemberIds: memberIds,
	}

	// 构建 URL
	url, err := url.Parse(baseUrl+router)
	if err != nil {
		return []models.Member{},fmt.Errorf("url解析失败: %v", err)
	}

	// 设置查询参数
	query := url.Query()
	query.Set("sessionKey", requestBody.SessionKey)
	query.Set("target", strconv.Itoa(requestBody.Target))
	query.Set("memberIds", strconv.Itoa(requestBody.MemberIds))

	url.RawQuery = query.Encode()	
	
	invaildParam := "memberIds=0"
	requestUrl := strings.Replace(url.String(),invaildParam,"memberIds",-1)

	response, err := http.Get(requestUrl)
	if err != nil {
		return []models.Member{},fmt.Errorf("HTTP request failed: %v", err)
	}
	defer response.Body.Close()

	//解析响应体
	var latestMemberList miraiResponse.LatestMemberList
	err = json.NewDecoder(response.Body).Decode(&latestMemberList)
	if err != nil {
		fmt.Println(latestMemberList.Data)
		return []models.Member{}, fmt.Errorf("failed to decode response body: %v", err)
	}
	
	return latestMemberList.Data, nil		
}