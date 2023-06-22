package models

type MemberInfo struct {
	GroupID         int64  `xorm:"'group_id' not null pk" json:"group_id"`                // 群号
	UserID          int64  `xorm:"'user_id' not null pk" json:"user_id"`                  // QQ 号
	Nickname        string `xorm:"'nickname' not null" json:"nickname"`                   // 昵称
	Card            string `xorm:"'card' not null" json:"card"`                           // 群名片／备注
	Sex             string `xorm:"'sex' not null" json:"sex"`                             // 性别, male 或 female 或 unknown
	Age             int32  `xorm:"'age' not null" json:"age"`                             // 年龄
	Area            string `xorm:"'area' not null" json:"area"`                           // 地区
	JoinTime        int32  `xorm:"'join_time' not null" json:"join_time"`                 // 加群时间戳
	LastSentTime    int32  `xorm:"'last_sent_time' not null" json:"last_sent_time"`       // 最后发言时间戳
	Level           string `xorm:"'level' not null" json:"level"`                         // 成员等级
	Role            string `xorm:"'role' not null" json:"role"`                           // 角色, owner 或 admin 或 member
	Unfriendly      bool   `xorm:"'unfriendly' not null" json:"unfriendly"`               // 是否不良记录成员
	Title           string `xorm:"'title' not null" json:"title"`                         // 专属头衔
	TitleExpireTime int64  `xorm:"'title_expire_time' not null" json:"title_expire_time"` // 专属头衔过期时间戳
	CardChangeable  bool   `xorm:"'card_changeable' not null" json:"card_changeable"`     // 是否允许修改群名片
	ShutUpTimestamp int64  `xorm:"'shut_up_timestamp' not null" json:"shut_up_timestamp"` // 禁言到期时间
}
