package models

type MessageUniversal struct {
	MessageType string             `json:"message_type"`
	SubType     string             `json:"sub_type"`
	MessageID   int32              `json:"message_id"`
	UserID      int64              `json:"user_id"`
	Message     string             `json:"message"`
	RawMessage  string             `json:"raw_message"`
	Font        int                `json:"font"`
	Sender      GroupMessageSender `json:"sender"`
}

type MessageBody struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

type MessageSender struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

type GroupMessageSender struct {
	MessageSender

	Card  string `json:"card"`
	Area  string `json:"area"`
	Level string `json:"level"`
	Role  string `json:"role"`
	Title string `json:"title"`
}

type Anonymous struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}
