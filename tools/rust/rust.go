package rust

import (
	"log"
	"nullbot/tools"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func DiscordRust(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	if option, ok := optionMap["code"]; ok {
		test := i.Interaction.ApplicationCommandData().Resolved.Attachments[option.Value.(string)]
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You're code return:",
				Files: []*discordgo.File{
					{
						ContentType: "text/plain",
						Name:        "output.txt",
						Reader:      strings.NewReader(runRust(tools.GetFile(test.URL))),
					},
				},
			},
		})
	}
}

func runRust(code string) string {
	result, err := exec.Command("docker", "run", "-i", "--rm", "--name", "script-rust", "rust-compile", "sh", "compiler.sh", code).Output()
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	return string(result)
}
