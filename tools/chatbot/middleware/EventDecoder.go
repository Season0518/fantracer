package middleware

import (
	"encoding/json"
	"fmt"
)

// building
func EventDecoder(data []byte) ([]byte, map[string]interface{}, error) {
	var event map[string]interface{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, nil, err
	}

	switch event["post_type"] {
	case "message":
		fmt.Print("this is a message. trying... : ")
		//models.GroupInfo{GroupID: }
		//json.Unmarshal(data,&)

	case "notice":
		fmt.Print("this is a notice. trying... : ")
		//json.Unmarshal()

	default:
		return nil, nil, nil
	}

	return nil, nil, nil
}
