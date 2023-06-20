package models

type GroupMessageEvent struct {
	EventUniversal
	MessageUniversal

	GroupID   int64     `json:"group_id"`
	Anonymous Anonymous `json:"anonymous"`
}
