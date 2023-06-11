package models

type VerifyRes struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

type BindRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ReleaseRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LatestMemberList struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []Member `json:"data"`
}
