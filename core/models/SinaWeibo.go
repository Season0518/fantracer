package models

type SinaWeiboResp struct {
	Ok   int `json:"ok" xorm:"-"`
	Data struct {
		Cards []struct {
			Timestamp int64  `json:"-"`
			Scheme    string `json:"scheme" xorm:"'scheme'"`
			Mblog     struct {
				Text      string `json:"text" xorm:"'text'"`
				CreatedAt string `json:"created_at" xorm:"-"`
			} `json:"mblog"`
		} `json:"cards"`
	} `json:"data"`
}
