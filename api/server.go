package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rlarkin212/wumpdump/api/routes"
	"github.com/rlarkin212/wumpdump/discord"
)

type api struct {
	Router *gin.Engine
	Bot    *discord.Bot
}

func Init(bot *discord.Bot) *api {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.ForceConsoleColor()

	api := &api{
		Router: router,
		Bot:    bot,
	}

	return api
}

func (a *api) Start() {
	routes.RegisterRoutes(a.Router, a.Bot)

	a.Router.Run(":5000")
}
