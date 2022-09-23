package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/rlarkin212/wumpdump/api/routes/create"
	"github.com/rlarkin212/wumpdump/api/routes/drop"
	"github.com/rlarkin212/wumpdump/api/routes/health"
	"github.com/rlarkin212/wumpdump/api/routes/insert"
	"github.com/rlarkin212/wumpdump/discord"
)

func RegisterRoutes(router *gin.Engine, bot *discord.Bot) {
	v1 := router.Group("v1")
	register(v1, bot)
}

func register(rg *gin.RouterGroup, bot *discord.Bot) {
	insert := insert.New(bot)
	health := health.New()

	// tables := tables.New(bot)
	create := create.New(bot)
	drop := drop.New(bot)

	//!data routes
	rg.POST("/insert", insert.Insert)
	rg.GET("/health", health.Health)

	//!table routes
	// rg.GET("/tables", tables.Tables)
	rg.POST("/create", create.Create)
	rg.DELETE("/drop", drop.Drop)
}
