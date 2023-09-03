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

type node struct {
	skip bool
	n    *html.Node
}

// processAttributes function to process attributes of HTML
func processAttributes(fa *node, b *strings.Builder) {
	for _, subAttr := range fa.n.FirstChild.Attr {
		if subAttr.Key == "src" && strings.Contains(subAttr.Val, "timeline_card_small_super_default.png") {
			//如果是超话则标记skip标签，并丢弃所有子节点
			fa.skip = true
			return
		} else if subAttr.Key == "src" && strings.Contains(subAttr.Val, "timeline_card_small_location_default.png") {
			//如果是位置信息则向stringBuilder添加圆图钉emoji
			b.WriteString("📍")
		}
	}
	return
}

// BlogTextParser function to extract text from HTML
// 20230903 Updated: 丢弃超话的超链接及文本，优化位置信息的显示
func BlogTextParser(s string) string {
	doc, _ := html.Parse(strings.NewReader(s))

	var b strings.Builder
	var f func(fa *node)

	//递归处理获取到的微博HTML对象
	f = func(fa *node) {
		if fa.n.Type == html.ElementNode {
			for _, attr := range fa.n.Attr {
				//根据url-icon的图标分类超链接的类型
				if attr.Key == "class" && attr.Val == "url-icon" {
					processAttributes(fa, &b)
				}
			}
		}

		if fa.n.Type == html.TextNode {
			b.WriteString(fa.n.Data)
		}
		for c := fa.n.FirstChild; c != nil; c = c.NextSibling {
			sub := node{skip: false, n: c}
			f(&sub)
			if sub.skip {
				return
			}
		}
	}
	f(&node{skip: false, n: doc})

	return b.String()
}

// GetLatestBlog function is used to get all the blogs of a user's homepage on Weibo
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
		card.Mblog.Text = BlogTextParser(card.Mblog.Text)

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
