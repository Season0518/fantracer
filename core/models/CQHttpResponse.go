package models

type CQUniversalResp struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
	Msg     string `json:"msg,omitempty"`
	Wording string `json:"wording,omitempty"`
	Echo    string `json:"echo,omitempty"`
}
