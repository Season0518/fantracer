package middleware

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

func AccessWebSocket(c *websocket.Conn) ([]byte, map[string]interface{}, error) {
	_, message, err := c.ReadMessage()
	if err != nil {
		return nil, nil, err
	}
	var data map[string]interface{}

	err = json.Unmarshal(message, &data)
	if err != nil {
		return nil, nil, err
	}

	return message, data, nil
}
