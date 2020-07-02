package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	ytc, err := NewYouTubeClient(os.Getenv("YT_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	deltaInt, err := strconv.Atoi(os.Getenv("YT_DELTA_MINUTES"))
	if err != nil {
		log.Fatal(err)
	}

	videos, err := ytc.ListRecentVideosForUsernames(strings.Split(os.Getenv("YT_CHANNELS"), ","), time.Duration(deltaInt)*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d videos to post", len(videos))

	dc, err := NewDiscordClient(os.Getenv("DISCORD_BOT_TOKEN"), os.Getenv("DISCORD_CHANNEL_ID"))
	if err != nil {
		log.Fatal(err)
	}

	var ls []Postable
	for _, v := range videos {
		ls = append(ls, v)
	}

	err = dc.PostVideos(ls)
	if err != nil {
		log.Fatal(err)
	}
}
