data "aws_region" "current" {
}

variable "youtube_api_key" {}
variable "youtube_channels" {}
variable "discord_bot_token" {}
variable "discord_channel_id" {}
variable "delta_minutes" {}

resource "aws_lambda_function" "youtube_discord_bot" {
  function_name    = "youtube-discord-bot"
  filename         = "youtube-discord-bot.zip"
  handler          = "lambdabot"
  source_code_hash = filebase64sha256("youtube-discord-bot.zip")
  role             = aws_iam_role.youtube_discord_bot.arn
  runtime          = "go1.x"
  memory_size      = 512
  timeout          = 20
  environment {
    variables = {
      YT_API_KEY = var.youtube_api_key
      YT_CHANNELS = var.youtube_channels
      DISCORD_BOT_TOKEN = var.discord_bot_token
      DISCORD_CHANNEL_ID = var.discord_channel_id
      DELTA_MINUTES = var.delta_minutes
    }
  }
}

resource "aws_iam_role" "youtube_discord_bot" {
  name               = "youtube-discord-bot-role"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": {
    "Action": "sts:AssumeRole",
    "Principal": {
      "Service": "lambda.amazonaws.com"
    },
    "Effect": "Allow"
  }
}
POLICY
}
