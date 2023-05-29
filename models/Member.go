package models

type Member struct {
    ID                 int    `json:"id"`
    MemberName         string `json:"memberName"`
    SpecialTitle       string `json:"specialTitle"`
    Permission         string `json:"permission"`
    JoinTimestamp      int    `json:"joinTimestamp"`
    LastSpeakTimestamp int    `json:"lastSpeakTimestamp"`
    MuteTimeRemaining  int    `json:"muteTimeRemaining"`
    Group              Group  `json:"group"`
}