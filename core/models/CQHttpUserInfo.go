package models

type UserInfo struct {
	UserID    int64  `json:"user_id"`    // QQ 号
	Nickname  string `json:"nickname"`   // 昵称
	Sex       string `json:"sex"`        // 性别, male 或 female 或 unknown
	Age       int32  `json:"age"`        // 年龄
	QID       string `json:"qid"`        // qid ID身份卡
	Level     int32  `json:"level"`      // 等级
	LoginDays int32  `json:"login_days"` // 登录天数
}
