package bot

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Run() {
	ytc, err := NewYouTubeClient(os.Getenv("YT_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	deltaInt, err := strconv.Atoi(os.Getenv("DELTA_MINUTES"))
	if err != nil {
		log.Fatal(err)
	}

	ytChannels := strings.Split(os.Getenv("YT_CHANNELS"), ",")

	log.Printf("Searching for videos from the past %d minutes from %d channels: %s", deltaInt, len(ytChannels), ytChannels)

	videos, err := ytc.ListRecentVideosForUsernames(ytChannels, time.Duration(deltaInt)*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Attempting to post  %d videos", len(videos))

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

	log.Println("Successfully posted videos to Discord")
}
