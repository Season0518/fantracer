package models

type GroupIncreaseEvent struct {
	EventUniversal
	NoticeUniversal

	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
}

type GroupDecreaseEvent struct {
	EventUniversal
	NoticeUniversal
	
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
}
