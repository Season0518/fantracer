package models

type SendGroupMsg struct {
	GroupID    int64       `json:"group_id"`
	Message    MessageBody `json:"message"`
	AutoEscape bool        `json:"auto_escape"`
}
