package create

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/wumpdump/discord"
)

type table struct {
	Name string `json:"name"`
}

type createHandler struct {
	bot *discord.Bot
}

func New(bot *discord.Bot) *createHandler {
	return &createHandler{
		bot: bot,
	}
}

func (h *createHandler) Create(c *gin.Context) {
	var input table

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	channel, err := h.bot.Discord.GuildChannelCreate(h.bot.Config.Discord.GuildId, input.Name, discordgo.ChannelTypeGuildText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    channel.ID,
		"table": channel.Name,
	})
}
