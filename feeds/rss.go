package feeds

// rss support
// validation done according to spec here:
//    http://cyber.law.harvard.edu/rss/rss.html

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
	"strings"
)

// private wrapper around the RssFeed which gives us the <rss>..</rss> xml
type rssFeedXml struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"xmlns:media,attr"`
	Channel *RssFeed
}

type RssImage struct {
	XMLName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Width   string   `xml:"width,omitempty"`
	Height  string   `xml:"height,omitempty"`
}

type RssTextInput struct {
	XMLName     xml.Name `xml:"textInput"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Name        string   `xml:"name"`
	Link        string   `xml:"link"`
}

type RssFeed struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`       // required
	Link           string   `xml:"link"`        // required
	Description    string   `xml:"description"` // required
	Language       string   `xml:"language,omitempty"`
	Copyright      string   `xml:"copyright,omitempty"`
	ManagingEditor string   `xml:"managingEditor,omitempty"` // Author used
	WebMaster      string   `xml:"webMaster,omitempty"`
	PubDate        string   `xml:"pubDate,omitempty"`       // created or updated
	LastBuildDate  string   `xml:"lastBuildDate,omitempty"` // updated used
	Category       string   `xml:"category,omitempty"`
	Generator      string   `xml:"generator,omitempty"`
	Docs           string   `xml:"docs,omitempty"`
	Cloud          string   `xml:"cloud,omitempty"`
	Ttl            int      `xml:"ttl,omitempty"`
	Rating         string   `xml:"rating,omitempty"`
	SkipHours      string   `xml:"skipHours,omitempty"`
	SkipDays       string   `xml:"skipDays,omitempty"`
	Image          *RssImage
	TextInput      *RssTextInput
	Items          []*RssItem
}

type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`       // required
	Link        string   `xml:"link"`        // required
	Description string   `xml:"description"` // required
	Author      string   `xml:"author,omitempty"`
	Category    string   `xml:"category,omitempty"`
	Comments    string   `xml:"comments,omitempty"`
	Enclosure   *RssEnclosure
	Guid        string `xml:"guid,omitempty"`    // Id used
	PubDate     string `xml:"pubDate,omitempty"` // created or updated
	Source      string `xml:"source,omitempty"`
	MediaGroup  *RssMediaGroup `xml:"media:group,omitempty"`
	MediaContent *RssMedia `xml:"media:content,omitempty"`
}

type RssMediaGroup struct {
	XMLName      xml.Name `xml:"media:group"`
	MediaContent []*RssMedia
}

type RssMedia struct {
	XMLName xml.Name `xml:"media:content"`
	Medium string `xml:"medium,attr,omitempty"`
	Url    string `xml:"url,attr,omitempty"`
	Height string `xml:"height,attr,omitempty"`
	Width  string `xml:"width,attr,omitempty"`
}

type RssEnclosure struct {
	//RSS 2.0 <enclosure url="http://example.com/file.mp3" length="123456789" type="audio/mpeg" />
	XMLName xml.Name `xml:"enclosure"`
	Url     string   `xml:"url,attr"`
	Length  string   `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

type Rss struct {
	*Feed
}

// create a new RssItem with a generic Item struct's data
func newRssItem(i *Item) *RssItem {
	item := &RssItem{
		Title:       fmt.Sprintf("<![CDATA[%s]]>", i.Title),
		Link:        i.Link.Href,
		Description: fmt.Sprintf("<![CDATA[%s]]>", i.Description),
		Guid:        i.Id,
		PubDate:     anyTimeFormat(time.RFC1123Z, i.Created, i.Updated),
	}

	if len(i.Media) > 1 {
		item.MediaGroup = &RssMediaGroup{}
		item.MediaGroup.MediaContent = []*RssMedia{}
		for counter := 0; counter < len(i.Media); counter++ {
			rssMediaContentItem := &RssMedia{
				Url: i.Media[counter].Url,
				Medium: i.Media[counter].Medium,
			}

			if len(i.Media[counter].Height) > 0 && strings.Compare(i.Media[counter].Height, "0") != 0 {
				rssMediaContentItem.Height = i.Media[counter].Height
			}

			if len(i.Media[counter].Width) > 0 && strings.Compare(i.Media[counter].Width, "0") != 0 {
				rssMediaContentItem.Width = i.Media[counter].Width
			}

			item.MediaGroup.MediaContent = append(item.MediaGroup.MediaContent, rssMediaContentItem)
		}
	} else if len(i.Media) == 1 {
		// treat 1 as a special case as it does not need the enclosing media:group tag
		rssMediaContentItem := &RssMedia{
			Url: i.Media[0].Url,
			Medium: i.Media[0].Medium,
		}

		if len(i.Media[0].Height) > 0 && strings.Compare(i.Media[0].Height, "0") != 0 {
			rssMediaContentItem.Height = i.Media[0].Height
		}

		if len(i.Media[0].Width) > 0 && strings.Compare(i.Media[0].Width, "0") != 0 {
			rssMediaContentItem.Width = i.Media[0].Width
		}

		item.MediaContent = rssMediaContentItem
	}

	intLength, err := strconv.ParseInt(i.Link.Length, 10, 64)

	if err == nil && (intLength > 0 || i.Link.Type != "") {
		item.Enclosure = &RssEnclosure{Url: i.Link.Href, Type: i.Link.Type, Length: i.Link.Length}
	}
	if i.Author != nil {
		item.Author = i.Author.Name
	}
	return item
}

// create a new RssFeed with a generic Feed struct's data
func (r *Rss) RssFeed() *RssFeed {
	pub := anyTimeFormat(time.RFC1123Z, r.Created, r.Updated)
	build := anyTimeFormat(time.RFC1123Z, r.Updated)
	author := ""
	if r.Author != nil {
		author = r.Author.Email
		if len(r.Author.Name) > 0 {
			author = fmt.Sprintf("%s (%s)", r.Author.Email, r.Author.Name)
		}
	}

	channel := &RssFeed{
		Title:          r.Title,
		Link:           r.Link.Href,
		Description:    r.Description,
		ManagingEditor: author,
		PubDate:        pub,
		LastBuildDate:  build,
		Copyright:      r.Copyright,
	}
	for _, i := range r.Items {
		channel.Items = append(channel.Items, newRssItem(i))
	}
	return channel
}

// return an XML-Ready object for an Rss object
func (r *Rss) FeedXml() interface{} {
	// only generate version 2.0 feeds for now
	return r.RssFeed().FeedXml()

}

// return an XML-ready object for an RssFeed object
func (r *RssFeed) FeedXml() interface{} {
	return &rssFeedXml{Version: "2.0", Channel: r, Media: "http://search.yahoo.com/mrss/"}
}
