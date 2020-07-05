package main

import (
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mleone10/youtube-discord-bot/internal"
)

func initAndRun() {
	b := bot.NewBot()

	delta, err := strconv.Atoi(os.Getenv("DELTA_MINUTES"))
	if err != nil {
		log.Fatal(err)
	}
	b.DeltaMinutes = delta

	b.Run()
}

func main() {
	lambda.Start(initAndRun)
}
