// Package douyin Description: 该部分提供的抖音接口是在本地部署的。如需推送请自行搭建接口服务器
package douyin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/tidwall/gjson"
)

type Agent interface {
	RetrieveFromAgent(endpoint string) ([]byte, error)
}

// Todo: 未对直播信息进行处理，其中直播见ID可以从room_data字段中获取
type agent struct {
	StatusCode int             `json:"status_code"`
	Data       json.RawMessage `json:"data"`
}

type Api struct {
	Server Agent
}

type Post struct {
	SecUid       string
	Nickname     string
	AwemeId      string
	CreateTime   int64
	Desc         string
	Duration     int64
	PreviewTitle string
	DynamicCover string
	Cover        string
}

func NewApiServer() *Api {
	return &Api{
		Server: &agent{},
	}
}

func (a *agent) RetrieveFromAgent(endpoint string) ([]byte, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("在请求抖音代理服务器时发生异常: %v", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, a)
	if err != nil {
		return nil, fmt.Errorf("在解析抖音代理服务器响应时发生异常: %v", err)
	}

	if a.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("抖音代理服务器响应异常, Code: %d", a.StatusCode)
	}

	return a.Data, nil
}

func (a *Api) GetLiveStatus(secUserId string) (bool, error) {
	requestURL := fmt.Sprintf("http://localhost:15000/get_profile?sec_user_id=%s", secUserId)

	data, err := a.Server.RetrieveFromAgent(requestURL)
	if err != nil {
		return false, err
	}
	status := gjson.Get(string(data), "user.live_status").Int()

	return status != 0, nil
}

func (a *Api) GetLatestAweme(secUserId string) (Post, error) {
	requestURL := fmt.Sprintf("http://localhost:15000/get_post?sec_user_id=%s", secUserId)

	data, err := a.Server.RetrieveFromAgent(requestURL)
	if err != nil {
		return Post{}, err
	}

	var userPosts []Post

	//fmt.Println(string(data))
	result := gjson.ParseBytes(data)
	result.Get("aweme_list").ForEach(func(key, value gjson.Result) bool {
		userPosts = append(userPosts, Post{
			SecUid:       value.Get("author.sec_uid").String(),
			Nickname:     value.Get("author.nickname").String(),
			AwemeId:      value.Get("aweme_id").String(),
			CreateTime:   value.Get("create_time").Int(),
			PreviewTitle: value.Get("preview_title").String(),
			Desc:         value.Get("desc").String(),
			Duration:     value.Get("duration").Int(),
			Cover:        value.Get("video.cover.url_list.0").String(),
			DynamicCover: value.Get("video.dynamic_cover.url_list.0").String(),
		})
		return true
	})

	sort.Slice(userPosts, func(i, j int) bool {
		return userPosts[i].CreateTime > userPosts[j].CreateTime
	})

	//Todo: 判断空响应
	return userPosts[0], nil
}
