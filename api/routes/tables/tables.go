package tables

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/glinq"
	"github.com/rlarkin212/wumpdump/discord"
)

type table struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type tablesHandler struct {
	bot *discord.Bot
}

func New(bot *discord.Bot) *tablesHandler {
	return &tablesHandler{
		bot: bot,
	}
}

func (h *tablesHandler) Tables(c *gin.Context) {
	channels, err := h.bot.Discord.GuildChannels(h.bot.Config.Discord.GuildId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	textChannels := glinq.Where(channels, func(x *discordgo.Channel) bool {
		return x.Type == discordgo.ChannelTypeGuildText
	})

	tables := glinq.Select(textChannels, func(x *discordgo.Channel) *table {
		return &table{
			Id:   x.ID,
			Name: x.Name,
		}
	})

	c.JSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}
