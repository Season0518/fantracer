package miraiRequest

type Universial struct {
	SessionKey string `url:"sessionKey"`
	Target     int64 `url:"target"`
	MemberIds  int64 `url:"memberIds"`
}

type Verify struct {
	VerifyKey string `json:"verifyKey"`
	QQ        int64    `json:"qq"`
}

type Bind struct {
	SessionKey string `json:"sessionKey"`
	QQ         int64    `json:"qq"`
}

type Release struct {
	SessionKey string `json:"sessionKey"`
	QQ         int64    `json:"qq"`	
}
