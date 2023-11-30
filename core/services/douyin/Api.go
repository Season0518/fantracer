// Description: 该部分提供的抖音接口是在本地部署的。如需推送请自行搭建接口服务器
package douyin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type agent struct {
	StatusCode int    `json:"status_code"`
	Data       string `json:"data"`
}

type Profile struct {
	StatusCode int `json:"status_code,omitempty"`
	StatusMsg  any `json:"status_msg,omitempty"`
	User       struct {
		AwemeCount                        int    `json:"aweme_count,omitempty"`
		AwemeCountCorrectionThreshold     int    `json:"aweme_count_correction_threshold,omitempty"`
		City                              string `json:"city,omitempty"`
		Country                           string `json:"country,omitempty"`
		District                          string `json:"district,omitempty"`
		DongtaiCount                      int    `json:"dongtai_count,omitempty"`
		FavoritingCount                   int    `json:"favoriting_count,omitempty"`
		FollowerCount                     int    `json:"follower_count,omitempty"`
		FollowingCount                    int    `json:"following_count,omitempty"`
		ForwardCount                      int    `json:"forward_count,omitempty"`
		Gender                            int    `json:"gender,omitempty"`
		IsActivityUser                    bool   `json:"is_activity_user,omitempty"`
		IsBan                             bool   `json:"is_ban,omitempty"`
		LiveCommerce                      bool   `json:"live_commerce,omitempty"`
		LiveStatus                        int    `json:"live_status,omitempty"`
		MaxFollowerCount                  int    `json:"max_follower_count,omitempty"`
		MplatformFollowersCount           int    `json:"mplatform_followers_count,omitempty"`
		Nickname                          string `json:"nickname,omitempty"`
		Province                          string `json:"province,omitempty"`
		RoleID                            string `json:"role_id,omitempty"`
		RoomID                            int    `json:"room_id,omitempty"`
		SecUID                            string `json:"sec_uid,omitempty"`
		Signature                         string `json:"signature,omitempty"`
		TotalFavorited                    int    `json:"total_favorited,omitempty"`
		TotalFavoritedCorrectionThreshold int    `json:"total_favorited_correction_threshold,omitempty"`
		UID                               string `json:"uid,omitempty"`
		UniqueID                          string `json:"unique_id,omitempty"`
		UserAge                           int    `json:"user_age,omitempty"`
	} `json:"user,omitempty"`
}

type UserInfo struct {
	Profile
	Posts struct {
		StatusCode int `json:"status_code,omitempty"`
		AwemeList  []struct {
			AwemeID    string `json:"aweme_id,omitempty"`
			Desc       string `json:"desc,omitempty"`
			CreateTime int    `json:"create_time,omitempty"`
			ItemTitle  string `json:"item_title,omitempty"`
			Video      struct {
				Cover struct {
					URI     string   `json:"uri,omitempty"`
					URLList []string `json:"url_list,omitempty"`
					Width   int      `json:"width,omitempty"`
					Height  int      `json:"height,omitempty"`
				} `json:"cover,omitempty"`
				Height       int `json:"height,omitempty"`
				Width        int `json:"width,omitempty"`
				DynamicCover struct {
					URI     string   `json:"uri,omitempty"`
					URLList []string `json:"url_list,omitempty"`
					Width   int      `json:"width,omitempty"`
					Height  int      `json:"height,omitempty"`
				} `json:"dynamic_cover,omitempty"`
				OriginCover struct {
					URI     string   `json:"uri,omitempty"`
					URLList []string `json:"url_list,omitempty"`
					Width   int      `json:"width,omitempty"`
					Height  int      `json:"height,omitempty"`
				} `json:"origin_cover,omitempty"`
			} `json:"video,omitempty"`
			ShareURL  string `json:"share_url,omitempty"`
			ShareInfo struct {
				ShareURL      string `json:"share_url,omitempty"`
				ShareLinkDesc string `json:"share_link_desc,omitempty"`
			} `json:"share_info,omitempty"`
			Original         int  `json:"original,omitempty"`
			Duration         int  `json:"duration,omitempty"`
			AwemeType        int  `json:"aweme_type,omitempty"`
			MediaType        int  `json:"media_type,omitempty"`
			ReportAction     bool `json:"report_action,omitempty"`
			VisualSearchInfo struct {
				IsShowEntrance        bool   `json:"is_show_entrance,omitempty"`
				Extra                 string `json:"extra,omitempty"`
				VisualSearchLongpress int    `json:"visual_search_longpress,omitempty"`
			} `json:"visual_search_info,omitempty"`
			PreviewTitle string `json:"preview_title,omitempty"`
		} `json:"aweme_list,omitempty"`
	}
}

func (a *agent) GetFromAgent(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(body))

	err = json.Unmarshal(body, a)
	if err != nil {
		return err
	}

	if a.StatusCode != 0 {
		return fmt.Errorf("抖音代理服务器响应异常: %d", a.StatusCode)
	}

	return nil
}

func (u *UserInfo) GetPosts(secUserId string) error {
	var resp agent
	requestURL := fmt.Sprintf("http://localhost:15000/get_post?sec_user_id=%s", secUserId)

	err := resp.GetFromAgent(requestURL)
	if err != nil {
		return err
	}

	_ = json.Unmarshal([]byte(resp.Data), &u.Posts)

	return nil
}

func (u *UserInfo) GetProfile(secUserId string) error {
	var resp agent
	requestURL := fmt.Sprintf("http://localhost:15000/get_profile?sec_user_id=%s", secUserId)

	err := resp.GetFromAgent(requestURL)
	if err != nil {
		return err
	}
	_ = json.Unmarshal([]byte(resp.Data), &u.Posts)

	return nil
}
