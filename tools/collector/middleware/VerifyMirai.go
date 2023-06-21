package middleware

import (
	"bytes"
	"core/models"
	"core/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func VerifySession(verifyKey string, qq int64) (string, error) {
	router := "/verify"

	baseUrl, err := utils.ReadMiraiBaseUrl()
	if err != nil {
		log.Panicln("baseUrl非法, 无法链接mirai-http服务")
	}

	requestBody := models.VerifyReq{
		VerifyKey: verifyKey,
		QQ:        qq,
	}

	// 发送 POST 请求
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("序列化失败: %v", err)
	}

	// 读取响应体
	response, err := http.Post(baseUrl+router, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("获取SessionKey失败: %v", err)
	}
	defer response.Body.Close()

	// 解析响应体
	var session models.VerifyRes
	err = json.NewDecoder(response.Body).Decode(&session)
	if err != nil {
		return "", fmt.Errorf("解析/verify的返回值失败: %v", err)
	}

	// 处理响应
	if session.Code != 0 {
		return "", fmt.Errorf("/verify返回值异常: %d", session.Code)
	}

	return session.Session, nil
}

func BindSession(sessionKey string, qq int64) (string, error) {
	router := "/bind"

	baseUrl, err := utils.ReadMiraiBaseUrl()
	if err != nil {
		log.Panicln("baseUrl非法, 无法链接mirai-http服务")
	}

	requestBody := models.BindReq{
		SessionKey: sessionKey,
		QQ:         qq,
	}

	// 发送 POST 请求
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("序列化失败: %v", err)
	}

	// 读取响应体
	response, err := http.Post(baseUrl+router, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("获取SessionKey失败: %v", err)
	}
	defer response.Body.Close()

	// 解析响应体
	var session models.BindRes
	err = json.NewDecoder(response.Body).Decode(&session)
	if err != nil {
		return "", fmt.Errorf("解析/bind的返回值失败: %v", err)
	}

	// 处理响应
	if session.Code != 0 {
		return session.Msg, fmt.Errorf("/bind返回值异常: %d", session.Code)
	}

	return session.Msg, nil
}

func ReleaseSession(sessionKey string, qq int64) (string, error) {
	router := "/release"

	baseUrl, err := utils.ReadMiraiBaseUrl()
	if err != nil {
		log.Panicln("baseUrl非法, 无法链接mirai-http服务")
	}

	requestBody := models.ReleaseReq{
		SessionKey: sessionKey,
		QQ:         qq,
	}

	// 发送 POST 请求
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("序列化失败: %v", err)
	}

	// 读取响应体
	response, err := http.Post(baseUrl+router, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("获取SessionKey失败: %v", err)
	}
	defer response.Body.Close()

	// 解析响应体
	var session models.ReleaseRes
	err = json.NewDecoder(response.Body).Decode(&session)
	if err != nil {
		return "", fmt.Errorf("解析/release的返回值失败: %v", err)
	}

	// 处理响应
	if session.Code != 0 {
		return session.Msg, fmt.Errorf("/release返回值异常: %d", session.Code)
	}

	return session.Msg, nil
}
