package mirai

type Universial struct {
	SessionKey string `url:"sessionKey"`
	Target     string `url:"target"`
	MemberIds  string `url:"memberIds"`
}

type Verify struct {
	VerifyKey string `json:"verifyKey"`
	Qq        int    `json:"qq"`
}

type Bind struct {
	SessionKey string `json:"sessionKey"`
	QQ         int    `json:"qq"`
}