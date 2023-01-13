package config

import (
	"log"
	"nullbot/tools"
	"os"
)

type App struct {
	Ctx tools.Ctx
}

func New() *App {
	app := new(App)
	dir, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	app.Ctx = tools.Ctx{
		TmpDir: dir,
	}
	return app
}
