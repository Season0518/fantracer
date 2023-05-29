package miraiRequest

type Universial struct {
	SessionKey string `url:"sessionKey"`
	Target     int `url:"target"`
	MemberIds  int `url:"memberIds"`
}

type Verify struct {
	VerifyKey string `json:"verifyKey"`
	QQ        int    `json:"qq"`
}

type Bind struct {
	SessionKey string `json:"sessionKey"`
	QQ         int    `json:"qq"`
}

type Release struct {
	SessionKey string `json:"sessionKey"`
	QQ         int    `json:"qq"`	
}
