package models

type Member struct {
    ID                 int    `json:"id" xorm:"unique(unique_member_group)"`
    GroupID            int    `json:"-" xorm:"unique(unique_member_group)"`
    MemberName         string `json:"memberName"`
    SpecialTitle       string `json:"specialTitle"`
    Permission         string `json:"permission"`
    JoinTimestamp      int    `json:"joinTimestamp"`
    LastSpeakTimestamp int    `json:"lastSpeakTimestamp"`
    MuteTimeRemaining  int    `json:"muteTimeRemaining"`
    Group              Group  `json:"group" xorm:"-"`
}