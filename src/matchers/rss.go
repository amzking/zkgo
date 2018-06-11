package matchers

import (
	"encoding/xml"
	"zkgo/src/search"
	"errors"
	"net/http"
	"fmt"
	"log"
	"regexp"
)

type (
	item struct {
		XMLName xml.Name `xml:"item"`
		PubDate string `xml:"pubDate"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		GUID string `xml:"guid"`
		GeoRssPoint string 	`xml:"georss:point"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL string `xml:"url"`
		Title string `xml:"title"`
		Link string `xml:"link"`
	}

	channel struct {
		XMLName xml.Name `xml:"channel"`
		Title string 	`xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		PubDate string `xml:"pubDate"`
		LastBuildDate string `xml:"lastBuildDate"`
		TTL string `xml:"ttl"`
		Language string `xml:"language"`
		ManaingEditor string `xml:"managingEditor"`
		WebMaster string `xml:"webMaster"`
		Image image `xml:"image"`
		Item []item `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel `xml:"channel"`
	}
)


type rssMatcher struct {}

/**
 * 每个包可以包含任意多个init函数，会在程序开始的时候被调用
 * 所有被编译器发现的init函数都会安排在main函数之前执行。
 */
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func init() {
	fmt.Println("joie")
}


func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Http response error %d\n", resp.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error){
	var results []*search.Result
	log.Printf("Search feed Type[%s] Site[%s] For Uri[%s]\n", feed.Type, feed.Name, feed.URI)

	document, err := m.retrieve(feed)

	if err != nil {
		return nil, err
	}


	for _, channelItem := range document.Channel.Item {

		matched, err := regexp.MatchString(searchTerm, channelItem.Title)

		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field: "Title",
				Content: channelItem.Title,
			})
		}

		matched1, err := regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		if matched1 {
			results = append(results, &search.Result{
				Field: "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}