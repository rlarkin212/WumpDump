package alter

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/wumpdump/discord"
	"github.com/rlarkin212/wumpdump/entity/dto"
)

type alterHandler struct {
	bot *discord.Bot
}

func New(bot *discord.Bot) *alterHandler {
	return &alterHandler{
		bot: bot,
	}
}

func (h *alterHandler) Alter(c *gin.Context) {
	channelId := c.Param("id")
	if channelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "supply channel id",
		})

		return
	}

	var input dto.Table
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	channel, err := h.bot.Discord.ChannelEdit(channelId, &discordgo.ChannelEdit{Name: input.Name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    channel.ID,
		"table": channel.Name,
	})
}
