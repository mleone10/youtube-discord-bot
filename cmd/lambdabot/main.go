package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	bot "github.com/mleone10/youtube-discord-bot/internal"
)

func main() {
	lambda.Start(bot.Run)
}
