package main

import (
	"log"
	"nullbot/config/run"
	"nullbot/plugins/gcc"
	"nullbot/plugins/python"
	"nullbot/plugins/rust"
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
	bot.RegisterPlugin(python.RegisterPlugin())
	bot.RegisterPlugin(rust.RegisterPlugin())

	bot.Run()
}
