package models

type UniversialReq struct {
	SessionKey string `url:"sessionKey"`
	Target     int64  `url:"target"`
	MemberIds  int64  `url:"memberIds"`
}

type VerifyReq struct {
	VerifyKey string `json:"verifyKey"`
	QQ        int64  `json:"qq"`
}

type BindReq struct {
	SessionKey string `json:"sessionKey"`
	QQ         int64  `json:"qq"`
}

type ReleaseReq struct {
	SessionKey string `json:"sessionKey"`
	QQ         int64  `json:"qq"`
}
