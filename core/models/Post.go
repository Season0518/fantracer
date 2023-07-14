package models

type PostRecord struct {
	UserID    int64  `xorm:"'user_id' not null pk"`
	TimeStamp int64  `xorm:"'time_stamp' not null pk"`
	Refer     string `xorm:"'refer'"`
	Text      string `xorm:"'text'"`
}

type PostInfo struct {
	UserID    int64  `xorm:"'user_id' not null pk"`
	TimeStamp int64  `xorm:"'time_stamp'"`
	Platform  string `xorm:"'platform'"`
}
