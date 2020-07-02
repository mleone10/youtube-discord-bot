package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	youtubeApiRoot string = "https://www.googleapis.com"
)

type YouTubeSearchResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	ItemId  ItemId  `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type ItemId struct {
	VideoId string `json:"videoId"`
}

type Snippet struct {
	Title       string `json:"title"`
	PublishedAt string `json:"publishedAt"`
}

type YouTubeClient struct {
	client     *http.Client
	apiRootUrl *url.URL
	apiKey     string
}

type YouTubeChannelResponse struct {
	Items []ChannelItem `json:"items"`
}

type ChannelItem struct {
	Id string `json:"id"`
}

func NewYouTubeClient(apiKey string) (*YouTubeClient, error) {
	url, err := url.Parse(youtubeApiRoot)
	if err != nil {
		return nil, err
	}

	return &YouTubeClient{
		client:     &http.Client{},
		apiRootUrl: url,
		apiKey:     apiKey,
	}, nil
}

func (c *YouTubeClient) ListRecentVideosForUsernames(usernames []string, delta time.Duration) ([]Item, error) {
	var channelIds []string
	for _, u := range usernames {
		cid, err := c.getChannelId(u)
		if err != nil {
			return nil, err
		}
		channelIds = append(channelIds, cid)
	}

	var items []Item
	for _, cid := range channelIds {
		newItems, err := c.getRecentVideos(cid, delta)
		if err != nil {
			return nil, err
		}

		items = append(items, newItems...)
	}

	return items, nil
}

func (c *YouTubeClient) getChannelId(username string) (string, error) {
	ytChannelSearchUrl, _ := url.Parse("/youtube/v3/channels")
	req, err := http.NewRequest("GET", c.apiRootUrl.ResolveReference(ytChannelSearchUrl).String(), nil)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	q.Set("key", c.apiKey)
	q.Set("part", "id")
	q.Set("forUsername", username)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	var parsedResponse YouTubeChannelResponse
	parseResponse(resp, &parsedResponse)

	if l := len(parsedResponse.Items); l != 1 {
		return "", fmt.Errorf("Expected 1 channel ID for username %s, found %d", username, l)
	}

	return parsedResponse.Items[0].Id, nil
}

func (c *YouTubeClient) getRecentVideos(channelId string, delta time.Duration) ([]Item, error) {
	ytSearchUrl, _ := url.Parse("/youtube/v3/search")
	req, err := http.NewRequest("GET", c.apiRootUrl.ResolveReference(ytSearchUrl).String(), nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Set("key", c.apiKey)
	q.Set("part", "snippet,id")
	q.Set("order", "date")
	q.Set("maxResults", strconv.Itoa(5))
	q.Set("publishedAfter", time.Now().Add(-1*delta).Format(time.RFC3339))
	q.Set("channelId", channelId)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var parsedResponse YouTubeSearchResponse
	parseResponse(resp, &parsedResponse)

	return parsedResponse.Items, nil
}

func (i Item) Id() string {
	return i.ItemId.VideoId
}

func (i Item) Title() string {
	return i.Snippet.Title
}

func parseResponse(r *http.Response, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
