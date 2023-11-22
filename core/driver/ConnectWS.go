package driver

import (
	"github.com/gorilla/websocket"
)

var Conn *websocket.Conn

func InitWS() error {
	var err error

	Conn, _, err = websocket.DefaultDialer.Dial(Base.Bot.WebSocket+"/", nil)
	if err != nil {
		return err
	}

	return nil
}
