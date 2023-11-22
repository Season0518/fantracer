package weibo

import (
	"fmt"
	"testing"
)

func TestGetLatestBlogFunc(t *testing.T) {
	uid := int64(6593497650)

	latestBlog, err := GetLatestBlog(uid, 1)
	if err != nil || latestBlog.Ok != 1 {
		t.Error("Failed to get latest weibo")
	}

	if !(len(latestBlog.Data.Cards) >= 0) {
		t.Error("Failed to get latest weibo")
	}

	return
}

func TestExtractTextFunc(t *testing.T) {
	testBlogs := []string{
		//特殊符号&超话&位置
		"[相爱]要结束啦<br /><a  href=\"https://m.weibo.cn/p/index?extparam=%E8%8A%B1%E5%9B%A2%E5%B9%B4%E7%B3%95%E5%85%AC%E4%B8%BB&containerid=100808c6f28071bf2aef1f074a8d0fa24f0ca5&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png'></span><span class=\"surl-text\">花团年糕公主</span></a><a  href=\"https://m.weibo.cn/p/index?extparam=%E8%8A%B1%E5%9B%A2Family&containerid=100808cbaddfc50864f4459a48faea71f09109&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png'></span><span class=\"surl-text\">花团Family</span></a><a  href=\"https://m.weibo.cn/p/index?extparam=%E5%AE%88%E6%8A%A4%E8%8A%B1%E5%9B%A2%E6%9C%80%E5%A5%BD%E7%9A%84Gen&containerid=1008082f62c668c7f7316d66ee90c5549c2bf8&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png'></span><span class=\"surl-text\">守护花团最好的Gen</span></a> <a  href=\"http://weibo.com/p/1001018008621020000000000\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_location_default.png'></span><span class=\"surl-text\">大连</span></a> ",
		//话题
		"请大家一起见证我成为更好的偶像吧！<a  href=\"https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23starlink%E5%81%B6%E5%83%8F%E8%AE%A1%E5%88%92%23&extparam=%23starlink%E5%81%B6%E5%83%8F%E8%AE%A1%E5%88%92%23&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class=\"surl-text\">#starlink偶像计划#</span></a> ",
		//At&超话&位置
		"三年风雨三年晴，谁行谁用折叠屏<a href='/n/BLOSSOM-冰灵'>@BLOSSOM-冰灵</a> <a  href=\"https://m.weibo.cn/p/index?extparam=%E8%8A%B1%E5%9B%A2%E5%B9%B4%E7%B3%95%E5%85%AC%E4%B8%BB&containerid=100808c6f28071bf2aef1f074a8d0fa24f0ca5&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png'></span><span class=\"surl-text\">花团年糕公主</span></a><a  href=\"https://m.weibo.cn/p/index?extparam=%E8%8A%B1%E5%9B%A2Family&containerid=100808cbaddfc50864f4459a48faea71f09109&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png'></span><span class=\"surl-text\">花团Family</span></a> <a  href=\"http://weibo.com/p/100101B2094551D56DAAFA439F\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_location_default.png'></span><span class=\"surl-text\">长春·摩天・活力城Mall购物中心</span></a> ",
		//Emoji&超话
		"这次我是最快的啦！七夕快乐！我永远爱你们<span class=\"url-icon\"><img alt=[心] src=\"https://h5.sinaimg.cn/m/emoticon/icon/others/l_xin-43af9086c0.png\" style=\"width:1em; height:1em;\" /></span><a  href=\"https://m.weibo.cn/p/index?extparam=%E8%8A%B1%E5%9B%A2%E5%B9%B4%E7%B3%95%E5%85%AC%E4%B8%BB&containerid=100808c6f28071bf2aef1f074a8d0fa24f0ca5&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png'></span><span class=\"surl-text\">花团年糕公主</span></a><a  href=\"https://m.weibo.cn/p/index?extparam=%E8%8A%B1%E5%9B%A2Family&containerid=100808cbaddfc50864f4459a48faea71f09109&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png'></span><span class=\"surl-text\">花团Family</span></a> <a  href=\"http://weibo.com/p/100101B2094257D26CAAF8419F\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_location_default.png'></span><span class=\"surl-text\">长春·草台喜剧馆</span></a> ",
		//超话&位置(无正文)
		"<a  href=\"https://m.weibo.cn/p/index?extparam=%E8%8A%B1%E5%9B%A2%E5%B9%B4%E7%B3%95%E5%85%AC%E4%B8%BB&containerid=100808c6f28071bf2aef1f074a8d0fa24f0ca5&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png'></span><span class=\"surl-text\">花团年糕公主</span></a><a  href=\"https://m.weibo.cn/p/index?extparam=%E8%8A%B1%E5%9B%A2Family&containerid=100808cbaddfc50864f4459a48faea71f09109&luicode=10000011&lfid=1076036593497650\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png'></span><span class=\"surl-text\">花团Family</span></a> <a  href=\"http://weibo.com/p/1001018008622010000000000\" data-hide=\"\"><span class='url-icon'><img style='width: 1rem;height: 1rem' src='https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_location_default.png'></span><span class=\"surl-text\">长春</span></a>",
	}

	for _, blog := range testBlogs {
		fmt.Println(BlogTextParser(blog))
	}

	return
}
