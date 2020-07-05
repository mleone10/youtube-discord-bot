package bot

import (
	"log"
	"os"
	"strings"
	"time"
)

type Bot struct {
	YouTubeApiKey    string
	DiscordApiKey    string
	DiscordChannelId string
	YouTubeChannels  []string
	DeltaMinutes     int
}

func NewBot() *Bot {
	return &Bot{
		YouTubeApiKey:    os.Getenv("YT_API_KEY"),
		DiscordApiKey:    os.Getenv("DISCORD_BOT_TOKEN"),
		DiscordChannelId: os.Getenv("DISCORD_CHANNEL_ID"),
		YouTubeChannels:  strings.Split(os.Getenv("YT_CHANNELS"), ","),
		DeltaMinutes:     30,
	}
}

func (b *Bot) Run() {
	ytc, err := NewYouTubeClient(b.YouTubeApiKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Searching for videos from the past %d minutes from %d channels: %s", b.DeltaMinutes, len(b.YouTubeChannels), b.YouTubeChannels)

	videos, err := ytc.ListRecentVideosForUsernames(b.YouTubeChannels, time.Duration(b.DeltaMinutes)*time.Minute)
	if err != nil {
		log.Fatal(err)
	}
	if len(videos) == 0 {
		log.Print("No new videos found")
		return
	}

	log.Printf("Attempting to post %d videos", len(videos))

	dc, err := NewDiscordClient(b.DiscordApiKey, b.DiscordChannelId)
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
