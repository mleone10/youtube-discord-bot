package bot

import (
	"fmt"
	"html"

	"github.com/bwmarrin/discordgo"
)

type DiscordClient struct {
	client    *discordgo.Session
	channelId string
}

type Postable interface {
	Title() string
	Id() string
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

func (dc *DiscordClient) PostVideos(vs []Postable) error {
	for _, v := range vs {
		err := dc.PostVideo(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dc *DiscordClient) PostVideo(v Postable) error {
	_, err := dc.client.ChannelMessageSend(dc.channelId, messageString(v))
	if err != nil {
		return err
	}

	return nil
}

func messageString(v Postable) string {
	return fmt.Sprintf("**%s**\nhttps://www.youtube.com/watch?v=%s", html.UnescapeString(v.Title()), v.Id())
}
