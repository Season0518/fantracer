package models

type GroupInfo struct {
	ID              int64  `xorm:"'id' not null pk autoincr" json:"-"`                    // 自增ID
	GroupID         int64  `xorm:"'group_id' not null" json:"group_id"`                   // 群号
	GroupName       string `xorm:"'group_name' not null" json:"group_name"`               // 群名称
	GroupMemo       string `xorm:"'group_memo' not null" json:"group_memo"`               // 群备注
	GroupCreateTime uint32 `xorm:"'group_create_time' not null" json:"group_create_time"` // 群创建时间
	GroupLevel      uint32 `xorm:"'group_level' not null" json:"group_level"`             // 群等级
	MemberCount     int32  `xorm:"'member_count' not null" json:"member_count"`           // 成员数
	MaxMemberCount  int32  `xorm:"'max_member_count' not null" json:"max_member_count"`   // 最大成员数（群容量）
	InfoRetrievedAt int64  `xorm:"'info_retrieved_at' not null" json:"-"`                 // 获取信息的时间
}

type GroupBlackList struct {
	ID         int64  `xorm:"'id' not null pk autoincr" json:"-"` // 自增ID
	GroupID    int64  `xorm:"'group_id' not null" json:"-"`       // 群号
	UserID     int64  `xorm:"'user_id' not null" json:"-"`        // 用户ID
	Time       int64  `xorm:"'time'" json:"-"`                    // 操作时间
	OperatorID int64  `xorm:"'operator_id' not null" json:"-"`    // 执行人ID
	SubType    string `xorm:"'sub_type' not null" json:"-"`       // 事件子类型
}
