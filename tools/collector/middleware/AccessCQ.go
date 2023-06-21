package middleware

import (
	"core/models"
	"core/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func FetchGroupMembers(accessToken string, groupID int64) ([]models.MemberInfo, error) {
	router := "/get_group_member_list"

	baseUrl, err := utils.ReadCQBaseUrl()
	if err != nil {
		log.Panicln("baseUrl非法, 无法链接go-cqHttp服务")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+router, nil)
	if err != nil {
		return []models.MemberInfo{}, err
	}

	// 设置HTTP请求参数
	query := req.URL.Query()
	query.Add("group_id", strconv.FormatInt(groupID, 10))
	query.Add("no_cache", strconv.FormatBool(true))
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return []models.MemberInfo{}, err
	}

	//读取返回值的byte字节流
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.MemberInfo{}, err
	}

	//序列化响应数据
	var response struct {
		models.CQUniversalResp
		Data []models.MemberInfo `json:"data,omitempty"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []models.MemberInfo{}, err
	}
	if response.RetCode != 0 && response.RetCode != 1 {
		return []models.MemberInfo{}, fmt.Errorf("%s", response.Wording)
	}

	return response.Data, nil
}
