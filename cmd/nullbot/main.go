package main

import (
	"log"
	"nullbot/config/run"
	"nullbot/plugins/compiler"
	"nullbot/plugins/ping"
)

func main() {
	configuration, err := run.Init()
	if err != nil {
		log.Panic(err)
	}

	bot, err := run.NewBot(configuration.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.RegisterPlugin(compiler.RegisterPlugin())
	bot.RegisterPlugin(ping.RegisterPlugin())

	bot.Run()
}
