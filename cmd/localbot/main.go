package main

import (
	"log"
	"os"
	"strconv"

	"github.com/mleone10/youtube-discord-bot/internal"
)

func main() {
	b := bot.NewBot()

	delta, err := strconv.Atoi(os.Getenv("DELTA_MINUTES"))
	if err != nil {
		log.Fatal(err)
	}
	b.DeltaMinutes = delta

	b.Run()
}
