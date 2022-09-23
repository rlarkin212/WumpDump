package drop

import (
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
	params := c.Request.URL.Query()

	channelId := params.Get("cid")
	if channelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "supply channel id",
		})
	}

	channel, err := h.bot.Discord.ChannelDelete(channelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    channel.ID,
		"table": channel.Name,
		"state": "deleted",
	})
}
