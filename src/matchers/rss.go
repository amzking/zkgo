package matchers

import "encoding/xml"

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