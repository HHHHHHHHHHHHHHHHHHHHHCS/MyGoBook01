package matchers

import (
	"../../Search"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type (
	// item 根据 item 字段的标签，将定义的字段
	// 与 rss 文档的字段关联起来
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image 根据 image 字段的标签，将定义的字段
	// 与 rss 文档的字段关联起来
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel 根据 channel 字段的标签，将定义的字段
	// 与 rss 文档的字段关联起来
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument 定义了与 rss 文档关联的字段
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

//rssMatcher 实现了Matcher接口
type RssMatcher struct {
}

//将匹配器注册到程序里
func init() {
	var matcher RssMatcher
	Search.Register("rss", matcher)
}

//发送http get 请求 获取rss数据源并解码
func (m RssMatcher) Retrieve(feed *Search.Feed) (*rssDocument, error) {
	if feed.URL == "" {
		return nil, errors.New("No rss feed URL provided")
	}

	//从网络获得rss数据源文档
	resp, err := http.Get(feed.URL)
	if err != nil {
		return nil, err
	}

	//退出时,关闭返回的响应链接
	defer resp.Body.Close()

	//检查状态码是不是200,是不是收到了正确的响应
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

//在文档中搜索特定的选项
func (m RssMatcher) Search(feed *Search.Feed, searchTerm string) ([]*Search.Result, error) {
	var results []*Search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For Url [%s]\n", feed.Type, feed.Name, feed.URL)

	docment, err := m.Retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range docment.Channel.Item {
		//检查标题部分是否包含搜索选项
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		//如果找到匹配的项,将其作为结果保存
		if matched {
			results = append(results, &Search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		//检查描述部分是否包含搜索项
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		//如果找到匹配的项,将其作为结果保存
		if matched{
			results=append(results,&Search.Result{
				Field:"Desc",
				Content:channelItem.Description,
			})
		}
	}

	return results,nil
}
