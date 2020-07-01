package main

import (
	"log"
	"os"
)

func main() {
	ytClient, err := NewYouTubeClient(os.Getenv("YT_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	ytClient.ListRecentVideosFromChannels([]string{os.Getenv("FOOD_WISHES_CHANNEL_ID")}, 5)
}
