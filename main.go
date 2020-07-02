package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	ytClient, err := NewYouTubeClient(os.Getenv("YT_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	deltaInt, err := strconv.Atoi(os.Getenv("YT_DELTA_MINUTES"))
	if err != nil {
		log.Fatal(err)
	}

	ytClient.ListRecentVideosFromChannels([]string{os.Getenv("FOOD_WISHES_CHANNEL_ID")}, 5, time.Duration(deltaInt)*time.Minute)

	dClient, err := NewDiscordClient(os.Getenv("DISCORD_BOT_TOKEN"), os.Getenv("DISCORD_CHANNEL_ID"))
	if err != nil {
		log.Fatal(err)
	}

	err = dClient.PostVideo(&Video{title: "Test Title", id: "I7OcL8j6rhk"})
	if err != nil {
		log.Fatal(err)
	}
}
