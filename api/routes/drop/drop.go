package drop

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/wumpdump/discord"
)

type dropHandler struct {
	bot *discord.Bot
}

func New(bot *discord.Bot) *dropHandler {
	return &dropHandler{
		bot: bot,
	}
}

func (h *dropHandler) Drop(c *gin.Context) {
	channelId := c.Param("id")
	if channelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "supply channel id",
		})

		return
	}

	channel, err := h.bot.Discord.ChannelDelete(channelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"dropped": fmt.Sprintf("%s - %s", channelId, channel.Name),
	})
}
