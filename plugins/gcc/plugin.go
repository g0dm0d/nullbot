package gcc

import (
	"nullbot/config/run"

	"github.com/bwmarrin/discordgo"
)

func RegisterPlugin() *run.Plugin {
	return &run.Plugin{
		Name: "gcc",
		Commands: []run.Command{
			{
				Command: &discordgo.ApplicationCommand{
					Name:        "gcc",
					Description: "compile C code",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionAttachment,
							Name:        "code",
							Description: "drop code file",
							Required:    true,
						},
					},
				},
				Func: DiscordGcc,
			},
		},
	}
}
