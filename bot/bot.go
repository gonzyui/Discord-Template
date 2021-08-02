package bot

import (
	"Test/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var botID string
var client *discordgo.Session

func Start() {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	session.AddHandler(message)
	session.AddHandler(ready)

	fmt.Print("Bot is online")
	defer session.Close()
	if err = session.Open(); err != nil {
		fmt.Println(err)
		return
	}

	scall := make(chan os.Signal, 1)
	signal.Notify(scall, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-scall
}

func ready(bot *discordgo.Session, event *discordgo.Ready) {
	bot.UpdateGameStatus(0, "&ping")
}

func message(bot *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}
	switch {
	case strings.HasPrefix(message.Content, config.BotPrefix):
		ping := bot.HeartbeatLatency().Truncate(60)
		if message.Content == "&ping" {
			bot.ChannelMessageSend(message.ChannelID,`My latency is **` + ping.String() + `**!`)
		}
		if message.Content == "&author" {
			bot.ChannelMessageSend(message.ChannelID, "My author is Gonz#0001, I'm only a template discord bot made in golang.")
		}
		if message.Content == "&github" {
			embed := &discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: message.Author.AvatarURL("1024"),
				},
				Title: "My repository",
				Description: "Just click [here](https://github.com/gonzyui) to go to my repository and see if I was updated :)",
				Color: 0x00ff00,
			}
			bot.ChannelMessageSendEmbed(message.ChannelID, embed)
		}
	}
}
