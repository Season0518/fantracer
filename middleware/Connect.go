package middleware

import (
	"bytes"
	"encoding/json"
	miraiRequest "fantracer/models/mirai/Request"
	miraiResponse "fantracer/models/mirai/Response"
	"fantracer/utils"
	"fmt"
	"log"
	"net/http"
)

func VerifySession(verifyKey string,qq int64) (string, error) {
	router := "/verify"
	
	baseUrl,err := utils.ReadBaseUrl() 
	if err != nil {
		log.Panicln("baseUrl非法, 无法链接mirai-http服务")
	}
	
	requestBody := miraiRequest.Verify {
		VerifyKey: verifyKey,
		QQ: qq,
	}

	// 发送 POST 请求
	jsonData,err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("序列化失败: %v",err)
	}

	// 读取响应体
	response,err := http.Post(baseUrl+router,"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("获取SessionKey失败: %v",err)
	}
	defer response.Body.Close()
	
	// 解析响应体
	var session miraiResponse.Verify
	err = json.NewDecoder(response.Body).Decode(&session)
	if err != nil {
		return "", fmt.Errorf("解析/verify的返回值失败: %v",err)
	}

	// 处理响应
	if session.Code != 0 {
		return "", fmt.Errorf("/verify返回值异常: %d", session.Code)
	}
	
	return session.Session, nil
}

func BindSession(sessionKey string,qq int64) (string, error) {
	router := "/bind"

	baseUrl,err := utils.ReadBaseUrl() 
	if err != nil {
		log.Panicln("baseUrl非法, 无法链接mirai-http服务")
	}

	requestBody := miraiRequest.Bind {
		SessionKey: sessionKey,
		QQ: qq,
	}

	// 发送 POST 请求
	jsonData,err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("序列化失败: %v",err)
	}

	// 读取响应体
	response,err := http.Post(baseUrl+router,"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("获取SessionKey失败: %v",err)
	}
	defer response.Body.Close()

	// 解析响应体
	var session miraiResponse.Bind
	err = json.NewDecoder(response.Body).Decode(&session)
	if err != nil {
		return "", fmt.Errorf("解析/bind的返回值失败: %v",err)
	}

	// 处理响应
	if session.Code != 0 {
		return session.Msg, fmt.Errorf("/bind返回值异常: %d", session.Code)
	}
	
	return session.Msg, nil
}

func ReleaseSession(sessionKey string,qq int64) (string, error) {
	router := "/release"

	baseUrl,err := utils.ReadBaseUrl() 
	if err != nil {
		log.Panicln("baseUrl非法, 无法链接mirai-http服务")
	}

	requestBody := miraiRequest.Release {
		SessionKey: sessionKey,
		QQ: qq,
	}

	// 发送 POST 请求
	jsonData,err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("序列化失败: %v",err)
	}

	// 读取响应体
	response,err := http.Post(baseUrl+router,"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("获取SessionKey失败: %v",err)
	}
	defer response.Body.Close()

	// 解析响应体
	var session miraiResponse.Release
	err = json.NewDecoder(response.Body).Decode(&session)
	if err != nil {
		return "", fmt.Errorf("解析/release的返回值失败: %v",err)
	}

	// 处理响应
	if session.Code != 0 {
		return session.Msg, fmt.Errorf("/release返回值异常: %d", session.Code)
	}
	
	return session.Msg, nil	
}


