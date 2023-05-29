package miraiResponse

import (
	"fantracer/models"
)

type Verify struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

type Bind struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}

type Release struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}

type LatestMemberList struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []models.Member `json:"data"`
}