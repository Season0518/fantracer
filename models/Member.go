package models

type Member struct {
    ID                 int64    `json:"id" xorm:"unique(unique_member_group)"`
    GroupID            int64    `json:"-" xorm:"unique(unique_member_group)"`
    MemberName         string `json:"memberName"`
    SpecialTitle       string `json:"specialTitle"`
    Permission         string `json:"permission"`
    JoinTimestamp      int64    `json:"joinTimestamp"`
    LastSpeakTimestamp int64    `json:"lastSpeakTimestamp"`
    MuteTimeRemaining  int64    `json:"muteTimeRemaining"`
    Group              Group  `json:"group" xorm:"-"`
}