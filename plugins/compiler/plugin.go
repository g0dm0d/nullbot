package compiler

import (
	"nullbot/config/run"
	"nullbot/plugins/compiler/gcc"
	"nullbot/plugins/compiler/python"
	"nullbot/plugins/compiler/rust"

	"github.com/bwmarrin/discordgo"
)

func RegisterPlugin() *run.Plugin {
	return &run.Plugin{
		Name: "Compiler",
		Commands: []run.Command{
			{
				Command: &discordgo.ApplicationCommand{
					Name:        "rust",
					Description: "compile rust code",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionAttachment,
							Name:        "code",
							Description: "drop code file",
							Required:    true,
						},
					},
				},
				Func: rust.DiscordRust,
			},
			{
				Command: &discordgo.ApplicationCommand{
					Name:        "python",
					Description: "run python script",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionAttachment,
							Name:        "code",
							Description: "drop code file",
							Required:    true,
						},
					},
				},
				Func: python.DiscordPython,
			},
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
				Func: gcc.DiscordGcc,
			},
		},
	}
}
