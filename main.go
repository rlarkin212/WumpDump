package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rlarkin212/wumpdump/api"
	"github.com/rlarkin212/wumpdump/config"
	"github.com/rlarkin212/wumpdump/discord"
)

func main() {
	config, err := config.LoadConfig("./", "config", "yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	dg := discord.Init(config)
	go dg.Start()

	api := api.Init(dg)
	go api.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
