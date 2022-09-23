package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthHandler struct {
}

func New() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"UwU": "https://cdn.discordapp.com/attachments/736954184909193220/958125603935105124/FO9kDDvVUAAFPHr.png",
	})
}
