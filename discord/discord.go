package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/rlarkin212/wumpdump/config"
)

type Bot struct {
	Discord *discordgo.Session
	Config  *config.Config
}

func Init(config *config.Config) *Bot {
	dg, err := discordgo.New("Bot " + config.Discord.Token)
	if err != nil {
		log.Fatal(err)
	}

	discord := &Bot{
		Discord: dg,
		Config:  config,
	}

	return discord
}

func (b *Bot) Start() {
	err := b.Discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
}

func (b *Bot) Close() {
	b.Discord.Close()
}
