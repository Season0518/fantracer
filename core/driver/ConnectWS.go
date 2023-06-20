package driver

import (
	"core/utils"
	"net/url"

	"github.com/gorilla/websocket"
)

var Conn *websocket.Conn

func InitWS() error {
	var err error

	host, err := utils.ReadWebsocketHost()
	if err != nil {
		return err
	}

	u := url.URL{Scheme: "ws", Host: host, Path: "/"}
	Conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	return nil
}
