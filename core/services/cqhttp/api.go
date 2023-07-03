package cqhttp

import (
	"core/utils"
	"io"
	"log"
	"net/http"
)

func FetchHttpData(accessToken string, route string, params map[string]string) ([]byte, error) {
	baseUrl, err := utils.ReadCQBaseUrl()
	if err != nil {
		log.Panicln("baseUrl非法, 无法链接go-cqHttp服务")
	}

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
