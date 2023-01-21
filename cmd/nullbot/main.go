package main

import (
	"log"
	"nullbot/config/run"
	"nullbot/plugins/gcc"
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

	bot.RegisterPlugin(gcc.RegisterPlugin())

	bot.Run()
}
