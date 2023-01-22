package main

import (
	"log"
	"nullbot/config/run"
	"nullbot/plugins/compiler"
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

	bot.Run()
}
