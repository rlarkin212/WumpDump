package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/rlarkin212/wumpdump/api/routes/alter"
	"github.com/rlarkin212/wumpdump/api/routes/create"
	"github.com/rlarkin212/wumpdump/api/routes/drop"
	"github.com/rlarkin212/wumpdump/api/routes/health"
	"github.com/rlarkin212/wumpdump/api/routes/insert"
	"github.com/rlarkin212/wumpdump/api/routes/tables"
	"github.com/rlarkin212/wumpdump/discord"
)

func RegisterRoutes(router *gin.Engine, bot *discord.Bot) {
	v1 := router.Group("v1")
	register(v1, bot)
}

func register(rg *gin.RouterGroup, bot *discord.Bot) {
	insert := insert.New(bot)
	health := health.New()

	tables := tables.New(bot)
	create := create.New(bot)
	drop := drop.New(bot)
	alter := alter.New(bot)

	//!data routes
	rg.POST("/insert/:id", insert.Insert)
	rg.GET("/health", health.Health)

	//!table routes
	rg.GET("/tables", tables.Tables)
	rg.POST("/create", create.Create)
	rg.DELETE("/drop/:id", drop.Drop)
	rg.PATCH("/alter/:id", alter.Alter)
}
