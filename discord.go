package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordClient struct {
	client    *discordgo.Session
	channelId string
}

type Video struct {
	title string
	id    string
}

func NewDiscordClient(token, channelId string) (*DiscordClient, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}

	err = session.Open()
	if err != nil {
		return nil, err
	}

	return &DiscordClient{
		client:    session,
		channelId: channelId,
	}, nil
}

func (dc *DiscordClient) PostVideo(v *Video) error {
	_, err := dc.client.ChannelMessageSend(dc.channelId, v.MessageString())
	if err != nil {
		return err
	}

	return nil
}

func (v *Video) MessageString() string {
	return fmt.Sprintf("**%s**\nhttps://www.youtube.com/watch?v=%s", v.title, v.id)
}
