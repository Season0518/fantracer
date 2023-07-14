package weibo

import (
	"core/models"
	"encoding/json"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// extractText function to extract text from HTML
func extractText(s string) string {
	doc, _ := html.Parse(strings.NewReader(s))
	var b strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			b.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return b.String()
}

func GetLatestBlog(uid int64, page int) (models.SinaWeiboResp, error) {
	requestURL := "https://m.weibo.cn/api/container/getIndex"

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return models.SinaWeiboResp{}, err
	}

	// 设置HTTP请求参数
	query := req.URL.Query()
	query.Add("containerid", "107603"+strconv.FormatInt(uid, 10))
	if page != 1 {
		query.Add("page", strconv.Itoa(page))
	}

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return models.SinaWeiboResp{}, err
	}

	//读取返回值的byte字节流
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.SinaWeiboResp{}, err
	}

	var latestBlog models.SinaWeiboResp
	err = json.Unmarshal(body, &latestBlog)
	if err != nil {
		return models.SinaWeiboResp{}, err
	}

	//Format the text
	for i, card := range latestBlog.Data.Cards {
		card.Mblog.Text = extractText(card.Mblog.Text)

		t, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", card.Mblog.CreatedAt)
		if err == nil {
			card.Timestamp = t.Unix()
		}

		latestBlog.Data.Cards[i] = card
	}

	// Sort cards by Timestamp in descending order
	sort.Slice(latestBlog.Data.Cards, func(i, j int) bool {
		return latestBlog.Data.Cards[i].Timestamp > latestBlog.Data.Cards[j].Timestamp
	})

	return latestBlog, nil
}
