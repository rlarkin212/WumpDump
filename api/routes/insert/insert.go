package insert

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/wumpdump/discord"
)

type insertHandler struct {
	bot *discord.Bot
}

func New(bot *discord.Bot) *insertHandler {
	return &insertHandler{
		bot: bot,
	}
}

func (h *insertHandler) Insert(c *gin.Context) {
	channelId := c.Param("id")
	if channelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "supply channel id",
		})
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	input := string(body)

	msg, err := h.bot.Discord.ChannelMessageSend(channelId, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
