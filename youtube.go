package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	youtubeApiRoot string = "https://www.googleapis.com"
)

type YouTubeSearchResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id      ItemId  `json:"id"`
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

func (c *YouTubeClient) ListRecentVideosFromChannels(channelIds []string, maxResults int) error {
	ytSearchUrl, _ := url.Parse("/youtube/v3/search")
	req, err := http.NewRequest("GET", c.apiRootUrl.ResolveReference(ytSearchUrl).String(), nil)
	if err != nil {
		return err
	}

	q := url.Values{}
	q.Set("key", c.apiKey)
	q.Set("part", "snippet,id")
	q.Set("order", "date")
	q.Set("maxResults", strconv.Itoa(maxResults))
	for _, cid := range channelIds {
		q.Add("channelId", cid)
	}
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	var parsedResponse YouTubeSearchResponse
	parseResponse(resp, &parsedResponse)

	log.Printf("%+v", parsedResponse)

	// TODO: This should actually return something...
	return nil
}
