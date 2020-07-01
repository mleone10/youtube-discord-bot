package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
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
	PublishedAt string `json:"publishedAt`
}

func main() {
	ytApiRootUrl, _ := url.Parse(youtubeApiRoot)
	ytSearchUrl, _ := url.Parse("/youtube/v3/search")

	req, err := http.NewRequest("GET", ytApiRootUrl.ResolveReference(ytSearchUrl).String(), nil)
	if err != nil {
		log.Fatal("Could not form YouTube search URL")
	}

	q := url.Values{}
	q.Set("key", os.Getenv("YT_API_KEY"))
	q.Set("channelId", os.Getenv("FOOD_WISHES_CHANNEL_ID"))
	q.Set("part", "snippet,id")
	q.Set("order", "date")
	q.Set("maxResults", "20")
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error performing YouTube search request")
	}

	var parsedResponse YouTubeSearchResponse
	parseResponse(resp, &parsedResponse)
	log.Printf("%+v", parsedResponse)
}

func parseResponse(r *http.Response, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
