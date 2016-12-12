package helpers

import (
	"github.com/pushapps/urlstorss/feeds"
	"io/ioutil"
	"strings"
	"net/http"
	"github.com/pushapps/urlstorss/opengraph"
	"strconv"
)

func UrlToRssItem(articleUrl string) *feeds.Item {
	rssItem := &feeds.Item{}

	response,  err := http.Get(articleUrl)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	og := opengraph.NewOpenGraph()
	err = og.ProcessHTML(strings.NewReader(string(contents)))

	if err != nil {
		return nil
	}

	rssItem.Id = og.URL
	rssItem.Title = og.Title
	rssItem.Description = og.Description
	rssItem.Link = &feeds.Link{}
	rssItem.Link.Href = og.URL
	rssItem.Media = []*feeds.Media{}
	if len(og.Videos) > 0 {
		rssItem.Media = append(rssItem.Media, &feeds.Media{
			Medium: "video",
			Url: og.Videos[0].URL,
			Height: strconv.FormatUint(og.Videos[0].Height, 16),
			Width: strconv.FormatUint(og.Videos[0].Width, 16),
		})
	}
	if len(og.Images)>0 {
		rssItem.Media = append(rssItem.Media, &feeds.Media{
			Medium: "image",
			Url: og.Images[0].URL,
			Height: strconv.FormatUint(og.Images[0].Height, 16),
			Width: strconv.FormatUint(og.Images[0].Width, 16),
		})
	}

	return rssItem
}
