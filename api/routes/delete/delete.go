package delete

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/wumpdump/discord"
	"github.com/rlarkin212/wumpdump/entity/dto"
)

type deleteHandler struct {
	bot *discord.Bot
}

func New(bot *discord.Bot) *deleteHandler {
	return &deleteHandler{
		bot: bot,
	}
}

func (h *deleteHandler) Delete(c *gin.Context) {
	channelId := c.Param("id")
	messageId := c.Param("m")

	if channelId == "" || messageId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "please supply channel id & message id",
		})
		return
	}

	err := h.bot.Discord.ChannelMessageDelete(channelId, messageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted": messageId,
	})
}

func (h *deleteHandler) BulkDelete(c *gin.Context) {
	channelId := c.Param("id")
	if channelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "please supply channel id & message id",
		})

		return
	}

	var bulkDelete dto.BulkDelte
	if err := c.ShouldBindJSON(&bulkDelete); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := h.bot.Discord.ChannelMessagesBulkDelete(channelId, bulkDelete.Messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted": bulkDelete.Messages,
	})
}
