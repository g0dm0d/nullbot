package ping

import (
	"nullbot/config/run"

	"github.com/bwmarrin/discordgo"
)

func RegisterPlugin() *run.Plugin {
	return &run.Plugin{
		Name: "Ping",
		Commands: []run.Command{
			{
				Command: &discordgo.ApplicationCommand{
					Name:        "ping",
					Description: "test command",
				},
				Func: PingPong,
			},
		},
	}
}
